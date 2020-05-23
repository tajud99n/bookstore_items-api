package services

import (
	"net/http"

	"github.com/tajud99n/bookstore_items-api/domain/items"
	"github.com/tajud99n/bookstore_utils-go/errors"
)

var (
	ItemsService ItemServiceInterface = &itemsService{}
)

type ItemServiceInterface interface {
	Create(items.Item) (*items.Item, errors.RestErr)
	Get(string) (*items.Item, errors.RestErr)
}

type itemsService struct{}

func (s *itemsService) Create(itemRequest items.Item) (*items.Item, errors.RestErr) {
	if err := itemRequest.Save(); err != nil {
		return nil, err
	}
	return &item, nil
}

func (s *itemsService) Get(id string) (*items.Item, errors.RestErr) {
	item := items.Item{Id: id}
	if err := item.Get(); err != nil {
		return nil, err
	}
	return &item, nil
}
