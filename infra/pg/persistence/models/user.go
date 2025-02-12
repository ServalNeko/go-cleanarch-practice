package models

import "github.com/uptrace/bun"

type UserDataModel struct {
	bun.BaseModel `bun:"users"`
	Id            string `bun:"id,pk"`
	Name          string `bun:"name,notnull"`
	UserType      int    `bun:"type,notnull"`
}
