package services

import (
	"github.com/core-go/search"
	"github.com/core-go/service"
)

type UserService interface {
	search.SearchService
	service.GenericService
}
