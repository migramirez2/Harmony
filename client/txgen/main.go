package main

import (
	"encoding/binary"
	"encoding/hex"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"path"
	"runtime"
	"sync"
	"time"

	"github.com/simple-rules/harmony-benchmark/blockchain"
	"github.com/simple-rules/harmony-benchmark/client"
	client_config "github.com/simple-rules/harmony-benchmark/client/config"
	"github.com/simple-rules/harmony-benchmark/consensus"
	"github.com/simple-rules/harmony-benchmark/crypto/pki"
	"github.com/simple-rules/harmony-benchmark/log"
	"github.com/simple-rules/harmony-benchmark/node"
	"github.com/simple-rules/harmony-benchmark/p2p"
	proto_node "github.com/simple-rules/harmony-benchmark/proto/node"
)

var (
	version string
	builtBy string
	builtAt string
	commit  string
)

type txGenSettings struct {
	numOfAddress      int
	crossShard        bool
	maxNumTxsPerBatch int
	crossShardRatio   int
}

var (
	utxoPoolMutex sync.Mutex
	setting       txGenSettings
)

type TxInfo struct {
	// Global Input
	shardID   int
	dataNodes []*node.Node
	// Temp Input
	id      [32]byte
	index   uint32
	value   int
	address [20]byte
	// Output
	txs      []*blockchain.Transaction
	crossTxs []*blockchain.Transaction
	txCount  int
}

// Generates at most "maxNumTxs" number of simulated transactions based on the current UtxoPools of all shards.
// The transactions are generated by going through the existing utxos and
// randomly select a subset of them as the input for each new transaction. The output
// address of the new transaction are randomly selected from [0 - N), where N is the total number of fake addresses.
//
// When crossShard=true, besides the selected utxo input, select another valid utxo as input from the same address in a second shard.
// Similarly, generate another utxo output in that second shard.
//
// NOTE: the genesis block should contain N coinbase transactions which add
//       token (1000) to each address in [0 - N). See node.AddTestingAddresses()
//
// Params:
//     subsetId                   - the which subset of the utxo to work on (used to select addresses)
//     shardID                    - the shardID for current shard
//     dataNodes                  - nodes containing utxopools of all shards
// Returns:
//     all single-shard txs
//     all cross-shard txs
func generateSimulatedTransactions(subsetId, numSubset int, shardID int, dataNodes []*node.Node) ([]*blockchain.Transaction, []*blockchain.Transaction) {
	/*
	  UTXO map structure:
	     address - [
	                txId1 - [
	                        outputIndex1 - value1
	                        outputIndex2 - value2
	                       ]
	                txId2 - [
	                        outputIndex1 - value1
	                        outputIndex2 - value2
	                       ]
	               ]
	*/

	txInfo := TxInfo{}
	txInfo.shardID = shardID
	txInfo.dataNodes = dataNodes
	txInfo.txCount = 0

UTXOLOOP:
	// Loop over all addresses
	for address, txMap := range dataNodes[shardID].UtxoPool.UtxoMap {
		if int(binary.BigEndian.Uint32(address[:]))%numSubset == subsetId%numSubset { // Work on one subset of utxo at a time
			txInfo.address = address
			// Loop over all txIds for the address
			for txIdStr, utxoMap := range txMap {
				// Parse TxId
				id, err := hex.DecodeString(txIdStr)
				if err != nil {
					continue
				}
				copy(txInfo.id[:], id[:])

				// Loop over all utxos for the txId
				utxoSize := len(utxoMap)
				batchSize := utxoSize / numSubset
				i := subsetId % numSubset
				counter := 0
				for index, value := range utxoMap {
					counter++
					if batchSize*i < counter && counter > batchSize*(i+1) {
						continue
					}
					txInfo.index = index
					txInfo.value = value

					randNum := rand.Intn(100)

					subsetRatio := 100         // / numSubset
					if randNum < subsetRatio { // Sample based on batch size
						if setting.crossShard && randNum < subsetRatio*setting.crossShardRatio/100 { // 30% cross shard transactions: add another txinput from another shard
							generateCrossShardTx(&txInfo)
						} else {
							generateSingleShardTx(&txInfo)
						}
						if txInfo.txCount >= setting.maxNumTxsPerBatch {
							break UTXOLOOP
						}
					}
				}
			}
		}
	}
	log.Info("UTXO CLIENT", "numUtxo", dataNodes[shardID].UtxoPool.CountNumOfUtxos(), "shardID", shardID)
	log.Debug("[Generator] generated transations", "single-shard", len(txInfo.txs), "cross-shard", len(txInfo.crossTxs))
	return txInfo.txs, txInfo.crossTxs
}

