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
