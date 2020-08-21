package api

import (
	bookstore "bookstore/bookstore"
	arfupdatecustomer "bookstore/pkg/rule/ARF_Updatecustomer"
	"reflect"
)

import (
	"context"
)

func Updatecustomer(ctx context.Context, book *bookstore.Book) (*bookstore.Book, error) {

	result := arfupdatecustomer.ARF_Updatecustomer(&book)
	if reflect.TypeOf(result) == reflect.TypeOf(bookstore.Book{}) || reflect.TypeOf(result) == reflect.TypeOf(&bookstore.Book{}) {
		return result, nil
	} else {
		return result, nil
	}
}
