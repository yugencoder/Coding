package X

import "strings"

func simplifyPath(path string) string {
	paths := strings.Split(path, "/")
	validPaths := []string{}

	for _, p := range paths {
		if p == "" || p == "." {
			continue
		}

		if p == ".." {
			if len(validPaths) > 0 {
				validPaths = validPaths[:len(validPaths)-1]
			}
			continue
		}

		validPaths = append(validPaths, p)
	}

	return "/" + strings.Join(validPaths, "/")
}
