package config

import (
    "math/big"
    "github.com/ethereum/go-ethereum/params"
)

var LuwaChainConfig = &params.ChainConfig{
    ChainID:                 big.NewInt(9882),     // ✔ Chain ID cusub
    HomesteadBlock:          big.NewInt(0),
    EIP150Block:             big.NewInt(0),
    EIP155Block:             big.NewInt(0),
    EIP158Block:             big.NewInt(0),
    ByzantiumBlock:          big.NewInt(0),
    ConstantinopleBlock:     big.NewInt(0),
    PetersburgBlock:         big.NewInt(0),
    IstanbulBlock:           big.NewInt(0),
    MuirGlacierBlock:        big.NewInt(0),
    BerlinBlock:             big.NewInt(0),
    LondonBlock:             big.NewInt(0),
    ArrowGlacierBlock:       big.NewInt(0),
    GrayGlacierBlock:        big.NewInt(0),
    ShanghaiTime:            newUint64(0),
    CancunTime:              newUint64(0),
    PragueTime:              newUint64(0),
    OsakaTime:               newUint64(0),
    TerminalTotalDifficulty: big.NewInt(0), // PoS off (full PoW)
}

// Helper function
func newUint64(v uint64) *uint64 {
    return &v
}
