package main

import (
	"fmt"
	"regexp"
)

var latLonRegexp = regexp.MustCompile(`^(?P<lat>\d+\.\d*),[ ]?(?P<lon>-?\d+\.\d*)$`)
var latLonRegexp2 = regexp.MustCompile(`^{\s*"?lat"?:\s*"?(?P<lat>\d+\.\d*)"?,\s*"?lon"?:\s*"?(?P<lon>-?\d+\.\d*)"?\s*}$`)

func geoCoords(sel string) []item {
	l1 := latLonRegexp.FindSubmatch([]byte(sel))
	l2 := latLonRegexp2.FindSubmatch([]byte(sel))
	var ll [][]byte
	if len(l1) == 3 {
		ll = l1
	}
	if len(l2) == 3 {
		ll = l2
	}

	if len(ll) == 3 {
		items := []item{
			{
				Label: "Google Maps",
				Value: fmt.Sprintf(
					"https://www.google.com/maps/search/?api=1&query=%s,%s",
					string(ll[1]),
					string(ll[2]),
				),
			},
			{
				Label: "Apple Maps",
				Value: fmt.Sprintf(
					"https://maps.apple.com/?t=m&ll=%s,%s",
					string(ll[1]),
					string(ll[2]),
				),
			},
		}
		return items
	}

	return nil
}
