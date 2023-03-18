package repository

import (
	"github.com/Willw99/Project-Frontier-Go-Edition/models/request"
)

type RequestStateRepository interface {
	CreateState(request.Request) error
	GetAll()
}
