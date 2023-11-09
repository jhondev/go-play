package main

import (
	"flag"
	"fmt"
	"log"

	"logprocessor.37Widgets.co/processor"
)

type config struct {
	logPath string
}

func main() {
	var cfg config
	flag.StringVar(&cfg.logPath, "path", "", "Log path")
	flag.Parse()
	if cfg.logPath == "" {
		fmt.Println("You must provide a log path using the '-path' flag")
		return
	}
	qualityCtrl, err := processor.ProcessLog(cfg.logPath)
	if err != nil {
		log.Fatalf("Error processing log file: %s", err)
		return
	}

	fmt.Printf("\nResults for log file '%s':\n\n", cfg.logPath)
	for sensor, quality := range qualityCtrl {
		fmt.Printf("%s: %s\n", sensor, quality)
	}
}
