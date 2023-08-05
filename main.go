package gopher_orm

import (
	gopher "github.com/AbbasSrour/gopher-orm/schema"
)

type User struct {
	gopher.Entity
	ID int64
}

func (*User) Declare() *gopher.Schema {
	return &gopher.Schema{
		Fields: []*gopher.Field{
			gopher.Column("id").Int(),
		},
	}
}
