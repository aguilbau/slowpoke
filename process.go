package main

import "encoding/json"

type Result map[string]map[string][]string

func (r Result) Add(f *Footprint) {
	if _, ok := r[f.Hash]; ok == false {
		r[f.Hash] = make(map[string][]string)
		r[f.Hash][f.Title] = []string{}
	}
	r[f.Hash][f.Title] = append(r[f.Hash][f.Title], f.Url)
}

func process(c <-chan *Footprint) ([]byte, error) {
	r := Result{}
	for f := range c {
		r.Add(f)
	}
	return json.Marshal(r)
}
