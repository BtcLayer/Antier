package eth

import (
	"context"
	"net/rpc/jsonrpc"
	"strconv"

	"github.com/qtumproject/ethereum-block-processor/jsonrpc"
	"github.com/sirupsen/logrus"
)

func GetLatestBlock(ctx context.Context, logger *logrus.Entry, url string) (latestBlock int64, err error) {
	rpcClient := jsonrpc.NewClient(url, 0)
	rpcResponse, err := rpcClient.Call(ctx, "eth_getblockbynumber", "latest", false)
	if err != nil {
		logger.Error("Invalid endpoint:", err)
		return
	}

	if rpcResponse.Error != nil {
		logger.Error("rpc response error:", rpcResponse.Error)
	}

	var qtumBlock jsonrpc.GetBlockByNumberResponse
	err = jsonrpc.GetBlockByNumberResponse(rpcResponse, &qtumBlock)
	if err != nil {
		logger.Error("could not convert result to qtum.GetBlockByNumberResponse", err)
		return
	}
	latestBlock, _ =strconv.ParseInt(qtumBlock.Number, 0, 64)
	logger.Debug("LatestBlock:", latestBlock)
	return

}

func GetMissingBlock(ctx context.Context, logger *logrus.Entry, url string) (missingBlock int64, err error) {
	
}