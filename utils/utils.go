package utils

import (
	"bytes"
	"encoding/binary"
	"log"
	"strconv"
	"strings"
)

// ConvertFixedDataIntoByteArray converts an empty interface data to a byte array
func ConvertFixedDataIntoByteArray(data interface{}) []byte {
	buff := new(bytes.Buffer)
	err := binary.Write(buff, binary.BigEndian, data)
	if err != nil {
		log.Panic(err)
	}
	return buff.Bytes()
}

// ConvertIntoInts is to convert '1,2,3,4' into []int{1,2,3,4}.
func ConvertIntoInts(data string) []int {
	var res = []int{}
	items := strings.Split(data, ",")
	for _, value := range items {
		intValue, err := strconv.Atoi(value)
		if err == nil {
			res = append(res, intValue)
		}
	}
	return res
}
