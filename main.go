package main

import (
	"os"

	"github.com/elastic/beats/libbeat/beat"

	"github.com/caibaoying/axtx_beats/beater"
)

func main() {
	err := beat.Run("axtx_beats", "", beater.New)
	if err != nil {
		os.Exit(1)
	}
}
