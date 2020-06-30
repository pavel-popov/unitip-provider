// "LogStreamName": "sfexec-stg/default/96efd666-1f1f-4aa6-87f8-22353e82d850"
package main

import (
	"fmt"
	"regexp"
	"strings"
)

var cloudWatchRegexp = regexp.MustCompile(`"LogStreamName": "(.*)"$`)

func formatCloudWatchLink(logStreamName string) string {
	baseUrl := "https://eu-west-1.console.aws.amazon.com/cloudwatch/home?region=eu-west-1#logsV2:log-groups/log-group/$252Faws$252Fbatch$252Fjob/log-events"
	return fmt.Sprintf("%s/%s", baseUrl, strings.Replace(logStreamName, "/", "$252F", -1))
}

func cloudWatchLink(sel string) []item {
	ll := cloudWatchRegexp.FindSubmatch([]byte(sel))
	if len(ll) == 2 {
		return []item{
			{
				Label: "CloudWatch Batch Logs",
				Value: formatCloudWatchLink(string(ll[1])),
			},
		}
	}
	return nil
}
