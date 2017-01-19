package goScraping

import ()

type Empty struct{}

func (Empty) Get() error     { return nil }
func (Empty) Analyse() error { return nil }
func (Empty) Save() error    { return nil }
