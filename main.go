package main

import (
	"flag"
	"fmt"
	"sync"
	"time"

	"github.com/gen2brain/beeep"
)

var (
	title    string
	duration string
)

func main() {
	flag.StringVar(&title, "title", "Timer", "Title of the notification")
	flag.StringVar(&duration, "duration", "5s", "Duration of the timer")
	flag.Parse()

	var wg sync.WaitGroup
	alerts := getAlerts(duration)
	for _, a := range alerts {
		wg.Add(1)
		go func(a alert) {
			time.Sleep(a.Duration)
			err := beeep.Notify(title, a.string, "")
			if err != nil {
				fmt.Println(err)
			}
			wg.Done()
		}(a)
	}
	wg.Wait()
}

type alert struct {
	time.Duration
	string
}

func getAlerts(d string) []alert {
	duration, err := time.ParseDuration(d)
	if err != nil {
		panic(err)
	}
	return []alert{
		{duration / 4, "Quarter of the way through timer for " + title},
		{duration / 2, "halfway through " + title},
		{duration * 3 / 4, "three quarters of the way through " + title},
		{duration, "time is up for " + title + "!"},
	}
}
