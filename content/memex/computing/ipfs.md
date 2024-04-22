---
title: IPFS
summary: The InterPlanetary FileSystem, and its adjacent distributed technologies.
tags:
- computing
- communication
references:
  - url: https://github.com/twdragon/ipfs-cid-local
    name: IPFS CID Calculator
  - url: https://go.dev/
    name: The Go Programming Language
---

## Generating IPFS hashes

For some reason I often find myself wanting to generate the CID that IPFS would generate in [Go](https://go.dev/), but without having to run an entire node. I found [this repo](https://github.com/twdragon/ipfs-cid-local) from which I've extracted and tweaked the following code to generate Blake3 CIDs from arbitrary bytes.

```go
package hash

import (
  "bytes"
  "fmt"

  "github.com/ipfs/boxo/blockservice"
  blockstore "github.com/ipfs/boxo/blockstore"
  chunker "github.com/ipfs/boxo/chunker"
  offline "github.com/ipfs/boxo/exchange/offline"
  "github.com/ipfs/boxo/ipld/merkledag"
  "github.com/ipfs/boxo/ipld/unixfs/importer/balanced"
  uih "github.com/ipfs/boxo/ipld/unixfs/importer/helpers"
  "github.com/ipfs/go-cid"
  "github.com/ipfs/go-datastore"
  dsync "github.com/ipfs/go-datastore/sync"
  multicodec "github.com/multiformats/go-multicodec"
)

func CID(data []byte) (string, error) {
  ds := dsync.MutexWrap(datastore.NewNullDatastore())
  bs := blockstore.NewIdStore(blockstore.NewBlockstore(ds))
  bsrv := blockservice.New(bs, offline.Exchange(bs))
  dsrv := merkledag.NewDAGService(bsrv)

  ufsImportParams := uih.DagBuilderParams{
    Maxlinks:  uih.DefaultLinksPerBlock,
    RawLeaves: true,
    CidBuilder: cid.V1Builder{
      Codec:    uint64(multicodec.Raw),
      MhType:   uint64(multicodec.Blake3),
      MhLength: -1,
    },
    Dagserv: dsrv,
    NoCopy:  false,
  }

  r := bytes.NewReader(data)

  ufsBuilder, err := ufsImportParams.New(chunker.NewSizeSplitter(r, chunker.DefaultBlockSize))
  if err != nil {
    return "", fmt.Errorf("unable to chunk the event to create a Blake3 hash: %w", err)
  }

  nd, err := balanced.Layout(ufsBuilder)
  if err != nil {
    return "", fmt.Errorf("unable to layout the DAG while hashing the event: %w", err)
  }

  return nd.Cid().String(), nil
}
```

This can be done _much_ more simply if you know that your data is never going to exceed 256KiB (which is the default chunker size) with the following. This works cos, without chunking, the data that'd be stored is identical to what you're providing!

```go
import (
  "github.com/ipfs/go-cid"
  "github.com/multiformats/go-multihash"
)

func CID(data []byte) (string, error) {
  hash, err := multihash.Sum(data, multihash.BLAKE3, -1)
  if err != nil {
    return "", err
  }

  return cid.NewCidV1(cid.Raw, hash).String(), nil
}
```
