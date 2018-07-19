package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/TerrenceHo/ABFeature"
)

var (
	Version string
)

func main() {
	file := Flags()
	ABFeature.Start(file)
}

func Flags() string {
	v := flag.Bool("v", false, "Show version of ABFeature.")
	version := flag.Bool("version", false, "Show version of ABFeature.")
	f := flag.String("f", "./config/config.yaml", "Set path to config file.")
	file := flag.String("file", "./config/config.yaml", "Set path to config file.")
	flag.Parse()
	if *version || *v {
		fmt.Printf("ABFeature version %s\n", Version)
		os.Exit(0)
	}

	if f != nil {
		return *f
	} else {
		return *file
	}
}
