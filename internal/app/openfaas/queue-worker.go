package exporter

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func GetQueueworkerTotalMessage(nexusurl string) int64 {

	var url string
	// curl "http://10.42.1.226:8222/streaming/channelsz?channel=faas-request&subs=1" -s | jq ".subscriptions[].last_sent" | sort -n"
	// curl "http://10.42.1.226:8222/streaming/serverz" -s | jq ".total_msgs

	url = nexusurl

	fmt.Printf("your request url : %s\n", url)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		// handle err
	}
	req.Header.Set("Accept", "application/json")

	resp, err := http.DefaultClient.Do(req)
	defer resp.Body.Close()
	if err != nil {
		log.Println("request failed")
	}
	responseData, err := ioutil.ReadAll(resp.Body)
	//log.Println(string(responseData))
	temp_total_msgs := JsonParseInt64(string(responseData), "total_msgs")
	//total_msgs, err := strconv.Atoi(temp_total_msgs)
	log.Printf("get total_msgs : %d \n", temp_total_msgs)
	return temp_total_msgs
}

func GetQueueworkerLastsent(url1 string, url2 string) int64 {

	var url string
	url = url1
	log.Printf("your request url : %s\n", url)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		// handle err
	}
	req.Header.Set("Accept", "application/json")

	resp, err := http.DefaultClient.Do(req)
	defer resp.Body.Close()
	if err != nil {
		log.Println("request failed")
	}
	responseData, err := ioutil.ReadAll(resp.Body)
	//log.Println(string(responseData))
	temp_subscriptions := JsonParseInt64(string(responseData), "subscriptions")

	// curl "http://10.42.1.226:8222/streaming/channelsz?channel=faas-request&subs=1" -s | jq ".subscriptions[].last_sent" | sort -n"
	// curl "http://10.42.1.226:8222/streaming/serverz" -s | jq ".total_msgs
	url = url2
	log.Printf("your request url : %s\n", url)
	req2, err := http.NewRequest("GET", url, nil)
	if err != nil {
		// handle err
	}
	req.Header.Set("Accept", "application/json")

	resp, err = http.DefaultClient.Do(req2)
	defer resp.Body.Close()
	if err != nil {
		log.Println("request failed")
	}
	responseData, err = ioutil.ReadAll(resp.Body)
	//log.Println("----------srart of responseData-----------")
	//log.Println(string(responseData))

	temp_subscriptions_last_sent := JsonParseArray(string(responseData), int(temp_subscriptions))
	log.Printf("get subscriptions_last_sent : %d \n", temp_subscriptions_last_sent)
	return temp_subscriptions_last_sent
}
