# go-steam-resources

This is a Go wrapper around Steam resources
([Protobufs](https://github.com/SteamRE/SteamKit/tree/master/Resources/Protobufs) and
[SteamLanguage](https://github.com/SteamRE/SteamKit/tree/master/Resources/SteamLanguage)) from
[SteamKit](https://github.com/SteamRE/SteamKit). It's primarily intended for
[go-steam](https://github.com/Philipp15b/go-steam) extension module authors.

This repository is inspired by [node-steam-resources](https://github.com/seishun/node-steam-resources).

Philipp15b's [go-steam](https://github.com/Philipp15b/go-steam) repository already contains (a
version) of these packages, but this repository was created mainly for three reasons:

1. Decoupling auto-generated protobuf packages sounds like a good idea. It can be used on any sort
of project that needs it, not only with `go-steam`. It also enables different repository life-cycles
(sometimes not all protobuf changes require changes to code dealing with them and vice-versa).
2. `go-steam`'s protobuf packages organization doesn't follow SteamKit's organization (which seems
saner) and are a bit messy. It's possible to fix it, but since reason #1 seems enough for me to
extract the repository, I took the opportunity to rework the organization too.
3. SteamLanguage resources in `go-steam` requires the `mono` platform (since it simply runs
SteamKit's steamlang generator which is implemented in C#). A generator in go leveraging `go
generate`, with no external dependencies, would be nice.

*NOTE: SteamLanguage resources are WIP. For now, the only option is to use go-steam's SteamLanguage
resources.*

# Protobufs

For each protobuf message or enum available in SteamKit2, there is an equivalently named struct or
constant. They lie in the same hierarchy as in SteamKit2, with packages as lower-cased namespaces,
except that `Internal` namespaces are removed. (If you see a mismatch, consider that a bug.)

For example, the `CMsgGCTopCustomGamesList` message from
[dota_gcmessages_common.proto](https://github.com/SteamRE/SteamKit/blob/master/Resources/Protobufs/dota/dota_gcmessages_common.proto)
is available as `SteamKit2.GC.Dota.Internal.CMsgGCTopCustomGamesList` in SteamKit2 and as
`CMsgGCTopCustomGamesList` in package `github.com/13k/go-steam-resources/protobuf/gc/dota`.

A table of the mapping between SteamKit2's namespaces to go-steam-resources packages (relative to
`github.com/13k/go-steam-resources`):

SteamKit2 | go-steam-resources
--------- | --------------------------------------------------------------------
SteamKit2.Internal | protobuf
SteamKit2.Unified.Internal | protobuf/unified
SteamKit2.Unified.Internal.Steamworks | protobuf/unified/steamworks
SteamKit2.GC.Internal | protobuf/gc
SteamKit2.GC.CSGO.Internal | protobuf/gc/csgo
SteamKit2.GC.Dota.Internal | protobuf/gc/dota
SteamKit2.GC.TF2.Internal | protobuf/gc/tf2

## Installation & Usage

Simply `go get` the desired package and use the exported protobuf messages accordingly.

`go get github.com/13k/go-steam-resources/protobuf`

```go
import (
  "github.com/13k/go-steam-resources/protobuf"
  "github.com/golang/protobuf/proto"
)

steamId := uint64(1234567890)
msgHdr := protobuf.CMsgProtoBufHeader{}
msgHdr.Steamid = proto.Uint64(steamId)
// ...
```

# SteamLanguage

WIP.

# Development

## Updating SteamKit

Use the provided `update.sh` shell script. See `update.sh -h` for usage help.

## Generating protobufs

Run `go generate` at the top directory of the repository. It will start the generation process that
consists of the steps:

1. `go generate` (without a package name, which means the package in the current directory, which
means the `main` package in `gen.go`) will catch the `go:generate` directive in `gen.go`, which
calls `update.sh`
2. `update.sh` will clean the `./protobuf` package and call `go run gen.go`
3. `gen.go` reads protobuf package definitions in `config.json` and generates package stub files
(`generate.go`, generated from the template in `generate.go.template`) containing `go:generate`
directives (these directives simply call `protoc` with generated arguments from the package
definition). The `./protobuf` package will then be ready to be generated using `go generate`
4. `update.sh` calls `go generate ./protobuf/...`. This will recursively read all package stub files
in `./protobuf` package and delegate to `protoc`

**WARNING:** This process breaks `go generate ./...` at the top directory, since it's actually a
two-phase generation (generating stubs with `gen.go` then generating actual go files with `protoc`).
Ideally a single invocation should be used, but I'm not sure it's possible.
