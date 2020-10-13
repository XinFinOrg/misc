package main

import "github.com/xinfinorg/xdcchain/crypto"
import "encoding/hex"

func main() {

	
	// Create an account
	key, err := crypto.GenerateKey()
	
	// Get the address
	address := crypto.PubkeyToAddress(key.PublicKey).Hex()
	// 0x8ee3333cDE801ceE9471ADf23370c48b011f82a6
	
	// Get the private key
	privateKey := hex.EncodeToString(key.D.Bytes())
	// 05b14254a1d0c77a49eae3bdf080f926a2df17d8e2ebdf7af941ea001481e57f

	println(err,address,privateKey)

}
