package parser

import (
	"encoding/json"
	"io"
	"log"
	"os"
)

type Data struct {
	Temperature float32 `json:"Temperature"`
	Humidity    float32 `json:"Humidity"`
}

func Parsing(date string) Data {

	content, err := os.Open("data_" + date + ".json")

	if err != nil {
		log.Fatal(err)
	}

	defer content.Close()

	byteRes, _ := io.ReadAll(content)

	var data Data

	json.Unmarshal(byteRes, &data)

	return data
}
