package goScraping

import (
	"fmt"
	"time"
)

type Scrapper interface {
	Get() error
	Analyse() error
	Save() error
}

func Scraping(scrapper Scrapper, name string, isDebug bool) error {
	if isDebug {
		fmt.Printf("scraping %v get\n", name)
	}
	err := scrapper.Get()
	if err != nil {
		return err
	}

	if isDebug {
		fmt.Printf("scraping %v analyse\n", name)
	}
	err = scrapper.Analyse()
	if err != nil {
		return err
	}

	if isDebug {
		fmt.Printf("scraping %v save\n", name)
	}
	err = scrapper.Save()
	if err != nil {
		return err
	}

	// sleep any time
	time.Sleep(time.Second)

	return nil
}
