// Copyright (c) 2014-2016 The btcsuite developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package chaincfg

import (
	"time"

	"github.com/ltcsuite/ltcd/chaincfg/chainhash"
	"github.com/ltcsuite/ltcd/wire"
)

// genesisCoinbaseTx is the coinbase transaction for the genesis blocks for
// the main network, regression test network, and test network (version 3).
var genesisCoinbaseTx = wire.MsgTx{
	Version: 1,
	TxIn: []*wire.TxIn{
		{
			PreviousOutPoint: wire.OutPoint{
				Hash:  chainhash.Hash{},
				Index: 0xffffffff,
			},
			SignatureScript: []byte{
				0x61, 0x69, 0x6c, 0x61, 0x72, 0x74, 0x73, 0x75, 0x41, 0x20, 0x6d, 0x6f, 0x72, 0x46, 0x20, 0x79,
				 0x72, 0x61, 0x6e, 0x6f, 0x69, 0x73, 0x73, 0x69, 0x4d, 0x20, 0x6e, 0x61, 0x69, 0x74, 0x73, 0x69,
				  0x72, 0x68, 0x43, 0x20, 0x73, 0x74, 0x73, 0x65, 0x72, 0x72, 0x41, 0x20, 0x61, 0x65, 0x72, 0x6f,
				   0x4b, 0x20, 0x68, 0x74, 0x72, 0x6f, 0x4e, 0x20, 0x34, 0x31, 0x30, 0x32, 0x2f, 0x62, 0x65, 0x46,
					0x2f, 0x39, 0x31, 0x20, 0x73, 0x65, 0x6d, 0x69, 0x54, 0x20, 0x59, 0x4e, 0x4c, 0x4c, 0x04, 0x01,
					 0x1d, 0x00, 0xff, 0xff, 0x04,

			},
			Sequence: 0xffffffff,
		},
	},
	TxOut: []*wire.TxOut{
		{
			Value: 0x12a05f200,
			PkScript: []byte{
				0xac, 0x88, 0x15, 0xf6, 0xed, 0x16, 0x1a, 0x61, 0xc7,
				 0x94, 0xd6, 0xf9, 0xeb, 0x08, 0xae, 0x69, 0xb7, 0xf9,
				  0xc9, 0x44, 0xec, 0x1c, 0x14, 0xa9, 0x76,
			},
		},
	},
	LockTime: 0,
}

// genesisHash is the hash of the first block in the block chain for the main
// network (genesis block).
var genesisHash = chainhash.Hash([chainhash.HashSize]byte{ // Make go vet happy.
	0x4b, 0xd9, 0xae, 0x11, 0x46, 0x71, 0x8f, 0x86,
	0x42, 0x7b, 0xeb, 0x97, 0x5b, 0x3b, 0x30, 0x84,
	0xf9, 0x03, 0x5f, 0x84, 0x1c, 0xff, 0x60, 0xf8,
	0x06, 0xac, 0xb8, 0xb7, 0x4b, 0x20, 0x56, 0x4e,
 })

// genesisMerkleRoot is the hash of the first transaction in the genesis block
// for the main network.
var genesisMerkleRoot = chainhash.Hash([chainhash.HashSize]byte{ // Make go vet happy.
	0x9b, 0x6f, 0xc2, 0x3c, 0xbe, 0x22, 0xf4, 0x97,
	0x66, 0x60, 0x5c, 0xa8, 0xcb, 0x5d, 0x44, 0x05,
	0xf0, 0x9d, 0x6b, 0x3e, 0x60, 0xcf, 0x2a, 0x35,
	0xe4, 0x2e, 0x01, 0x76, 0x87, 0x7e, 0xe4, 0xb3,
 })

// genesisBlock defines the genesis block of the block chain which serves as the
// public transaction ledger for the main network.
var genesisBlock = wire.MsgBlock{
	Header: wire.BlockHeader{
		Version:    1,
		PrevBlock:  chainhash.Hash{},         // 0000000000000000000000000000000000000000000000000000000000000000
		MerkleRoot: genesisMerkleRoot,        // 97ddfbbae6be97fd6cdf3e7ca13232a3afff2353e29badfab7f73011edd4ced9
		Timestamp:  time.Unix(1392841423, 0),
		Bits:       0x1e0ffff0,
		Nonce:      3236648,
	},
	Transactions: []*wire.MsgTx{&genesisCoinbaseTx},
}

// regTestGenesisHash is the hash of the first block in the block chain for the
// regression test network (genesis block).
var regTestGenesisHash = chainhash.Hash([chainhash.HashSize]byte{ // Make go vet happy.
	0xf9, 0x16, 0xc4, 0x56, 0xfc, 0x51, 0xdf, 0x62,
	0x78, 0x85, 0xd7, 0xd6, 0x74, 0xed, 0x02, 0xdc,
	0x88, 0xa2, 0x25, 0xad, 0xb3, 0xf0, 0x2a, 0xd1,
	0x3e, 0xb4, 0x93, 0x8f, 0xf3, 0x27, 0x08, 0x53,
})

