package user

import (
	"context"
	"fmt"
	"math/rand"
)

type Resolver struct{}

// Query resolver
func (r *Resolver) Users(ctx context.Context) ([]User, error) {
	// Mock data
	users := []User{
		{ID: "1", Name: "John Doe", Email: "john@example.com"},
		{ID: "2", Name: "Jane Doe", Email: "jane@example.com"},
	}
	return users, nil
}

// Mutation resolver
func (r *Resolver) CreateUser(ctx context.Context, name string, email string) (User, error) {
	// Create user with mock data
	user := User{
		ID:    fmt.Sprintf("%d", rand.Int()), // Generate a random ID for simplicity
		Name:  name,
		Email: email,
	}
	// Here you would typically add the user to your database
	return user, nil
}
