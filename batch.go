package main

import (
	"fmt"
	"regexp"
)

var batchJobIDRegexp = regexp.MustCompile(`"JobId": "(.*)"`)

func formatBatchLink(jobID string) string {
	baseURL := "https://eu-west-1.console.aws.amazon.com/batch/v2/home?region=eu-west-1#/jobs/detail/%s"
	return fmt.Sprintf(baseURL, jobID)
}

func batchLink(sel string) []item {
	id := batchJobIDRegexp.FindSubmatch([]byte(sel))
	if len(id) == 2 {
		return []item{
			{
				Label: fmt.Sprintf("AWS Batch GUIv2 for job %s", id[1]),
				Value: formatBatchLink(string(id[1])),
			},
		}
	}
	return nil
}
