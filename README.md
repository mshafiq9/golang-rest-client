# golang-rest-client
An example of a GoLang rest client consuming an API

## Installation
Download and install GoLang if you don't have [https://go.dev/doc/install](https://go.dev/doc/install)

#### Prerequisite
This rest client example need a TraderMade API key to run it. You can [sign up](https://marketdata.tradermade.com/signup) on TraderMade for free and then copy the API key from your [dashboard](https://marketdata.tradermade.com/myAccount).

```sh
# export the API key env var in your shell
export TRADERMADE_API_KEY="<your_own_api_key>"
```

#### Checkout and Run
```sh
$ git clone git@github.com:mshafiq9/golang-rest-client.git
$ cd golang-rest-client
$ go run .
```

#### Output
Read the console output. It should be similiar to as below:

```sh
{
  "endpoint": "live",
  "quotes": [
    {
      "ask": 1.12885,
      "base_currency": "EUR",
      "bid": 1.12885,
      "mid": 1.12885,
      "quote_currency": "USD"
    },
    {
      "ask": 1.35369,
      "base_currency": "GBP",
      "bid": 1.35369,
      "mid": 1.35369,
      "quote_currency": "USD"
    }
  ],
  "requested_time": "Wed, 02 Feb 2022 09:20:40 GMT",
  "timestamp": 1643793640
}

endpoint live requested time Wed, 02 Feb 2022 09:20:40 GMT timestamp 1643793640
0
symbol EUR USD bid 1.12885 ask 1.12885 mid 1.12885
1
symbol GBP USD bid 1.35369 ask 1.35369 mid 1.35369
```

## References
https://medium.com/nerd-for-tech/your-first-golang-rest-api-client-287c8dc0961

https://marketdata.tradermade.com/docs/restful-api#golang
