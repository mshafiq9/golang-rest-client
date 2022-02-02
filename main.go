package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "log"
    "net/http"
    "os"
)

// diclaring the data struct to handle TraderMade API output format
type data struct {
    Endpoint       string                   `json:'endpoint'`
    Quotes         []map[string]interface{} `json:'quotes'`
    Requested_time string                   `json:'requested_time'`
    Timestamp      int32                    `json:'timestamp'`
}


func main() {
    currencies := "EURUSD,GBPUSD"

    // reading env variable for secret key
    api_key := os.Getenv("TRADERMADE_API_KEY")

    url := "https://marketdata.tradermade.com/api/v1/live?currency=" +
            currencies + "&api_key=" + api_key

    // hitting the API url and verifying no errors
    res, getErr := http.Get(url)

    if getErr != nil {
        log.Fatal(getErr)
    }

    // reading the response body and verifying no errors
    body, readErr := ioutil.ReadAll(res.Body)

    if readErr != nil {
        log.Fatal(readErr)
    }

    // printing response body
    fmt.Println(string(body))

    // diclaring the data struct variable
    data_obj := data{}

    // unmarshalling response body json object, into data_obj variable
    jsonErr := json.Unmarshal(body, &data_obj)

    if jsonErr != nil {
        log.Fatal(jsonErr)
    }

    // printing the API response from data_obj var
    fmt.Println("endpoint", data_obj.Endpoint, "requested time",
        data_obj.Requested_time, "timestamp", data_obj.Timestamp)

    // iterating and printing the data_obj.Quotes sub structure
    for key, value := range data_obj.Quotes {
        fmt.Println(key)
        fmt.Println("symbol", value["base_currency"], value["quote_currency"],
            "bid", value["bid"], "ask", value["ask"], "mid", value["mid"])
    }
}
