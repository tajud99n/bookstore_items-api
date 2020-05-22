package items

import (
	"errors"

	"github.com/tajud99n/bookstore_items-api/clients/elasticsearch"

	rest_errors "github.com/tajud99n/bookstore_utils-go/errors"
)

const (
	indexItem = "items"
)

func (i *Item) Save() errors.RestErr {
	result, err := elasticsearch.Client.Index(indexItem, i)
	if err != nil {
		return rest_errors.NewInternalServerError("error when trying to save item", errors.New("database error"))
	}
	i.Id = result.Id
	return nil
}
