package main

import (
	medialog_client "medialog-client"
	"os"
)

func main() {

	configFile := os.Args[1]
	environment := os.Args[2]
	mlClient, err := medialog_client.GetClient(configFile, environment)
	if err != nil {
		panic(err)
	}

	if err := mlClient.GetEntries(); err != nil {
		panic(err)
	}

}
