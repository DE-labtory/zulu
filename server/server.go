package server

import (
	"github.com/DE-labtory/zulu/account"
	"github.com/DE-labtory/zulu/keychain"
	"github.com/DE-labtory/zulu/keychain/leveldb"
	"github.com/DE-labtory/zulu/server/api"
	"github.com/gin-gonic/gin"
)

func New() *gin.Engine {
	resolver := account.NewResolver()
	walletApi := api.NewWallet(
		resolver,
		keychain.NewKeyGenerator(),
		leveldb.NewKeyStore(".keystore"),
	)
	coinApi := api.NewCoin(resolver)

	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	// /wallets
	r.POST("/wallets", walletApi.CreateWallet)
	r.GET("/wallets", walletApi.GetWallets)
	r.GET("/wallets/:id", walletApi.GetWallet)
	r.POST("/wallets/:id/transfers", walletApi.Transfer)
	// /coins
	r.GET("/coins", coinApi.GetCoins)
	r.GET("/coins/:id", coinApi.GetCoin)
	return r
}
