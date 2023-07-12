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

	// Create a new IPNS record and bind it to a CID.
	ipns, err := client.CreateIPNS(cid)
	if err != nil {
		panic(err)
	}
	fmt.Println("IPNS is: ", ipns)

	// Update the IPNS record to a new CID.
	newCID := "QmNYERzV2LfD2kkfahtfv44ocHzEFK1sLBaE7zdcYT2GAZ"
	if err := client.UpdateIPNS(ipns, newCID); err != nil {
		panic(err)
	}
	fmt.Println("update IPNS success!")
}
