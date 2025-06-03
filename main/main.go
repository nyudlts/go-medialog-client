package main

import (
	"fmt"
	"os"

	medialog_client "medialog-client"
)

func main() {

	configFile := os.Args[1]
	environment := os.Args[2]
	mlClient, err := medialog_client.GetClient(configFile, environment)
	if err != nil {
		panic(err)
	}

	entryIDs, err := mlClient.GetEntryIDs()
	if err != nil {
		panic(err)
	}

	for i, entryID := range entryIDs {
		fmt.Println(i, entryID)
	}

}
