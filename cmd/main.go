package main

import (
	"context"
	"fmt"
	wazirx "wazirx-go/wazirx-client"
)

func main() {

	// Initializing client for accessing all API endpoints
	client := wazirx.Client{}
	// Initializing context
	ctx := context.TODO()

	// Public API Services init collection

	// Ping service initiates
	fmt.Println(client.NewPingService().FetchPingStatus(ctx))

	// Server service initiates
	fmt.Println(client.NewServerTimeService().FetchServerTime(ctx))

	// System status service initiates
	fmt.Println(client.NewSystemStatusService().FetchServerStatus(ctx))

	// Exchange status service initiates
	fmt.Println(client.NewExchangeInfoService().FetchExchangeInfo(ctx))

	// Tickers service initiates
	fmt.Println(client.NewTickersService().DisplayTickersData(ctx))

	// Ticker service initiates
	fmt.Println(client.NewTickerService().SetSymbol("manausdt").TickerData(ctx))

	// Market depth service initiates
	fmt.Println(client.NewOrderBookDepthService().SetLimit(10).SetSymbol("manausdt").DisplayOrdersBookDepth(ctx))
	// Market trade service initiates
	fmt.Println(client.NewRecentTradesService().SetLimit(10).SetSymbol("manausdt").DisplayRecentMarketTrade(ctx))

	// Historical trade service initiates
	fmt.Println(client.NewHistoricalTradesService().SetSymbol("manausdt").DisplayHistoricalTrade(ctx))

	// Account  API Services init collection

	// Account display service initiates
	fmt.Println(client.NewAccountInfoService().FetchAccountInfo(ctx))

	// Account funds service initiates
	fmt.Println(client.NewFundInfoService().FetchFundsInfo(ctx))

	/* EXAMPLE TO FETCH ASSET PRICE GREATER THAN "0.0" and asset name
	for _, a := range *client.NewFundInfoService().FetchFundsInfo(ctx) {

		if a.Free > "0.00" {
			fmt.Println("asset", a.Asset, "free", a.Free)
		}

	}

	*/

	// WebSocket-Client  API Services init collection

	// Generate auth token service initiates
	fmt.Println(client.NewAuthInfoService().CreateAuthToken(ctx))

}
