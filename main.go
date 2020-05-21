package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
)

type item struct {
	Label string `json:"label,omitempty"`
	Value string `json:"value"`
	Type  string `json:"type"`
}

func (i item) MarshalJSON() ([]byte, error) {
	type wrapper item
	w := wrapper(i)
	if i.Label == "" {
		w.Type = "text"
	} else {
		w.Type = "url"
	}
	return marshal(w)
}

func marshal(v interface{}) ([]byte, error) {
	sb := bytes.Buffer{}
	enc := json.NewEncoder(&sb)
	enc.SetEscapeHTML(false)
	if err := enc.Encode(v); err != nil {
		return nil, err
	}
	return sb.Bytes(), nil
}

type matcherFunc func(string) []item

func main() {
	selection := os.Args[1]
	matchers := []matcherFunc{
		geoCoords,
		unixTimestamp,
		jiraLink,
	}
	result := []item{}
	for _, matcher := range matchers {
		items := matcher(selection)
		if items != nil {
			result = append(result, items...)
		}
	}
	bytes, err := marshal(result)
	if err != nil {
		panic(err)
	}
	fmt.Print(string(bytes))
}
