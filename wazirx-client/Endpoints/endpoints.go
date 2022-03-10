package Endpoints

const (

	// Public API Endpoints
	BASE_URL = "https://api.wazirx.com/sapi/"
	PING     = "v1/ping"
	TIME     = "v1/time"

	SYSTEM_STATUS = "v1/systemStatus"

	EXCHAGE_INFO = "v1/exchangeInfo"

	TICKERS           = "v1/tickers/24hr"
	TICKER            = "v1/ticker/24hr"
	MARKET_DEPTH      = "v1/depth"
	MARKET_TRADES     = "v1/trades"
	HISTORICAL_TRADES = "v1/historicalTrades"

	// Account API Endpoints

	ORDER      = "v1/order"
	TEST_ORDER = "v1/order/test"
	ALL_ORDERS = "v1/allOrders"
	ACCOUNT    = "v1/account"
	FUNDS      = "v1/funds"

	// WebSocket-Client
	CREATE_AUTH_TOKEN = "v1/create_auth_token"

	//API KEYS

	API_KEY    = "YOUR_API_KEY"
	SECRET_KEY = "YOUR_SECRET_KEY"
)
