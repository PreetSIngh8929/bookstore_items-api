package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/PreetSIngh8929/bookstore-oauth-go/oauth"
	"github.com/PreetSIngh8929/bookstore_items-api/domain/items"
	"github.com/PreetSIngh8929/bookstore_items-api/services"
	"github.com/PreetSIngh8929/bookstore_items-api/utils/http_utils"
	"github.com/PreetSIngh8929/boookstore_utils-go/rest_errors"
)

var (
	ItemsController itemsControllerInterface = &itemsController{}
)

type itemsControllerInterface interface {
	Create(w http.ResponseWriter, r *http.Request)
	Get(w http.ResponseWriter, r *http.Request)
}
type itemsController struct {
}

func (c *itemsController) Create(w http.ResponseWriter, r *http.Request) {
	if err := oauth.AuthenticatRequest(r); err != nil {
		// http_utils.ResponseError(w, err)
		http_utils.ResponseJson(w, err.Status, err)
		return
	}
	var itemRequest items.Item

	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		respErr := rest_errors.NewBadRequestError("invalid request body")
		http_utils.ResponseError(w, *respErr)
		return
	}
	defer r.Body.Close()

	if err := json.Unmarshal(requestBody, &itemRequest); err != nil {
		respErr := rest_errors.NewBadRequestError("invalid json body")
		http_utils.ResponseError(w, *respErr)
		return
	}
	itemRequest.Seller = oauth.GetCallerId(r)

	result, createErr := services.ItemsService.Create(itemRequest)

	if createErr != nil {
		http_utils.ResponseError(w, *createErr)
		return
	}

	http_utils.ResponseJson(w, http.StatusCreated, result)
}

func (c *itemsController) Get(w http.ResponseWriter, r *http.Request) {

}
