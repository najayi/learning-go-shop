package resolver

import (
	"context"
	"errors"

	"github.com/najayi/learning-go-shop/internal/utils"
)

var (
	ErrUnauthorized = errors.New("unauthorized")
)

const (
	adminRole = "admin"
)

// GetUserIDFromContext function to extract user info from GraphQL context
func GetUserIDFromContext(ctx context.Context) (uint, error) {
	userID := ctx.Value(utils.UserIDKey)
	if userID == nil {
		return 0, ErrUnauthorized
	}

	if id, ok := userID.(uint); ok {
		return id, nil
	}

	return 0, ErrUnauthorized
}

func GetUserRoleFromContext(ctx context.Context) (string, error) {
	userRole := ctx.Value(utils.UserRoleKey)
	if userRole == nil {
		return "", ErrUnauthorized
	}

	if userRole, ok := userRole.(string); ok {
		return userRole, nil
	}

	return "", ErrUnauthorized
}

func IsAdminFromContext(ctx context.Context) bool {
	role, err := GetUserRoleFromContext(ctx)
	if err != nil {
		return false
	}

	return role == adminRole
}

func getPagingNumbers(page, limit *int) (pageNumber, pageLimit int) {
	var p, l = 0, 0

	if page != nil {
		p = *page
	}

	if limit != nil {
		l = *limit
	}

	if p <= 0 {
		p = 1
	}
	if l <= 0 {
		l = 10
	}
	return p, l
}