// regTestGenesisMerkleRoot is the hash of the first transaction in the genesis
// block for the regression test network.  It is NOT the same as the merkle root for
// the main network.
var regTestGenesisMerkleRoot = chainhash.Hash([chainhash.HashSize]byte{ // Make go vet happy.
	0xd9, 0xce, 0xd4, 0xed, 0x11, 0x30, 0xf7, 0xb7,
	0xfa, 0xad, 0x9b, 0xe2, 0x53, 0x23, 0xff, 0xaf,
	0xa3, 0x32, 0x32, 0xa1, 0x7c, 0x3e, 0xdf, 0x6c,
	0xfd, 0x97, 0xbe, 0xe6, 0xba, 0xfb, 0xdd, 0x97,
 })

// regTestGenesisBlock defines the genesis block of the block chain which serves
// as the public transaction ledger for the regression test network.
var regTestGenesisBlock = wire.MsgBlock{
	Header: wire.BlockHeader{
		Version:    1,
		PrevBlock:  chainhash.Hash{},         // 0000000000000000000000000000000000000000000000000000000000000000
		MerkleRoot: regTestGenesisMerkleRoot, // 97ddfbbae6be97fd6cdf3e7ca13232a3afff2353e29badfab7f73011edd4ced9
		Timestamp:  time.Unix(1296688602, 0), // 2011-02-02 23:16:42 +0000 UTC
		Bits:       0x207fffff,               // 545259519 [7fffff0000000000000000000000000000000000000000000000000000000000]
		Nonce:      0,
	},
	Transactions: []*wire.MsgTx{&genesisCoinbaseTx},
}

// testNet4GenesisHash is the hash of the first block in the block chain for the
// test network (version 4).
var testNet4GenesisHash = chainhash.Hash([chainhash.HashSize]byte{ // Make go vet happy.
	0xee, 0x6b, 0x40, 0x9e, 0x82, 0x15, 0x46, 0x57,
	0xbe, 0xaf, 0xb4, 0xf2, 0x55, 0x7a, 0x9a, 0x1e,
	0x74, 0x54, 0xd4, 0x76, 0x3a, 0x18, 0xe7, 0xc3, 
	0x92, 0x00, 0xe6, 0xb5, 0x88, 0x18, 0x27, 0xa4,
})

// testNet3GenesisMerkleRoot is the hash of the first transaction in the genesis
// block for the test network (version 3).  It is the same as the merkle root
// for the main network.
var testNet3GenesisMerkleRoot = genesisMerkleRoot

// testNet4GenesisMerkleRoot is the hash of the first transaction in the genesis
// block for the test network (version 4).  It is the same as the merkle root
// for the main network.
var testNet4GenesisMerkleRoot = chainhash.Hash([chainhash.HashSize]byte{ // Make go vet happy.
	0x9b, 0x6f, 0xc2, 0x3c, 0xbe, 0x22, 0xf4, 0x97,
	0x66, 0x60, 0x5c, 0xa8, 0xcb, 0x5d, 0x44, 0x05,
	0xf0, 0x9d, 0x6b, 0x3e, 0x60, 0xcf, 0x2a, 0x35,
	0xe4, 0x2e, 0x01, 0x76, 0x87, 0x7e, 0xe4, 0xb3,
})

// testNet4GenesisBlock defines the genesis block of the block chain which
// serves as the public transaction ledger for the test network (version 4).
var testNet4GenesisBlock = wire.MsgBlock{
	Header: wire.BlockHeader{
		Version:    1,
		PrevBlock:  chainhash.Hash{},          // 0000000000000000000000000000000000000000000000000000000000000000
		MerkleRoot: testNet4GenesisMerkleRoot, // 97ddfbbae6be97fd6cdf3e7ca13232a3afff2353e29badfab7f73011edd4ced9
		Timestamp:  time.Unix(1494757042, 0),
		Bits:       0x1e0ffff0,
		Nonce:      2231829,
	},
	Transactions: []*wire.MsgTx{&genesisCoinbaseTx},
}

// simNetGenesisHash is the hash of the first block in the block chain for the
// simulation test network.
var simNetGenesisHash = chainhash.Hash([chainhash.HashSize]byte{ // Make go vet happy.
	0xf6, 0x7a, 0xd7, 0x69, 0x5d, 0x9b, 0x66, 0x2a,
	0x72, 0xff, 0x3d, 0x8e, 0xdb, 0xbb, 0x2d, 0xe0,
	0xbf, 0xa6, 0x7b, 0x13, 0x97, 0x4b, 0xb9, 0x91,
	0x0d, 0x11, 0x6d, 0x5c, 0xbd, 0x86, 0x3e, 0x68,
})

// simNetGenesisMerkleRoot is the hash of the first transaction in the genesis
// block for the simulation test network.  It is the same as the merkle root for
// the main network.
var simNetGenesisMerkleRoot = genesisMerkleRoot

// simNetGenesisBlock defines the genesis block of the block chain which serves
// as the public transaction ledger for the simulation test network.
var simNetGenesisBlock = wire.MsgBlock{
	Header: wire.BlockHeader{
		Version:    1,
		PrevBlock:  chainhash.Hash{},         // 0000000000000000000000000000000000000000000000000000000000000000
		MerkleRoot: simNetGenesisMerkleRoot,  // 4a5e1e4baab89f3a32518a88c31bc87f618f76673e2cc77ab2127b7afdeda33b
		Timestamp:  time.Unix(1401292357, 0), // 2014-05-28 15:52:37 +0000 UTC
		Bits:       0x207fffff,               // 545259519 [7fffff0000000000000000000000000000000000000000000000000000000000]
		Nonce:      2,
	},
	Transactions: []*wire.MsgTx{&genesisCoinbaseTx},
}
