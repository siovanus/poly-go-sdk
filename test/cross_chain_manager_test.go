package test

import (
	"testing"
	"fmt"
	oc "github.com/ontio/multi-chain/common"

	"github.com/ontio/multi-chain-go-sdk/common"
	"github.com/ontio/ontology-crypto/signature"
	"encoding/json"
	"errors"
	"github.com/ontio/multi-chain-go-sdk"
)



//const (
//	TestNet = "172.168.3.73"
//)


func TestVote(t *testing.T) {
	sdk := ontology_go_sdk.NewMultiChainSdk()
	pri, _ := oc.HexToBytes("5f2fe68215476abb9852cfa7da31ef00aa1468782d5ca809da5c4e1390b8ee45")
	signer, _ := ontology_go_sdk.NewAccountFromPrivateKey(pri, signature.SHA256withECDSA)
	//to, _ := oc.AddressFromBase58("ASUwFccvYFrrWR6vsZhhNszLFNvCLA5qS6")
	sdk.NewWebSocketClient().Connect("ws://172.168.3.73:40335")
	var fromChainId uint64
	fromChainId = 100
	var address = "AQf4Mzu1YJrhz9f3aRkkwSm9n3qhXGSh4p"
	//txHash1, _ := oc.HexToBytes("7575526bc066a3acc6abb134119cd6d4a9041969")

	txHash, _ := sdk.Native.Ccm.Vote(fromChainId, address, "7575526bc066a3acc6abb134119cd6d4a9041969", signer)
	var results *common.SmartContactEvent
	//results := &common.SmartContactEvent{}
	err := errors.New("No result")
	for {
		results, err = sdk.GetSmartContractEvent(txHash.ToHexString())

		if err == nil {
			fmt.Printf("transfer result is %+v \n ", results)
			break
		}
	}
	jsonResultBytes, e := json.Marshal(results)
	if e !=  nil {
		fmt.Printf("marshal results struct error : %+v", e)
	} else {
		fmt.Printf("marshal results struct to json is %+v\n", string(jsonResultBytes))
	}
}


func TestRegisterSideChain(t *testing.T) {
	sdk := ontology_go_sdk.NewMultiChainSdk()
	pri, _ := oc.HexToBytes("5f2fe68215476abb9852cfa7da31ef00aa1468782d5ca809da5c4e1390b8ee45")
	signer, _ := ontology_go_sdk.NewAccountFromPrivateKey(pri, signature.SHA256withECDSA)
	//to, _ := oc.AddressFromBase58("ASUwFccvYFrrWR6vsZhhNszLFNvCLA5qS6")
	sdk.NewWebSocketClient().Connect("ws://138.91.6.125:40335")

	var address = "AQf4Mzu1YJrhz9f3aRkkwSm9n3qhXGSh4p"
	//txHash1, _ := oc.HexToBytes("7575526bc066a3acc6abb134119cd6d4a9041969")

	txHash, _ := sdk.Native.Scm.RegisterSideChain(address, 165, "chain165", 1, signer)
	var results *common.SmartContactEvent
	//results := &common.SmartContactEvent{}
	err := errors.New("No result")
	for {
		results, err = sdk.GetSmartContractEvent(txHash.ToHexString())

		if err == nil {
			fmt.Printf("RegisterSideChain result is %+v \n ", results)
			break
		}
	}
	jsonResultBytes, e := json.Marshal(results)
	if e !=  nil {
		fmt.Printf("marshal results struct error : %+v", e)
	} else {
		fmt.Printf("marshal results struct to json is %+v\n", string(jsonResultBytes))
	}
}