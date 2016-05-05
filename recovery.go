package main

import (
	"fmt"
	"log"
)

var cmdRecovery = &Command{
	Run:   runRecovery,
	Usage: "recovery [index]",
	Short: "prints index status",
	Long: `
Lists details of on-going index shard recoveries of all indices or the specified index.

Example:

  $ es recovery
  $ es recovery twitter
`,
	ApiUrl: "https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-recovery.html",
}

func runRecovery(cmd *Command, args []string) {
	index := ""
	if len(args) >= 1 {
		index = args[0]
	}

	var response map[string]interface{}

	var body string
	if len(index) > 0 {
		body = ESReq("GET", "/"+index+"/_recovery?pretty=1").Do(&response)
	} else {
		body = ESReq("GET", "/_recovery?pretty=1").Do(&response)
	}

	if error, ok := response["error"]; ok {
		status, _ := response["status"]
		log.Fatalf("Error: %v (%v)\n", error, status)
	}
	fmt.Print(body)
}
