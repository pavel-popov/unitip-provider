package main

import (
	"fmt"
	"regexp"
)

var sfnRegexp = regexp.MustCompile(`^arn:aws:states:(.*):([0-9]*):execution:(.*)$`)

func formatSFNLink(region, accountId, executionId string) string {
	baseUrl := "https://%s.console.aws.amazon.com/states/home?region=%s#/executions/details/arn:aws:states:%s:%s:execution:%s"
	return fmt.Sprintf(baseUrl, region, region, region, accountId, executionId)
}

func sfnLink(sel string) []item {
	arn := sfnRegexp.FindSubmatch([]byte(sel))
	if len(arn) == 4 {
		region := string(arn[1])
		accountId := string(arn[2])
		executionId := string(arn[3])
		return []item{
			{
				Label: fmt.Sprintf("SFN GUI for execution %s", executionId),
				Value: formatSFNLink(region, accountId, executionId),
			},
		}
	}
	return nil
}
