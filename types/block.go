// Copyright 2019 Conflux Foundation. All rights reserved.
// Conflux is free software and distributed under GNU General Public License.
// See http://www.gnu.org/licenses/

package types

import "github.com/ethereum/go-ethereum/common/hexutil"

// BlockHeader represents a block header in Conflux.
type BlockHeader struct {
	Hash                  Hash            `json:"hash"`
	ParentHash            Hash            `json:"parentHash"`
	Height                *hexutil.Big    `json:"height"`
	Miner                 Address         `json:"miner"`
	DeferredStateRoot     Hash            `json:"deferredStateRoot"`
	DeferredReceiptsRoot  Hash            `json:"deferredReceiptsRoot"`
	DeferredLogsBloomHash Hash            `json:"deferredLogsBloomHash"`
	Blame                 uint32          `json:"blame"`
	TransactionsRoot      Hash            `json:"transactionsRoot"`
	EpochNumber           *hexutil.Big    `json:"epochNumber,omitempty"`
	GasLimit              *hexutil.Big    `json:"gasLimit"`
	Timestamp             *hexutil.Uint64 `json:"timestamp"`
	Difficulty            *hexutil.Big    `json:"difficulty"`
	RefereeHashes         []Hash          `json:"refereeHashes"`
	Stable                *bool           `json:"stable,omitempty"`
	Adaptive              bool            `json:"adaptive"`
	Nonce                 *hexutil.Big    `json:"nonce"`
	Size                  *hexutil.Big    `json:"size,omitempty"`
}

// BlockSummary includes block header and a list transaction hashes
type BlockSummary struct {
	BlockHeader
	Transactions []Hash `json:"transactions"`
}

// Block represents a block in Conflux, including block header
// and a list of detailed transactions.
type Block struct {
	BlockHeader
	Transactions []Transaction `json:"transactions"`
}