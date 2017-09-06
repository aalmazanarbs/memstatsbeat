package main

import (
	"os"

	"github.com/elastic/beats/libbeat/beat"

	"github.com/aalmazanarbs/memstatsbeat/beater"
)

func main() {
	err := beat.Run("memstatsbeat", "", beater.New)
	if err != nil {
		os.Exit(1)
	}
}
