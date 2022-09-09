package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/PreetSIngh8929/bookstore-oauth-go/oauth"
	"github.com/PreetSIngh8929/bookstore_items-api/domain/items"
	"github.com/PreetSIngh8929/bookstore_items-api/domain/queries"
	"github.com/PreetSIngh8929/bookstore_items-api/services"
	"github.com/PreetSIngh8929/bookstore_items-api/utils/http_utils"
	"github.com/PreetSIngh8929/boookstore_utils-go/rest_errors"
	"github.com/gorilla/mux"
)

var (
	ItemsController itemsControllerInterface = &itemsController{}
)

type itemsControllerInterface interface {
	Create(w http.ResponseWriter, r *http.Request)
	Get(w http.ResponseWriter, r *http.Request)
	Search(w http.ResponseWriter, r *http.Request)
}
type itemsController struct {
}

func (c *itemsController) Create(w http.ResponseWriter, r *http.Request) {
	if err := oauth.AuthenticatRequest(r); err != nil {
		// http_utils.ResponseError(w, err)
		http_utils.ResponseJson(w, err.Status, err)
		return
	}
	sellerId := oauth.GetCallerId(r)
	if sellerId == 0 {
		respErr := rest_errors.NewUnauthorizedError("invalid request body")
		http_utils.ResponseError(w, respErr)
		return
	}

	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		respErr := rest_errors.NewBadRequestError("invalid request body")
		http_utils.ResponseError(w, respErr)
		return
	}
	defer r.Body.Close()
	var itemRequest items.Item
	if err := json.Unmarshal(requestBody, &itemRequest); err != nil {
		respErr := rest_errors.NewBadRequestError("invalid json body")
		http_utils.ResponseError(w, respErr)
		fmt.Println(respErr)
		return
	}
	itemRequest.Seller = sellerId

	result, createErr := services.ItemsService.Create(itemRequest)

	if createErr != nil {
		http_utils.ResponseError(w, createErr)
		return
	}

	http_utils.ResponseJson(w, http.StatusCreated, result)
}

func (c *itemsController) Get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	itemId := strings.TrimSpace(vars["id"])
	item, err := services.ItemsService.Get(itemId)
	if err != nil {
		http_utils.ResponseError(w, err)
		return
	}
	http_utils.ResponseJson(w, http.StatusOK, item)
}
func (c *itemsController) Search(w http.ResponseWriter, r *http.Request) {
	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		apiErr := rest_errors.NewBadRequestError("invalid json body")
		http_utils.ResponseError(w, apiErr)
		return
	}
	defer r.Body.Close()
	var query queries.EsQuery
	if err := json.Unmarshal(bytes, &query); err != nil {
		apiErr := rest_errors.NewBadRequestError("invalid json body")
		http_utils.ResponseError(w, apiErr)
		return
	}

	items, searchErr := services.ItemsService.Search(query)
	if searchErr != nil {
		http_utils.ResponseError(w, searchErr)
		return
	}
	http_utils.ResponseJson(w, http.StatusOK, items)
}
