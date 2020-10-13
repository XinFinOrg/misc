package main

import "github.com/xinfinorg/xdcchain/accounts/keystore"

func main() {
	capitalKeyStore := keystore.NewKeyStore("", keystore.StandardScryptN,
		keystore.StandardScryptP)


	println(capitalKeyStore)
}
