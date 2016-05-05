package main

import (
	"log"
	"os"
)

var cmdPutMapping = &Command{
	Run:   runPutMapping,
	Usage: "put-mapping <index> <type>",
	Short: "create or update mapping",
	Long: `
Creates or updates a mapping for a document type on an index.
https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-put-mapping.html

Example:

  $ es put-mapping twitter tweet < tweet-mapping.json
`,
	ApiUrl: "https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-put-mapping.html",
}

func runPutMapping(cmd *Command, args []string) {
	if len(args) != 2 {
		cmd.printUsage()
		os.Exit(1)
	}

	index := args[0]
	doctype := args[1]

	var response struct {
		Ok     bool   `json:"ok,omitempty"`
		Ack    bool   `json:"acknowledged,omitempty"`
		Error  string `json:"error,omitempty"`
		Status int    `json:"status,omitempty"`
	}

	data := getJsonFromStdin()
	req := ESReq("PUT", "/"+index+"/"+doctype+"/_mapping")
	req.SetBodyJson(data)
	req.Do(&response)
	if len(response.Error) > 0 {
		log.Fatalf("Error: %v (%v)\n", response.Error, response.Status)
	}
}
