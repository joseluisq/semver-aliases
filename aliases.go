// Package aliases creates deduplicated version aliases based on valid Semver release names.
//
// Library takes care of optional prefixed releases (`v`) as well as all version names are deduplicated and sorted in lexicographic order.
// For example this library can be used to create Semver and custom aliases for tagging Docker images.
package aliases

import (
	"sort"
	"strconv"
	"strings"

	"github.com/coreos/go-semver/semver"
)

// FromVersion creates tag aliases based on a valid Semver release.
func FromVersion(version string) []string {
	version = strings.TrimPrefix(
		strings.TrimSpace(version), "v",
	)

	if version == "" {
		return nil
	}

	v, err := semver.NewVersion(version)

	if err != nil {
		return nil
	}

	// just return pre-release versions
	if v.PreRelease != "" {
		return []string{version}
	}

	var tags []string
	var tag = ""

	for _, x := range v.Slice() {
		n := strconv.FormatInt(x, 10)

		if tag == "" {
			tag = n
		} else {
			tag = tag + "." + n
		}

		tags = append(tags, tag)
	}

	return tags
}

// FromVersionNames validates and returns deduplicated tag aliases.
func FromVersionNames(names []string) []string {
	names = filterInputNames(names)

	if len(names) == 0 {
		return nil
	}

	sort.Strings(names)
	j := 0
	for i := 1; i < len(names); i++ {
		if names[j] == names[i] {
			continue
		}
		j++
		names[j] = names[i]
	}

	return names[:j+1]
}

// GetVersionNamesSuffixed returns a version names with its items suffixed.
func GetVersionNamesSuffixed(names []string, suffix string) []string {
	names = FromVersionNames(names)

	if len(names) == 0 {
		return nil
	}

	if suffix == "" {
		return names
	}

	sort.Strings(names)

	var strs []string
	for _, s := range names {
		strs = append(strs, s+"-"+suffix)
	}

	return strs
}

// filterInputNames filters input names like empty values or `v` prefixes
func filterInputNames(names []string) []string {
	if len(names) == 0 {
		return nil
	}

	var strv []string
	for _, s := range names {
		s = strings.TrimSpace(s)
		// skip empty values
		if s != "" {
			// trim version prefixes append to the list
			s = strings.TrimPrefix(s, "v")
			strv = append(strv, s)
		}
	}

	return strv
}
