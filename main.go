package main

import "fmt"

var (
	GitRelease string
	GitCommit  string
	GitBranch  string
)

func main() {
	fmt.Printf("GitRelease : %s \nGitCommit : %s\nGitBranch : %s", GitRelease, GitCommit, GitBranch)
}
