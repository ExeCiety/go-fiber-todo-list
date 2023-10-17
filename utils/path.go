package utils

import (
	"os"
	"strings"
)

var RootPath *string

func GetRootPath() string {
	return *RootPath
}

func SetRootPath(cwd string) {
	RootPath = &cwd
}

// SetRootPathForTest is used to set the root path for testing
// Example project dir in /app, and working dir in /app/db/migration/tests
// to get the root path of the project, we need to go 3 levels up (cd ../../../)
// so set stepBack to 3, and we will get the root path
func SetRootPathForTest(stepBack int) {
	cwd := os.Getenv("PWD")
	path := strings.Split(cwd, "/")
	path = path[:len(path)-stepBack]
	fullPath := strings.Join(path, "/") + "/"
	RootPath = &fullPath
}
