# Farcaster-Go

A Go module that facilitates smooth and easy integration with the Farcaster network for Go programs. This module provides various packages and functions to interact with the Farcaster network, allowing developers to create, manage, and utilize Farcaster functionalities in their applications.

## Table of Contents
1. [Installation](#installation)
2. [Usage](#usage)
3. [Packages](#packages)
   - [Neynar](#neynar)
   - [Auth](#auth)
   - [Signer](#signer)
   - [Hub](#hub)
   - [Warpcast Client](#warpcast-client)
   - [Web3](#web3)
4. [Contributing](#contributing)
5. [License](#license)

## Installation

To install the Farcaster-Go module, run the following command:

```bash
go get github.com/vocdoni/farcaster-go
```

## Packages

### Neynar

The `neynar` package provides a client to interact with the Neynar API and its Farcaster hub.

**Purpose:**
- To manage interactions with the Neynar API, including setting and retrieving user data, handling channels, and processing webhooks.

**Basic Usage:**

```go
package main

import (
    "context"
    "fmt"
    "github.com/vocdoni/farcaster-go/neynar"
)

func main() {
    apiKey := "your-api-key"
    client, err := neynar.NewNeynarAPI(apiKey)
    if err != nil {
        fmt.Println("Error creating client:", err)
        return
    }

    fid := uint64(3)
    userData, err := client.UserDataByFID(context.Background(), fid)
    if err != nil {
        fmt.Println("Error getting user data:", err)
        return
    }

    fmt.Println("User Data:", userData)
}
```

### Auth

The `auth` package handles authentication processes within the Farcaster network.

**Purpose:**
- To manage the creation and status checking of authentication channels using Warpcast sign-in.

**Basic Usage:**

```go
package main

import (
    "fmt"
    "time"
    "errors"

    "github.com/vocdoni/farcaster-go/auth"
)

func main() {
    client := auth.New()
    response, err := client.CreateChannel("https://myapplication.com")
    if err != nil {
        fmt.Println("Error creating channel:", err)
        return
    }

    fmt.Println("Channel Response:", response)

    for {
        response, err = client.CheckStatus()
        if err != nil {
            if errors.Is(err, auth.ErrAuthenticationPending) {
                time.Sleep(10 * time.Second)
                continue
            }
            panic("something failed")
        }
        fmt.Printf("Auth completed: %s", response)
    }
}
```

### Signer

The `signer` package provides functionalities for create new signers.

**Purpose:**
- To generate key pairs, sign data, and register signers with the Warpcast API.

**Basic Usage:**

```go
package main

import (
    "crypto/ecdsa"
    "crypto/elliptic"
    "crypto/rand"
    "fmt"
    "github.com/vocdoni/farcaster-go/signer"
)

func main() {
    privKey, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
    publicKey, privateKey, _ := signer.GenerateSigner()

    signature, err := signer.SignData(privateKey, []byte("message"))
    if err != nil {
        fmt.Println("Error signing data:", err)
        return
    }

    fmt.Println("Signature:", signature)
}
```

### Hub

The `hub` package provides an API interface for interacting with a Farcaster Hub.

**Purpose:**
- To manage user data, mentions, casts, and followers within a Farcaster Hub.

**Basic Usage:**

```go
package main

import (
    "context"
    "fmt"
    "github.com/vocdoni/farcaster-go/hub"
)

func main() {
    apiKeys := []string{"key1", "value1", "key2", "value2"}
    client, err := hub.NewHubAPI("https://hub.endpoint", apiKeys)
    if err != nil {
        fmt.Println("Error creating hub client:", err)
        return
    }

    fid := uint64(3)
    userData, err := client.UserDataByFID(context.Background(), fid)
    if err != nil {
        fmt.Println("Error getting user data:", err)
        return
    }

    fmt.Println("User Data:", userData)
}
```

### Warpcast Client

The `warpcastclient` package provides access to public functions of the Warpcast API.

**Purpose:**
- To retrieve user profiles and Ethereum addresses associated with FIDs.

**Basic Usage:**

```go
package main

import (
    "fmt"
    "github.com/vocdoni/farcaster-go/warpcastapi"
)

func main() {
    fid := uint64(123456789)
    profile, err := warpcastapi.UserProfileByFID(fid)
    if err != nil {
        fmt.Println("Error getting user profile:", err)
        return
    }

    fmt.Println("User Profile:", profile)
}
```

### Web3

The `web3` package provides the bindings to interact with the keyregistry Farcaster contract on Optimism.

**Purpose:**
- To manage interactions with Ethereum clients and retrieve signers from FIDs.

**Basic Usage:**

```go
package main

import (
    "fmt"
    "github.com/vocdoni/farcaster-go/web3"
)

func main() {
    provider, err := web3.NewFarcasterProvider(nil)
    if err != nil {
        fmt.Println("Error creating provider:", err)
        return
    }

    fid := uint64(3)
    signers, err := provider.SignersFromFID(fid)
    if err != nil {
        fmt.Println("Error getting signers:", err)
        return
    }

    fmt.Println("Signers:", signers)
}
```

## Contributing

Contributions are welcome! Please open an issue or submit a pull request on GitHub.

## License

This project is licensed under the AGPLv3 License - see the [LICENSE](LICENSE) file for details.
```