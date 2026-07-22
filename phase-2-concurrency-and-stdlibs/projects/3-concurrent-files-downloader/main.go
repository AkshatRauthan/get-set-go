package main

import (
	"concurrent-files-downloader/downloader"
	"context"
	"log"
	"net/http"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	client := &http.Client{}

	err := downloader.ChunkedFileDownload(ctx, client, "https://drive.usercontent.google.com/download?id=1UtNW-_ANWGqpXqzEiBNJYeuxap5foYth&export=download&authuser=0&confirm=t&uuid=48c1bf11-fad3-467d-a199-4916a9143180&at=ABswASalMRgOVwFuSirzemuSQ2k4%3A1784745560587", 10)
	if err != nil {
		log.Fatal(err)
	}
}
