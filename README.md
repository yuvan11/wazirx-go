# Wazirx-go ![example workflow](https://github.com/yuvan11/wazirx-go/actions/workflows/go.yml/badge.svg)
A go client library for [wazirx](https://wazirx.com/exchange) REST [APIs](https://docs.wazirx.com/)

<p align="center" >
<img  src="https://user-images.githubusercontent.com/49576526/158990567-53ceb4a8-8a75-4d15-b407-9ee9f21168c7.png" width="200" height="200">
</p>

## Description

[wazirx-go](https://github.com/yuvan11/wazirx-go) is a go library for Wazirx Rest API integration.

### Available APIs
- [x] Common API
    - [x] General API ----------------- [docs](https://docs.wazirx.com/#general-endpoints) ,          [code](https://github.com/yuvan11/wazirx-go/blob/master/wazirx-client/general-api.go)
     - [x] Server Time ------------ [docs](https://docs.wazirx.com/#check-server-time) , [code](https://github.com/yuvan11/wazirx-go/blob/28b29c563fc9bb1e4b38f57477b40aedce85d8e5/wazirx-client/general-api.go#L57)
     - [x] System Status ---------- [docs](https://docs.wazirx.com/#system-status) , [code](https://github.com/yuvan11/wazirx-go/blob/28b29c563fc9bb1e4b38f57477b40aedce85d8e5/wazirx-client/general-api.go#L80)
     - [x] Exchange Info ---------- [docs](https://docs.wazirx.com/#exchange-info) , [code](https://github.com/yuvan11/wazirx-go/blob/28b29c563fc9bb1e4b38f57477b40aedce85d8e5/wazirx-client/general-api.go#L133)
    - [x] Market Data API ------------- [docs](https://docs.wazirx.com/#market-data-endpoints)  
     - [x] Tickers  ------------------ [docs](https://docs.wazirx.com/#24hr-tickers-price-change-statistics), [code](https://github.com/yuvan11/wazirx-go/blob/master/wazirx-client/tickers.go)
     - [x] Ticker ------------------- [docs](https://docs.wazirx.com/#24hr-ticker-price-change-statistics), [code](https://github.com/yuvan11/wazirx-go/blob/master/wazirx-client/ticker.go)
     - [x] OrderBookDepth ------- [docs](https://docs.wazirx.com/#order-book), [code](https://github.com/yuvan11/wazirx-go/blob/master/wazirx-client/orderBookDepth.go)
     - [x] Recent Trades ----------- [docs](https://docs.wazirx.com/#recent-trades-list), [code](https://github.com/yuvan11/wazirx-go/blob/master/wazirx-client/recentTrades.go)
     - [x] Historical Trades   -------- [docs](https://docs.wazirx.com/#old-trade-lookup-market_data), [code](https://github.com/yuvan11/wazirx-go/blob/master/wazirx-client/HistoricalTrade.go)
- [ ] Trading Endpoints  ----------------  [docs](https://docs.wazirx.com/#trading-endpoints)
     - [ ] New Order  ------------------- [docs](https://docs.wazirx.com/#new-order-trade)
     - [ ] Test New Order  -------------- [docs](https://docs.wazirx.com/#test-new-order-trade)
     - [ ] Query Order  ----------------- [docs](https://docs.wazirx.com/#query-order-user_data)
     - [ ] Query Open Orders  ----------[docs](https://docs.wazirx.com/#current-open-orders-user_data)
     - [ ] Query All Orders  -------------[docs](https://docs.wazirx.com/#all-orders-user_data)
     - [ ] Cancel Order  -----------------[docs](https://docs.wazirx.com/#cancel-order-trade)
     - [ ] Cancel All Orders  ------------[docs](https://docs.wazirx.com/#cancel-all-open-orders-on-a-symbol-trade)
- [x] Account API ----------------------- [docs](https://docs.wazirx.com/#account-endpoints)
     - [x] Account Info ----------------- [docs](https://docs.wazirx.com/#account-information-user_data), [code](https://github.com/yuvan11/wazirx-go/blob/master/wazirx-client/accountInfo.go)
     - [x] Funds ------------------------ [docs](https://docs.wazirx.com/#fund-details-user_data), [code](https://github.com/yuvan11/wazirx-go/blob/master/wazirx-client/funds.go)
- [x] Websocket Auth Token API  ------ [docs](https://docs.wazirx.com/#websocket-auth-tokens)
     - [x] Create Auth Token  ---------- [docs](https://docs.wazirx.com/#create-auth-token), [code](https://github.com/yuvan11/wazirx-go/blob/master/wazirx-client/authToken.go)

- [ ] Websocket Market Streams ------    [docs](https://docs.wazirx.com/#websocket-market-streams)
     - [ ] Live Subscribing & Unsubscribing  ------------------- [docs](https://docs.wazirx.com/#new-order-trade)
     - [ ] Trade Streams  -------------- [docs](https://docs.wazirx.com/#trade-streams)
     - [ ] All Market Tickers Stream  ----------------- [docs](https://docs.wazirx.com/#all-market-tickers-stream)
     - [ ] Depth Stream  ----------[docs](https://docs.wazirx.com/#depth-stream)
     - [ ] Account Update  -------------[docs](https://docs.wazirx.com/#account-update)
     - [ ] Order Update  -----------------[docs](https://docs.wazirx.com/#order-update)
     - [ ] Trade Update  ------------[docs](https://docs.wazirx.com/#trade-update)
     
    
## Getting Started

<!-- ### Dependencies

* Describe any prerequisites, libraries, OS version, etc., needed before installing program.
* ex. Windows 10 -->

### Prerequisites
It is prerequisite to have Go installed on your machine. [Steps to install Go](https://golang.org/doc/install?download)

### API Key Setup

* Some API requires an API key to verify the client request. For API Key creation,check [here](https://docs.wazirx.com/#api-key-setup).
* Sign in to the wazirx account and go to the [settings](https://wazirx.com/settings/profile) and then click [API KEY MANAGER](https://wazirx.com/settings/keys) at the bottom.
* Here, you can create your api key :smiley:

![image](https://user-images.githubusercontent.com/49576526/158996724-19916547-00be-4509-817f-b3cad1752f7d.png)


### Installing
	
	git clone https://github.com/yuvan11/wazirx-go.git

### Structure
<p align="center">
<img  src="https://user-images.githubusercontent.com/49576526/158992492-3a0197d7-7e36-45c0-b63c-5dfa277a295c.png" width="400" height="250">
</p>


### Executing program

    cd wazirx-go/cmd
    go run main.go

#### Note:
Don't forget to add API_KEY and SECRET_KEY in the code [here](https://github.com/yuvan11/wazirx-go/blob/bcae82aa8a01d7ca1173ba425b54dcae6ef59023/wazirx-client/Endpoints/endpoints.go#L33), as these are expected by some APIs for code & test execution.
		Otherwise, error will be thrown
- #### Starting a client
        
``` go
// Initializing client for accessing all API endpoints
client := wazirx.Client{}
    
// Initializing context
ctx := context.TODO()
``` 

- #### General API
    - *Server Time* 
             
	     	// Server service initiates
        	fmt.Println(client.NewServerTimeService().FetchServerTime(ctx))

    - *System Status*
   
              // System status service initiates
	        fmt.Println(client.NewSystemStatusService().FetchSystemStatus(ctx))
	- *Exchange Info*
	
	         // Exchange status service initiates
            fmt.Println(client.NewExchangeInfoService().FetchExchangeInfo(ctx))

- #### Market Data API
    - *Tickers*  
    
    	     // Tickers service initiates
        	fmt.Println(client.NewTickersService().DisplayTickersData(ctx))
    - *Ticker*
    
       	       // Ticker service initiates
               fmt.Println(client.NewTickerService().SetSymbol("manausdt").TickerData(ctx))
     - *OrderBook Depth*
      
	        // Market depth service initiates
	        fmt.Println(client.NewOrderBookDepthService().SetLimit(10).SetSymbol("manausdt").DisplayOrdersBookDepth(ctx))
    - *Recent Trades*
    
	        // Market trade service initiates
        	fmt.Println(client.NewRecentTradesService().SetLimit(10).SetSymbol("manausdt").DisplayRecentMarketTrade(ctx))
    - *Historical Trades*
    
	        // Historical trade service initiates
            fmt.Println(client.NewHistoricalTradesService().SetSymbol("manausdt").DisplayHistoricalTrade(ctx))

- #### Websocket Auth Token API
    - *Create Auth Token* 
    
            // Generate auth token service initiates
            fmt.Println(client.NewAuthInfoService().CreateAuthToken(ctx))

- #### Account API
    - *Account Info* 
    
            // Account display service initiates
            fmt.Println(client.NewAccountInfoService().FetchAccountInfo(ctx))
    - *Funds*
    
             // Account funds service initiates
	        fmt.Println(client.NewFundInfoService().FetchFundsInfo(ctx))
        

     -  *Example to fetch asset price greater than "0.0" and its asset name*

        ```go
        for _, a := range *client.NewFundInfoService().FetchFundsInfo(ctx) {

            if a.Free > "0.00" {
                fmt.Println("asset", a.Asset, "free", a.Free)
            }

        }
        ```
<!--
## Help

Any advise for common problems or issues.
```
command to run if program contains helper info
```
-->



## Unit Testing 
	
-	For running unit tests, 
	
		cd wazrix-client
		go test -v or go test
![image](https://user-images.githubusercontent.com/49576526/159119199-bf969747-f710-4fd7-9107-d3dcbbde17ac.png)

## Contributing
For contribution, have a look at [contributing.md](https://github.com/yuvan11/wazirx-go/blob/master/CONTRIBUTING.md)

## License

This project is licensed under the [MIT](https://en.wikipedia.org/wiki/MIT_License) License - see the [LICENSE](https://github.com/yuvan11/wazirx-go/blob/master/LICENSE) file for details

## Authors

Yuvaraj [@yuvan11](https://twitter.com/yuvaraj_11_)

<!--
## Version History

* 0.2
    * Various bug fixes and optimizations
    * See [commit change]() or See [release history]()
* 0.1
    * Initial Release


-->