func generateCrossShardTx(txInfo *TxInfo) {
	nodeShardID := txInfo.dataNodes[txInfo.shardID].Consensus.ShardID
	crossShardID := nodeShardID
	// a random shard to spend money to
	for true {
		crossShardID = uint32(rand.Intn(len(txInfo.dataNodes)))
		if crossShardID != nodeShardID {
			break
		}
	}

	//crossShardNode := txInfo.dataNodes[crossShardID]
	//crossShardUtxosMap := crossShardNode.UtxoPool.UtxoMap[txInfo.address]
	//
	//// Get the cross shard utxo from another shard
	//var crossTxin *blockchain.TXInput
	//crossUtxoValue := 0
	//// Loop over utxos for the same address from the other shard and use the first utxo as the second cross tx input
	//for crossTxIdStr, crossShardUtxos := range crossShardUtxosMap {
	//	// Parse TxId
	//	id, err := hex.DecodeString(crossTxIdStr)
	//	if err != nil {
	//		continue
	//	}
	//	crossTxId := [32]byte{}
	//	copy(crossTxId[:], id[:])
	//
	//	for crossShardIndex, crossShardValue := range crossShardUtxos {
	//		crossUtxoValue = crossShardValue
	//		crossTxin = blockchain.NewTXInput(blockchain.NewOutPoint(&crossTxId, crossShardIndex), txInfo.address, crossShardID)
	//		break
	//	}
	//	if crossTxin != nil {
	//		break
	//	}
	//}

	// Add the utxo from current shard
	txIn := blockchain.NewTXInput(blockchain.NewOutPoint(&txInfo.id, txInfo.index), txInfo.address, nodeShardID)
	txInputs := []blockchain.TXInput{*txIn}

	// Add the utxo from the other shard, if any
	//if crossTxin != nil { // This means the ratio of cross shard tx could be lower than 1/3
	//	txInputs = append(txInputs, *crossTxin)
	//}

	// Spend the utxo from the current shard to a random address in [0 - N)
	txout := blockchain.TXOutput{Amount: txInfo.value, Address: pki.GetAddressFromInt(rand.Intn(setting.numOfAddress) + 1), ShardID: crossShardID}

	txOutputs := []blockchain.TXOutput{txout}

	// Spend the utxo from the other shard, if any, to a random address in [0 - N)
	//if crossTxin != nil {
	//	crossTxout := blockchain.TXOutput{Amount: crossUtxoValue, Address: pki.GetAddressFromInt(rand.Intn(setting.numOfAddress) + 1), ShardID: crossShardID}
	//	txOutputs = append(txOutputs, crossTxout)
	//}

	// Construct the new transaction
	tx := blockchain.Transaction{ID: [32]byte{}, TxInput: txInputs, TxOutput: txOutputs, Proofs: nil}

	priKeyInt, ok := client.LookUpIntPriKey(txInfo.address)
	if ok {
		tx.PublicKey = pki.GetBytesFromPublicKey(pki.GetPublicKeyFromScalar(pki.GetPrivateKeyScalarFromInt(priKeyInt)))

		tx.SetID() // TODO(RJ): figure out the correct way to set Tx ID.
		tx.Sign(pki.GetPrivateKeyScalarFromInt(priKeyInt))
	} else {
		log.Error("Failed to look up the corresponding private key from address", "Address", txInfo.address)
		return
	}

	txInfo.crossTxs = append(txInfo.crossTxs, &tx)
	txInfo.txCount++
}

