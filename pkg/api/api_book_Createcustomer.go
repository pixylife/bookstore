package api

import (
	bookstore "bookstore/bookstore"
	arfcreatecustomer "bookstore/pkg/rule/ARF_Createcustomer"
	empty "github.com/golang/protobuf/ptypes/empty"
	"reflect"
)

import (
	"context"
)

func Createcustomer(ctx context.Context, book *bookstore.Book) (*empty.Empty, error) {

	result := arfcreatecustomer.ARF_Createcustomer(&book)
	if reflect.TypeOf(result) == reflect.TypeOf(bookstore.Book{}) || reflect.TypeOf(result) == reflect.TypeOf(&bookstore.Book{}) {
		return result, nil
	} else {
		return result, nil
	}
}
