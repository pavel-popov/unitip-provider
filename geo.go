package main

import (
	"fmt"
	"regexp"
)

var latLonRegexp = regexp.MustCompile(`^(?P<lat>\d+\.\d*),[ ]?(?P<lon>\d+\.\d*)$`)

func geoCoords(sel string) []item {
	ll := latLonRegexp.FindSubmatch([]byte(sel))
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
