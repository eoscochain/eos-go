package token

import "github.com/cochainio/eos-go"

func init() {
	eos.RegisterAction(AN("eosio.token"), ActN("transfer"), &Transfer{})
	eos.RegisterAction(AN("eosio.token"), ActN("issue"), &Issue{})
	eos.RegisterAction(AN("eosio.token"), ActN("create"), &Create{})
}
