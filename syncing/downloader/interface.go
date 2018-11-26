package downloader

import (
	pb "github.com/harmony-one/harmony/syncing/downloader/proto"
)

// DownloadInterface ...
type DownloadInterface interface {
	// Syncing blockchain from other peers.
	// The returned channel is the signal of syncing finish.
	CalculateResponse(request *pb.DownloaderRequest) (*pb.DownloaderResponse, error)
}