func generateSingleShardTx(txInfo *TxInfo) {
	nodeShardID := txInfo.dataNodes[txInfo.shardID].Consensus.ShardID
	// Add the utxo as new tx input
	txin := blockchain.NewTXInput(blockchain.NewOutPoint(&txInfo.id, txInfo.index), txInfo.address, nodeShardID)

	// Spend the utxo to a random address in [0 - N)
	txout := blockchain.TXOutput{Amount: txInfo.value, Address: pki.GetAddressFromInt(rand.Intn(setting.numOfAddress) + 1), ShardID: nodeShardID}
	tx := blockchain.Transaction{ID: [32]byte{}, TxInput: []blockchain.TXInput{*txin}, TxOutput: []blockchain.TXOutput{txout}, Proofs: nil}

	priKeyInt, ok := client.LookUpIntPriKey(txInfo.address)
	if ok {
		tx.PublicKey = pki.GetBytesFromPublicKey(pki.GetPublicKeyFromScalar(pki.GetPrivateKeyScalarFromInt(priKeyInt)))
		tx.SetID() // TODO(RJ): figure out the correct way to set Tx ID.
		tx.Sign(pki.GetPrivateKeyScalarFromInt(priKeyInt))
	} else {
		log.Error("Failed to look up the corresponding private key from address", "Address", txInfo.address)
		return
	}

	txInfo.txs = append(txInfo.txs, &tx)
	txInfo.txCount++
}

func printVersion(me string) {
	fmt.Fprintf(os.Stderr, "Harmony (C) 2018. %v, version %v-%v (%v %v)\n", path.Base(me), version, commit, builtBy, builtAt)
	os.Exit(0)
}

