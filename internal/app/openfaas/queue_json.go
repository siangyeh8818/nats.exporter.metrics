package exporter

import (
	"log"
	"strconv"

	"github.com/buger/jsonparser"
)

type JsonType struct {
	Array []string
}

func JsonParseInt64(jsondata string, pattern string) int64 {
	var data []byte = []byte(jsondata)

	total_msgs, err := jsonparser.GetInt(data, pattern)
	if err != nil {

		log.Println(err)
	}

	return total_msgs
}

func JsonParseArray(jsondata string, array_size int) int64 {
	var data []byte = []byte(jsondata)
	var max_last_sent int64

	for i := 0; i < array_size; i++ {
		var temp string
		temp = "[" + strconv.Itoa(i) + "]"
		last_sent, err := jsonparser.GetInt(data, "subscriptions", temp, "last_sent")
		if last_sent > max_last_sent {
			max_last_sent = last_sent
		}
		if err != nil {
			log.Println(err)
		}
	}
	return max_last_sent
}
