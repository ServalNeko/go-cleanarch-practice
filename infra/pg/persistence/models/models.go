package models

import "github.com/uptrace/bun"

type User struct {
	bun.BaseModel `bun:"users"`
	Id            string `bun:"id,pk"`
	Name          string `bun:"name,notnull"`
	UserType      int    `bun:"type,notnull"`
}

type Circle struct {
	bun.BaseModel `bun:"circles"`
	Id            string `bun:"id,pk"`
	Name          string `bun:"name,notnull"`
	QwnerId       string `bun:"owner_id"`
	Owner         *User  `bun:"rel:belongs-to,join:owner_id=id"`
	Members       []User `bun:"m2m:user_circles,join:User=Circle"`
}

type UserToCircle struct {
	bun.BaseModel `bun:"user_circles"`
	UserId        string  `bun:"user_id,pk"`
	User          *User   `bun:"rel:belongs-to,join:user_id=id"`
	CircleId      string  `bun:"circle_id,pk"`
	Circle        *Circle `bun:"rel:belongs-to,join:circle_id=id"`
}
