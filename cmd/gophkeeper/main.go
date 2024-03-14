package main

import (
	"fmt"

	"github.com/Dorrrke/GophKeeper/cmd"
	_ "github.com/mattn/go-sqlite3"
)

var (
	// buildVersion - версия сборки.
	buildVersion string
	// buildDate - дата сборки.
	buildDate string
	// buildCommit - комментарии к сборке.
	buildCommit string
)

func main() {
	if buildVersion == "" {
		buildVersion = "N/A"
	} else {
		fmt.Printf("Build version: %s\n", buildVersion)
	}

	if buildDate == "" {
		buildDate = "N/A"
	} else {
		fmt.Printf("Build date: %s\n", buildDate)
	}

	if buildCommit == "" {
		buildCommit = "N/A"
	} else {
		fmt.Printf("Build commit: %s\n", buildCommit)
	}
	cmd.Execute()
}
