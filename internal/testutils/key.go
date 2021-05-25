package testutils

import (
	"bufio"
	"encoding/hex"
	"os"

	"github.com/holiman/uint256"

	"github.com/ethereum/go-ethereum/common"
	gethcore "github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/smartbch/smartbch/internal/ethutils"
)

func GenKeyAndAddr() (string, common.Address) {
	key, _ := crypto.GenerateKey()
	keyHex := hex.EncodeToString(crypto.FromECDSA(key))
	addr := crypto.PubkeyToAddress(key.PublicKey)
	return keyHex, addr
}

func KeysToGenesisAlloc(balance *uint256.Int, keys []string) gethcore.GenesisAlloc {
	alloc := gethcore.GenesisAlloc{}
	for _, hexKey := range keys {
		privKey, data, err := ethutils.HexToPrivKey(hexKey)
		if err != nil {
			panic(err)
		}
		addr := ethutils.PrivKeyToAddr(privKey)
		alloc[addr] = gethcore.GenesisAccount{
			PrivateKey: data,
			Balance:    balance.ToBig(),
		}
	}
	return alloc
}

// read private keys from a file
func ReadKeysFromFile(fname string, count int) (res []string) {
	f, err := os.Open(fname)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	res = make([]string, 0, 8192)
	for scanner.Scan() && len(res) < count {
		//fmt.Printf("Now read %d\n", len(res))
		txt := scanner.Text()
		//_, err := crypto.HexToECDSA(txt)
		//if err != nil {
		//	panic(err)
		//}
		res = append(res, txt)
	}
	return
}
