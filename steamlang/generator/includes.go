package generator

import (
	"fmt"
	"sort"
)

var (
	includesSnippets = map[string]string{
		"emsgmask": `
const EMsgMask uint32 = ^uint32(0x80000000)

func MakeEMsg(e uint32) EMsg {
	return EMsg(e&EMsgMask)
}
		`,
		"protomask": `
const ProtoMask uint32 = 0x80000000

func MaskProto(e uint32) uint32 {
	return e|ProtoMask
}

func IsProto(e uint32) bool {
	return e&ProtoMask > 0
}
		`,
	}
)

// map of packageName => (map of snippetName => bool)
// the values are maps to avoid duplicates
type genIncludes map[string]map[string]bool

func (i genIncludes) add(pkg, name string) {
	if _, ok := includesSnippets[name]; !ok {
		panic(fmt.Errorf("Trying to include snippet %q that doesn't exist", name))
	}

	if i[pkg] == nil {
		i[pkg] = make(map[string]bool)
	}

	i[pkg][name] = true
}

func (i genIncludes) get(pkg string) (snippets []string) {
	names, ok := i[pkg]

	if !ok {
		return
	}

	for name := range names {
		if snippet, ok := includesSnippets[name]; ok {
			snippets = append(snippets, snippet)
		}
	}

	sort.Strings(snippets)
	return
}
