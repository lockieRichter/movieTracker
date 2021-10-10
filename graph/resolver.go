package graph

//go:generate go run github.com/99designs/gqlgen

import "github.com/lockieRichter/movieTracker/graph/model"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	movies []*model.Movie
}
