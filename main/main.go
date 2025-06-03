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

	entryID := entryIDs[0]
	entry, err := mlClient.GetEntry(entryID)
	if err != nil {
		panic(err)
	}

	fmt.Println(entry.RepositoryID)

	entry.ImagedBy = "Donald M."

	msg, err := mlClient.UpdateEntry(entry.ID, entry)
	if err != nil {
		panic(err)
	}

	fmt.Println(msg)

	entry2, err := mlClient.GetEntry(entry.ID)
	if err != nil {
		panic(err)
	}

	fmt.Println(entry2.ImagedBy)

}
