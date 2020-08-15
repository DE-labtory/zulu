package api

import (
	"errors"
	"github.com/DE-labtory/zulu/account"
	"github.com/DE-labtory/zulu/types"
	"github.com/gin-gonic/gin"
)

type CoinApi struct {
	resolver *account.Resolver
}

func NewCoin(resolver *account.Resolver) *CoinApi {
	return &CoinApi{
		resolver: resolver,
	}
}

func (c *CoinApi) GetCoins(context *gin.Context) {
	var coins []types.Coin
	for _, s := range c.resolver.GetAllServices() {
		coins = append(coins, s.GetInfo())
	}
	context.JSON(200, coins)
}

func (c *CoinApi) GetCoin(context *gin.Context) {
	var request struct {
		ID string `uri:"id" binding:"required"`
	}

	if err := context.ShouldBindUri(&request); err != nil {
		badrequestError(context, errors.New("path variable :id does not exists"))
		return
	}

	service, err := c.resolver.Resolve(request.ID)
	if err != nil{
		badrequestError(context, err)
		return
	}
	context.JSON(200, service.GetInfo())
}
