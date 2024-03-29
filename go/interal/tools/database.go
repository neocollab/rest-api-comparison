package tools

type LoginDetails struct {
	AuthToken string
	userName  string
}

type CoinBalanceDetails struct {
	Balance int64
}

type DatabaseInterface interface {
	GetUserLoginDetails(username string) *LoginDetails
	GetUserCoins(username string) *CoinBalanceDetails
	SetupDatabase() error
}
