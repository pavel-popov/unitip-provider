package main

import (
	"fmt"
	"regexp"
)

var jiraRegexp = regexp.MustCompile(`^(ch|CH)-?(\d+)$`)

func jiraLink(sel string) []item {
	ll := jiraRegexp.FindSubmatch([]byte(sel))
	if len(ll) == 3 {
		items := []item{
			{
				Label: fmt.Sprintf("ClubHouse Story CH%s", ll[2]),
				Value: fmt.Sprintf("https://app.clubhouse.io/findhotel/story/%s/", ll[2]),
			},
		}
		return items
	}

	return nil
}
