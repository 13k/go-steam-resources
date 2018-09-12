package generator

import (
	"fmt"
)

// map of importPath => alias
type genImports map[string]string

func (i genImports) add(pkg, alias string) {
	if pkgAlias, ok := i[pkg]; ok && alias != pkgAlias {
		panic(fmt.Errorf("Trying to import package %q with different aliases %q and %q", pkg, alias, pkgAlias))
	}

	i[pkg] = alias
}
