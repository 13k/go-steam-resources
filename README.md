# go-steam-resources

This is a Go wrapper around Steam resources
([Protobufs](https://github.com/SteamDatabase/Protobufs) and
[SteamLanguage](https://github.com/SteamRE/SteamKit/tree/master/Resources/SteamLanguage)).

It's primarily intended to use with [go-steam](https://github.com/Philipp15b/go-steam).

This repository is inspired by [node-steam-resources](https://github.com/seishun/node-steam-resources).

Philipp15b's [go-steam](https://github.com/Philipp15b/go-steam) repository already contains (a
version) of these packages, but this repository was created mainly for three reasons:

1. Decoupling auto-generated protobuf packages sounds like a good idea. It can be used on any sort
   of project that needs it, not only with `go-steam`. It also enables different repository
   life-cycles (sometimes not all protobuf changes require changes to code dealing with them and
   vice-versa).
2. `go-steam`'s protobuf packages organization is a bit messy. It's possible to fix it, but since
   reason #1 seems enough for me to extract the repository, I took the opportunity to rework the
   organization too.
3. SteamLanguage resources in `go-steam` requires the `mono` platform (since it simply runs
   SteamKit's steamlang generator which is implemented in C#). A generator in go leveraging `go
   generate`, with no external dependencies, would be nice.

*NOTE: SteamLanguage resources are WIP. For now, the only option is to use go-steam's SteamLanguage
resources.*

# Protobufs

Mapping between SteamDatabase's protobufs files (relative to `SteamDatabase/Protobufs`) to
go-steam-resources packages (relative to `github.com/13k/go-steam-resources`) are set in the file
`resources/config.yml`.

## Installation & Usage

Simply `go get` the desired package and use the exported protobuf messages accordingly.

`go get github.com/13k/go-steam-resources/protobuf`

```go
import (
  steampb "github.com/13k/go-steam-resources/protobuf/steam"
  "github.com/golang/protobuf/proto"
)

steamId := uint64(1234567890)
msgHdr := steampb.CMsgProtoBufHeader{}
msgHdr.Steamid = proto.Uint64(steamId)
// ...
```

# SteamLanguage

WIP.

# Development

## Updating protobufs

Use the provided `scripts/update.sh` shell script. See `update.sh -h` for usage help.

## Generating protobufs

Use the provided `scripts/generate.sh` shell script.
