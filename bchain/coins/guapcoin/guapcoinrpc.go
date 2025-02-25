package guapcoin

import (
	"encoding/json"

	"github.com/golang/glog"
	"github.com/NikunjVaghasiya01/blockbook/bchain"
	"github.com/NikunjVaghasiya01/blockbook/bchain/coins/btc"
)

// guapcoinRPC is an interface to JSON-RPC bitcoind service.
type guapcoinRPC struct {
	*btc.BitcoinRPC
}

// NewguapcoinRPC returns new guapcoinRPC instance.
func NewguapcoinRPC(config json.RawMessage, pushHandler func(bchain.NotificationType)) (bchain.BlockChain, error) {
	b, err := btc.NewBitcoinRPC(config, pushHandler)
	if err != nil {
		return nil, err
	}

	s := &guapcoinRPC{
		b.(*btc.BitcoinRPC),
	}
	s.RPCMarshaler = btc.JSONMarshalerV1{}
	s.ChainConfig.SupportsEstimateFee = true
	s.ChainConfig.SupportsEstimateSmartFee = false

	return s, nil
}

// Initialize initializes guapcoinRPC instance.
func (b *guapcoinRPC) Initialize() error {
	ci, err := b.GetChainInfo()
	if err != nil {
		return err
	}
	chainName := ci.Chain

	glog.Info("Chain name ", chainName)
	params := GetChainParams(chainName)

	// always create parser
	b.Parser = NewguapcoinParser(params, b.ChainConfig)

	// parameters for getInfo request
	if params.Net == MainnetMagic {
		b.Testnet = false
		b.Network = "livenet"
	} else {
		b.Testnet = true
		b.Network = "testnet"
	}

	glog.Info("rpc: block chain ", params.Name)

	return nil
}