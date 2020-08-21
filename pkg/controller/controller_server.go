package controller

import (
	bookstore "bookstore/bookstore"
	api "bookstore/pkg/api"
	empty "github.com/golang/protobuf/ptypes/empty"
)

import (
	"context"
)

type BookstoreServer struct{}

func (s *BookstoreServer) Createcustomer(ctx context.Context, parameters *bookstore.CreatecustomerParameters) (*empty.Empty, error) {

	result, err := api.Createcustomer(ctx, parameters.Book)
	return result, err
}
func (s *BookstoreServer) Updatecustomer(ctx context.Context, parameters *bookstore.UpdatecustomerParameters) (*bookstore.Book, error) {

	result, err := api.Updatecustomer(ctx, parameters.Book)
	return result, err
}
