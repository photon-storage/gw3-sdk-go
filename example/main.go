package main

import (
	"fmt"

	"github.com/photon-storage/gw3-sdk-go"
)

func main() {
	key := "YOUR-ACCESS-KEY"
	secret := "YOUR-ACCESS-SECRET"
	client, err := gw3.NewClient(key, secret)
	if err != nil {
		panic(err)
	}

	data := "EThe Times 03/Jan/2009 Chancellor on brink of second bailout for banks"

	// Post the data to the IPFS network, receiving a CID as a result
	cid, err := client.Post([]byte(data))
	if err != nil {
		panic(err)
	}
	fmt.Println("Data posted to IPFS network, CID is: ", cid)

	// Request the gateway to pin the CID data, ensuring its persistence
	if err := client.Pin(cid); err != nil {
		panic(err)
	}
	fmt.Println("CID data is pinned by the Gateway3")

	// Retrieve the data from the IPFS network using the CID
	got, err := client.Get(cid)
	if err != nil {
		panic(err)
	}
	fmt.Println("Data retrieved from IPFS network: ", string(got))
}
