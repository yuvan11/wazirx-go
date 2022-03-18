# Wazrix-go

A go client library for [wazrix](https://wazirx.com/exchange) REST [APIs](https://docs.wazirx.com/)

## Description

[Wazrix-go](https://github.com/yuvan11/wazirx-go) is a go library for Wazirx Rest API integration for spot exchange

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

### Installing

*  clone [repo](https://github.com/yuvan11/wazirx-go.git) 
        git clone https://github.com/yuvan11/wazirx-go.git


### Executing program

        cd wazirx-go/cmd
        go run main.go
        
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

## Authors



Yuvaraj [@yuvan11](https://twitter.com/yuvaraj_11_)

<!--
## Version History

* 0.2
    * Various bug fixes and optimizations
    * See [commit change]() or See [release history]()
* 0.1
    * Initial Release

## License

This project is licensed under the [NAME HERE] License - see the LICENSE.md file for details
-->
