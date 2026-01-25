package domain

import "buytun-backend/internal/domain/model"

var MIGRATION = []any{
	&model.User{},
	&model.Product{},
}
