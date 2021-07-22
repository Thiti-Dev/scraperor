package graph

import "github.com/Thiti-Dev/scraperor/postgres"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct{
	UsersRepository postgres.UsersRepository
}
