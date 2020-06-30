package main

import (
	"fmt"
	"regexp"
	"strings"
)

var batchJobIdRegexp = regexp.MustCompile(`"JobId": "(.*)"`)
var batchJobQueueRegexp = regexp.MustCompile(`"JobQueue": "(.*)"`)

func formatBatchLink(jobQueue, jobId string) string {
	baseUrl := "https://eu-west-1.console.aws.amazon.com/batch/home?region=eu-west-1#/jobs/queue/%s/job/%s?state=SUBMITTED"
	return fmt.Sprintf(baseUrl, strings.Replace(jobQueue, "/", "~2F", -1), jobId)
}

func batchLink(sel string) []item {
	id := batchJobIdRegexp.FindSubmatch([]byte(sel))
	queue := batchJobQueueRegexp.FindSubmatch([]byte(sel))
	if len(id) == 2 && len(queue) == 2 {
		return []item{
			{
				Label: fmt.Sprintf("AWS Batch GUI for job %s", id[1]),
				Value: formatBatchLink(string(queue[1]), string(id[1])),
			},
		}
	}
	return nil
}
