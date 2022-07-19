package log

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"os"
)

func notifyPapertrail(opts *Options) {
	token := os.Getenv("PAPERTRAIL_LOGGER_TOKEN")
	url := "https://logs.collector.solarwinds.com/v1/log"

	bs, err := json.Marshal(opts)
	if err != nil {
		log.Fatalf("failed to marshal data %+v", err)

	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(bs))
	if err != nil {
		log.Fatalf("failed to create req %+v", err)
	}

	req.SetBasicAuth("", token)
	client := http.Client{}
	_, err = client.Do(req)
	if err != nil {
		log.Fatalf("failed to do %+v", err)
	}
}
