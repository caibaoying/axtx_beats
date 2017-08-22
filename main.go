package main

import (
	"os"

	"github.com/elastic/beats/libbeat/beat"

	"github.com/caibaoying/axtx_beats/beater"
)

var Name = "axtx_beats"
var Version = "5.4.3"

func main() {
	err := beat.Run(Name, Version, beater.New)
	if err != nil {
		os.Exit(1)
	}
}
