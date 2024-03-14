package main

import (
	"github.com/Dorrrke/GophKeeper/cmd"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	cmd.Execute()
}
