# Go SemVer Aliases [![devel](https://github.com/joseluisq/semver-aliases/actions/workflows/devel.yml/badge.svg)](https://github.com/joseluisq/semver-aliases/actions/workflows/devel.yml) [![codecov](https://codecov.io/gh/joseluisq/semver-aliases/branch/master/graph/badge.svg)](https://codecov.io/gh/joseluisq/semver-aliases) [![Go Report Card](https://goreportcard.com/badge/github.com/joseluisq/semver-aliases)](https://goreportcard.com/report/github.com/joseluisq/semver-aliases) [![GoDoc](https://godoc.org/github.com/joseluisq/semver-aliases?status.svg)](https://pkg.go.dev/github.com/joseluisq/semver-aliases)

> A simple Go package to create deduplicated version aliases based on valid [SemVer](https://semver.org/) release names.

The library takes care of optional prefixed releases (`v`) as well as all version names are *deduplicated* and *sorted* in lexicographic order.

For example this library can be used to create SemVer and custom aliases for [tagging Docker images](https://docs.docker.com/engine/reference/commandline/tag/).

## Usage

```go
package main

import (
	"fmt"

	aliases "github.com/joseluisq/semver-aliases"
)

func main() {
	// 1. Create alias names based on a given release
	versionAliases := aliases.FromVersion("v1.0.0")
	fmt.Printf("%#v\n", versionAliases)
	//	[]string{"1", "1.0", "1.0.0"}

	extras := []string{"latest", "stable", "v1.0.0", "1.0"}
	composed := append(versionAliases, extras...)
	fmt.Printf("%#v\n", composed)
	//	[]string{"1", "1.0", "1.0.0", "latest", "stable", "v1.0.0", "1.0"}

	// 2. Or create alias names based on a list of names (deduplicated and sorted)
	tags := aliases.FromVersionNames(composed)
	fmt.Printf("%#v\n", tags)
	//	[]string{"1", "1.0", "1.0.0", "latest", "stable"}

	// 3. Or create version names with its items suffixed (sorted)
	suffixed := aliases.GetVersionNamesSuffixed(versionAliases, "linux-amd64")
	fmt.Printf("%#v\n", suffixed)
	//	[]string{"1-linux-amd64", "1.0-linux-amd64", "1.0.0-linux-amd64"}
}
```

## Examples

- [examples/aliases.go](./examples/aliases.go)
- [Playground code example](https://goplay.tools/snippet/g6zkaBTq60D)
- Tests examples at [aliases_test.go](./aliases_test.go)

## Documentation

[pkg.go.dev/github.com/joseluisq/semver-aliases](https://pkg.go.dev/github.com/joseluisq/semver-aliases)

## Contributions

Unless you explicitly state otherwise, any contribution intentionally submitted for inclusion in current work by you, as defined in the Apache-2.0 license, shall be dual licensed as described below, without any additional terms or conditions.

Feel free to send some [Pull request](https://github.com/joseluisq/semver-aliases/pulls) or [issue](https://github.com/joseluisq/semver-aliases/issues).

## License

This work is primarily distributed under the terms of both the [MIT license](LICENSE-MIT) and the [Apache License (Version 2.0)](LICENSE-APACHE).

© 2020-present [Jose Quintana](http://git.io/joseluisq)
