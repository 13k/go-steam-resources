# go-steam-resources

This is a Go wrapper around Steam resources
([Protobufs](https://github.com/SteamDatabase/Protobufs) and
[SteamLanguage](https://github.com/SteamRE/SteamKit/tree/master/Resources/SteamLanguage)).

It's primarily intended to use with [go-steam](https://github.com/13k/go-steam).

This repository is inspired by
[node-steam-resources](https://github.com/seishun/node-steam-resources).

[go-steam](https://github.com/Philipp15b/go-steam) (or the more recent
[fork by faceit](https://github.com/faceit/go-steam)) already contain (a version) of
these packages, but this repository was created mainly for three reasons:

1. Decoupling auto-generated protobuf packages sounds like a good idea. It can be used on any sort
   of project that needs it, not only with `go-steam`. It also enables different repository
   life-cycles (sometimes not all protobuf changes require changes to code dealing with them and
   vice-versa).
2. `go-steam`'s protobuf packages organization is a bit messy. It's possible to fix it, but since
   reason #1 seems enough for me to extract the repository, I took the opportunity to rework the
   organization too.
3. SteamLanguage resources in `go-steam` requires the `mono` platform (since it simply runs
   SteamKit's steamlang generator which is implemented in C#). A generator in go leveraging `go generate`, with no external dependencies, would be nice.

# Protobufs

Mapping between SteamDatabase's protobufs files (relative to `resources/SteamDatabase/Protobufs`) to
go-steam-resources packages (relative to `github.com/13k/go-steam-resources`) are set in the file
`resources/config.yml`.

**NOTE**: Importing of multiple games protobufs in the same package won't work, because some of them
define messages with the same name and protobuf will panic.

## Installation & Usage

`go get github.com/13k/go-steam-resources/v2/steampb`

```go
import (
  steampb "github.com/13k/go-steam-resources/v2/steampb/steam"
  "google.golang.org/protobuf/proto"
)

func f() {
  steamId := uint64(1234567890)
  msgHdr := steampb.CMsgProtoBufHeader{}
  msgHdr.Steamid = proto.Uint64(steamId)
  // ...
}
```

# SteamLanguage

## Installation & Usage

`go get github.com/13k/go-steam-resources/v2/steamlang`

```go
import (
  "github.com/13k/go-steam-resources/v2/steamlang"
)

func f() {
  if message.Result == steamlang.EResult_OK {
    // ...
  }
}
```

# Development

## Updating protobufs

Use the provided `scripts/update.sh` shell script. See `update.sh -h` for usage help.

## Generating protobufs

Use the provided `scripts/generate.sh` shell script.
