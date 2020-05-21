package main

import (
	"fmt"
	"regexp"
	"strconv"
	"time"
)

var unixSecondsRegexp = regexp.MustCompile(`^(\d{10})$`)

const isoTimestampLayoutLocal = "2006-01-02 15:04:05 -0700"
const isoTimestampLayoutUTC = "2006-01-02T15:04:05Z"

// TimeIn returns the time in UTC if the name is "" or "UTC".
// It returns the local time if the name is "Local".
// Otherwise, the name is taken to be a location name in
// the IANA Time Zone database, such as "Africa/Lagos".
func TimeIn(t time.Time, name string) (time.Time, error) {
	loc, err := time.LoadLocation(name)
	if err == nil {
		t = t.In(loc)
	}
	return t, err
}

func unixTimestamp(sel string) []item {
	ll := unixSecondsRegexp.FindSubmatch([]byte(sel))
	if len(ll) == 2 {
		ts, err := strconv.Atoi(string(ll[1]))
		if err != nil {
			panic(err)
		}
		t := time.Unix(int64(ts), 0)
		utcTs, err := TimeIn(t, "UTC")
		if err != nil {
			panic(err)
		}
		items := []item{
			{
				Value: fmt.Sprintf("%s", t.Format(isoTimestampLayoutLocal)),
			},
			{
				Value: fmt.Sprintf("%s", utcTs.Format(isoTimestampLayoutUTC)),
			},
		}
		return items
	}

	return nil
}
