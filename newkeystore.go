package main

import "github.com/XinFinOrg/XDPoSChain/accounts/keystore"
import "log"
import "fmt"

func main3() {
	// capitalKeyStore := keystore.NewKeyStore("", keystore.StandardScryptN,
	// 	keystore.StandardScryptP)


	// println(capitalKeyStore)
	ks := keystore.NewKeyStore("./wallets", keystore.StandardScryptN, keystore.StandardScryptP)
password := "secret"
account, err := ks.NewAccount(password)
if err != nil {
  log.Fatal(err)
}

fmt.Println(account.Address.Hex()) // 0x20F8D42FB0F667F2E53930fed426f225752453b3
}
