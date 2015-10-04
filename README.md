# go-steam-resources

This is a Go wrapper around Steam resources ([Protobufs](https://github.com/SteamRE/SteamKit/tree/master/Resources/Protobufs) and [SteamLanguage](https://github.com/SteamRE/SteamKit/tree/master/Resources/SteamLanguage)) from [SteamKit](https://github.com/SteamRE/SteamKit). It's primarily intended for [go-steam](https://github.com/13k/go-steam) extension module authors.

This repository is heavily based on [node-steam-resources](https://github.com/seishun/node-steam-resources).

*SteamLanguage resources are WIP*

# Protobufs

For each protobuf message or enum available in SteamKit2, there is an equivalently named struct or constant. They lie in the same hierarchy as in SteamKit2, with packages as lower-cased namespaces, except that `Internal` namespaces are removed. (If you see a mismatch, consider that a bug.)

For example, the `CMsgGCTopCustomGamesList` message from [dota_gcmessages_common.proto](https://github.com/SteamRE/SteamKit/blob/master/Resources/Protobufs/dota/dota_gcmessages_common.proto) is available as `SteamKit2.GC.Dota.Internal.CMsgGCTopCustomGamesList` in SteamKit2 and as `CMsgGCTopCustomGamesList` in package `github.com/13k/go-steam-resources/protobuf/gc/dota`.

A table of the mapping between SteamKit2's namespaces to go-steam-resources packages (relative to `github.com/13k/go-steam-resources`):

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
import "github.com/13k/go-steam-resources/protobuf"

steamId := uint64(1234567890)
msgHdr := protobuf.CMsgProtoBufHeader{}
msgHdr.Steamid = &steamId
// ...
```
