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

	// Create a new DAG by appending /example.txt
	data := "EThe Times 03/Jan/2009 Chancellor on brink of second bailout for banks"
	cid, err := client.DAGAdd(gw3.EmptyDAGRoot, "/example.txt", []byte(data))
	if err != nil {
		panic(err)
	}
	fmt.Println("Created a new DAG, CID is: ", cid)

	// Remove the /example.txt from the DAG that we just created.
	cid, err = client.DAGRemove(cid, "/example.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println("Removed the /example.txt, CID is: ", cid)
}
