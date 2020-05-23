package controllers

import (
	"github.com/tajud99n/bookstore_items-api/utils/http_utils"
	"net/http"

	"github.com/tajud99n/bookstore_items-api/domain/items"
	"github.com/tajud99n/bookstore_items-api/services"
	"github.com/tajud99n/bookstore_oauth-go/oauth"
	"github.com/tajud99n/bookstore_utils-go/errors"
)

var (
	ItemsController itemsControllerInterface = &itemsController{}
)

type itemsControllerInterface interface {
	Create(w http.ResponseWriter, r *http.Request)
	Get(w http.ResponseWriter, r *http.Request)
}

type itemsController struct {}

func (c *itemsController)Create(w http.ResponseWriter, r *http.Request) {
	if err := oauth.AuthenticateRequest(r); err != nil {
		http_utils.RespondError(w, err)
		return
	}
	
	requestBody, err := ioutil.ReadAll(r.body)
	if err != nil {
		respErr := errors.NewBadRequest("invalid request body")
		http_utils.RespondError(w, err)
		return
	}
	defer r.Body.Close()

	var itemRequest items.Item
	if err := json.Unmarshal(requestBody, &itemRequest); err != nil {
		respErr := errors.RestErr.NewBadRequest("invalid json body")
		http_utils.RespondError(w, err)
		return
	}

	itemRequest.Seller = Seller: oauth.GetCallerId(r),

	result, createErr := services.ItemsService.Create(itemRequest)
	if err != nil {
		http_utils.RespondError(w, createErr)
		return
	}
	http_utils.RespondJSON(w, http.StatusCreated, result)
}

func (c *itemsController)Get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	itemId := strings.TrimSpace(vars["id"])

	item, err := services.ItemsService.Get(itemId)
	if err != nil {
		http_utils.RespondError(w, err)
		return
	}
	http_utils.RespondJSON(w, http.StatusCreated, item)
}
