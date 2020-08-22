package api

import (
	"errors"

	"github.com/DE-labtory/zulu/account"
	"github.com/DE-labtory/zulu/interfaces"
	"github.com/DE-labtory/zulu/keychain"
	"github.com/DE-labtory/zulu/types"
	"github.com/gin-gonic/gin"
)

type WalletApi struct {
	resolver  *account.Resolver
	generator keychain.KeyGenerator
	store     keychain.Store
}

func NewWallet(resolver *account.Resolver) *WalletApi {
	return &WalletApi{
		resolver: resolver,
	}
}

func (w *WalletApi) CreateWallet(context *gin.Context) {
	key, err := w.generator.Generate()
	if err != nil {
		internalServerError(context, errors.New("failed to generate key pair"))
	}
	accounts, err := w.getAccounts(key)
	if err != nil {
		internalServerError(context, errors.New("failed to derive account"))
		return
	}

	if err := w.store.Store(key); err != nil {
		internalServerError(context, errors.New("failed to store key"))
		return
	}

	context.JSON(200, interfaces.CreateWalletResponse{
		Id:       key.ID,
		Accounts: accounts,
	})
}

func (w *WalletApi) GetWallet(context *gin.Context) {
	var pathParams struct {
		ID string `uri:"id" binding:"required"`
	}

	if err := context.ShouldBindUri(&pathParams); err != nil {
		badRequestError(context, errors.New("path variable :id does not exists"))
		return
	}

	key, err := w.store.Get(pathParams.ID)
	if err != nil {
		internalServerError(context, errors.New("failed to read key"))
		return
	}

	accounts, err := w.getAccounts(key)
	if err != nil {
		internalServerError(context, errors.New("failed to derive account"))
		return
	}

	context.JSON(200, interfaces.GetWalletResponse{
		Id:       key.ID,
		Accounts: accounts,
	})
}

func (w *WalletApi) GetWallets(context *gin.Context) {
	keys, err := w.store.GetAll()
	if err != nil {
		internalServerError(context, errors.New("failed to read key"))
		return
	}

	var response []interfaces.GetWalletResponse
	for _, key := range keys {
		accounts, err := w.getAccounts(key)
		if err != nil {
			internalServerError(context, errors.New("failed to derive account"))
			return
		}

		response = append(response, interfaces.GetWalletResponse{
			Id:       key.ID,
			Accounts: accounts,
		})
	}

	context.JSON(200, response)
}

func (w *WalletApi) getAccounts(key keychain.Key) ([]types.Account, error) {
	var accounts []types.Account
	for _, s := range w.resolver.GetAllServices() {
		a, err := s.DeriveAccount(key)
		if err != nil {
			return nil, err
		}
		accounts = append(accounts, a)
	}
	return accounts, nil
}

func (w *WalletApi) Transfer(context *gin.Context) {
	var requestBody interfaces.TransferRequest
	if err := context.ShouldBindJSON(&requestBody); err != nil {
		badRequestError(context, errors.New("failed to bind transfer requestBody body"))
		return
	}

	var pathParams struct {
		ID string `uri:"id" binding:"required"`
	}
	if err := context.ShouldBindUri(&pathParams); err != nil {
		badRequestError(context, errors.New("path variable :id does not exists"))
		return
	}

	service, err := w.resolver.Resolve(pathParams.ID)
	if err != nil {
		badRequestError(context, err)
		return
	}

	key, err := w.store.Get(pathParams.ID)
	if err != nil {
		internalServerError(context, errors.New("failed to read key"))
		return
	}

	transaction, err := service.Transfer(key, requestBody.To, requestBody.Amount)
	if err != nil {
		internalServerError(context, errors.New("failed to send transaction"))
		return
	}

	var response interfaces.TransferResponse = transaction
	context.JSON(200, response)
}
