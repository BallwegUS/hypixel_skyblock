package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

const (
	api_url = "https://api.hypixel.net"
)

var apiClient = &http.Client{Timeout: 10 * time.Second}

func call_api(url string, target interface{}, params map[string]string) error {
	headers := make(map[string]string)
	return call_api_with_headers(url, target, params, headers)
}

func call_api_with_headers(url string, target interface{}, params map[string]string, headers map[string]string) error {
	request, err := http.NewRequest("GET", url, nil)

	for key, value := range headers {
		request.Header.Add(key, value)
	}

	q := request.URL.Query()
	for key, value := range params {
		q.Add(key, value)
	}
	request.URL.RawQuery = q.Encode()

	r, err := apiClient.Do(request)
	if err != nil {
		fmt.Println(err)
		return err
	}

	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}

func GetPlayerName(uuid string, api_key string) (error, *PlayerData) {
	headers := map[string]string{}
	params := map[string]string{}
	data := new(PlayerReturn)
	headers["API-Key"] = api_key

	params["uuid"] = uuid

	err := call_api_with_headers(api_url+"/player", data, params, headers)
	if err != nil {
		fmt.Println(err)
		return err, nil
	}

	if !data.Success {
		return fmt.Errorf("Failure to call API"), nil
	}

	return nil, &data.Player
}

func GetProducts() (int, map[string]Product) {
	data := new(BazaarData)

	err := call_api(api_url+"/skyblock/bazaar", data, map[string]string{})

	if err != nil {
		fmt.Println(err)
		return -1, nil
	}

	if !data.Success {
		fmt.Println("Error calling the api")
		return -1, nil
	}

	return data.LastUpdated, data.Products
}

func GetRecentlyEndedAuctions() (int, []EndedAuction) {
	data := new(EndedAuctionReturn)

	err := call_api(api_url+"/skyblock/ended_auctions", data, map[string]string{})

	if err != nil {
		fmt.Println(err)
		return -1, nil
	}

	if !data.Success {
		fmt.Println("Error calling the api")
		return -1, nil
	}

	return data.LastUpdated, data.Auctions
}

func GetAuctions() []AuctionData {
	semaphoreChan := make(chan struct{}, 5)
	resultsChan := make(chan *AuctionReturn)
	auctions := []AuctionData{}
	cache := AuctionCache{}

	err := Load("auctions", &cache)
	if err == nil && time.Now().Sub(cache.LastUpdated).Minutes() < 5 {
		return cache.Auctions
	}

	data := AuctionReturn{}
	params := map[string]string{}
	params["page"] = "0"
	err = call_api(api_url+"/skyblock/auctions", &data, params)
	if err != nil {
		panic(err)
	}

	if data.Success {
		auctions = append(auctions, data.Auctions...)
	}

	page := 1
	for page <= data.TotalPages {
		go func(p string) {
			semaphoreChan <- struct{}{}
			params := map[string]string{}
			params["page"] = p
			d := AuctionReturn{}
			err = call_api(api_url+"/skyblock/auctions", &d, params)
			if err != nil {
				fmt.Println(err)
				<-semaphoreChan
				return
			}
			resultsChan <- &d
			<-semaphoreChan
		}(strconv.FormatInt(int64(page), 10))
		page += 1
	}

	resultCount := 1
	for {
		resultCount += 1
		result := <-resultsChan
		if result.Success {
			auctions = append(auctions, result.Auctions...)
		}

		if resultCount == data.TotalPages {
			break
		}
	}

	cache = AuctionCache{}
	cache.LastUpdated = time.Now()
	cache.Auctions = auctions
	Save("auctions", cache)

	return auctions
}
