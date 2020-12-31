package main

import "github.com/XinFinOrg/XDPoSChain/accounts/keystore"

func main() {
	capitalKeyStore := keystore.NewKeyStore("", keystore.StandardScryptN,
	keystore.StandardScryptP)
	println(capitalKeyStore)
}