func main() {
	configFile := flag.String("config_file", "local_config.txt", "file containing all ip addresses and config")
	maxNumTxsPerBatch := flag.Int("max_num_txs_per_batch", 20000, "number of transactions to send per message")
	logFolder := flag.String("log_folder", "latest", "the folder collecting the logs of this execution")
	numSubset := flag.Int("numSubset", 3, "the number of subsets of utxos to process separately")
	duration := flag.Int("duration", 120, "duration of the tx generation in second. If it's negative, the experiment runs forever.")
	versionFlag := flag.Bool("version", false, "Output version info")
	crossShardRatio := flag.Int("cross_shard_ratio", 30, "The percentage of cross shard transactions.")
	flag.Parse()

	if *versionFlag {
		printVersion(os.Args[0])
	}

	// Add GOMAXPROCS to achieve max performance.
	runtime.GOMAXPROCS(1024)

	// Read the configs
	config := client_config.NewConfig()
	config.ReadConfigFile(*configFile)
	shardIDLeaderMap := config.GetShardIDToLeaderMap()

	setting.numOfAddress = 10000
	// Do cross shard tx if there are more than one shard
	setting.crossShard = len(shardIDLeaderMap) > 1
	setting.maxNumTxsPerBatch = *maxNumTxsPerBatch
	setting.crossShardRatio = *crossShardRatio

	// TODO(Richard): refactor this chuck to a single method
	// Setup a logger to stdout and log file.
	logFileName := fmt.Sprintf("./%v/txgen.log", *logFolder)
	h := log.MultiHandler(
		log.StdoutHandler,
		log.Must.FileHandler(logFileName, log.LogfmtFormat()), // Log to file
	)
	log.Root().SetHandler(h)

	// Nodes containing utxopools to mirror the shards' data in the network
	nodes := []*node.Node{}
	for shardID, _ := range shardIDLeaderMap {
		node := node.New(&consensus.Consensus{ShardID: shardID}, nil)
		// Assign many fake addresses so we have enough address to play with at first
		node.AddTestingAddresses(setting.numOfAddress)
		nodes = append(nodes, node)
	}

	// Client/txgenerator server node setup
	clientPort := config.GetClientPort()
	consensusObj := consensus.NewConsensus("0", clientPort, "0", nil, p2p.Peer{})
	clientNode := node.New(consensusObj, nil)

	if clientPort != "" {
		clientNode.Client = client.NewClient(&shardIDLeaderMap)

		// This func is used to update the client's utxopool when new blocks are received from the leaders
		updateBlocksFunc := func(blocks []*blockchain.Block) {
			log.Debug("Received new block from leader", "len", len(blocks))
			for _, block := range blocks {
				for _, node := range nodes {
					if node.Consensus.ShardID == block.ShardID {
						log.Debug("Adding block from leader", "shardID", block.ShardID)
						// Add it to blockchain
						node.AddNewBlock(block)
						utxoPoolMutex.Lock()
						node.UpdateUtxoAndState(block)
						utxoPoolMutex.Unlock()
					} else {
						continue
					}
				}
			}
		}
		clientNode.Client.UpdateBlocks = updateBlocksFunc

		// Start the client server to listen to leader's message
		go func() {
			clientNode.StartServer(clientPort)
		}()
	}

	// Transaction generation process
	time.Sleep(10 * time.Second) // wait for nodes to be ready
	start := time.Now()
	totalTime := float64(*duration)

	client.InitLookUpIntPriKeyMap()
	subsetCounter := 0

	for true {
		t := time.Now()
		if totalTime > 0 && t.Sub(start).Seconds() >= totalTime {
			log.Debug("Generator timer ended.", "duration", (int(t.Sub(start))), "startTime", start, "totalTime", totalTime)
			break
		}
		shardIDTxsMap := make(map[uint32][]*blockchain.Transaction)
		lock := sync.Mutex{}
		var wg sync.WaitGroup
		wg.Add(len(shardIDLeaderMap))

		utxoPoolMutex.Lock()
		log.Warn("STARTING TX GEN", "gomaxprocs", runtime.GOMAXPROCS(0))
		for shardID, _ := range shardIDLeaderMap { // Generate simulated transactions
			go func(shardID uint32) {
				txs, crossTxs := generateSimulatedTransactions(subsetCounter, *numSubset, int(shardID), nodes)

				// Put cross shard tx into a pending list waiting for proofs from leaders
				if clientPort != "" {
					clientNode.Client.PendingCrossTxsMutex.Lock()
					for _, tx := range crossTxs {
						clientNode.Client.PendingCrossTxs[tx.ID] = tx
					}
					clientNode.Client.PendingCrossTxsMutex.Unlock()
				}

				lock.Lock()
				// Put txs into corresponding shards
				shardIDTxsMap[shardID] = append(shardIDTxsMap[shardID], txs...)
				for _, crossTx := range crossTxs {
					for curShardID, _ := range client.GetInputShardIDsOfCrossShardTx(crossTx) {
						shardIDTxsMap[curShardID] = append(shardIDTxsMap[curShardID], crossTx)
					}
				}
				lock.Unlock()
				wg.Done()
			}(shardID)
		}
		wg.Wait()
		utxoPoolMutex.Unlock()

		lock.Lock()
		for shardID, txs := range shardIDTxsMap { // Send the txs to corresponding shards
			go func(shardID uint32, txs []*blockchain.Transaction) {
				SendTxsToLeader(shardIDLeaderMap[shardID], txs)
			}(shardID, txs)
		}
		lock.Unlock()

		subsetCounter++
		time.Sleep(10000 * time.Millisecond)
	}

	// Send a stop message to stop the nodes at the end
	msg := proto_node.ConstructStopMessage()
	peers := append(config.GetValidators(), clientNode.Client.GetLeaders()...)
	p2p.BroadcastMessage(peers, msg)
}

func SendTxsToLeader(leader p2p.Peer, txs []*blockchain.Transaction) {
	log.Debug("[Generator] Sending txs to...", "leader", leader, "numTxs", len(txs))
	msg := proto_node.ConstructTransactionListMessage(txs)
	p2p.SendMessage(leader, msg)
}
