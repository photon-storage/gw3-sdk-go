# gw3-sdk-go

This repository contains the Golang SDK for Gateway3, an IPFS decentralized gateway. The Gateway3 SDK provides a simple and easy-to-use interface to interact with IPFS, enabling developers to build scalable and distributed applications.

## Features

- Seamless integration with IPFS network
- Store and retrieve files on the decentralized web
- Simplified and intuitive API
- Built-in support for pinning
- Secure and private data handling
- Lightweight and highly performant

## Getting Started

### Prerequisites

- Golang version 1.19 or higher

### Obtain Access Key and Access Secret

To use the Gateway3 SDK, you need to obtain an access key and access secret. You can get these by logging in to the Photon Storage website at https://www.gw3.io/.

### Installation

To install the Gateway3 SDK, use the following command:

```sh
go get -u github.com/photon-storage/gw3-sdk-go
```

### Usage

Import the Gateway3 SDK in your Golang project:

```go
import "github.com/photon-storage/gw3-sdk-go"
```

Here's a simple example demonstrating the usage of the Gateway3 IPFS Gateway SDK:

```go
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
```

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for more information.
