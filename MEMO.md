```
├── MEMO.md
├── domain
│   ├── circles
│   └── users
│       └── User.go
├── go.mod
├── go.sum
├── infra
│   └── pg
│       ├── db.go
│       └── persistence
│           ├── models
│           │   └── user.go
│           └── users
│               └── repository.go
├── main.go
├── usecase
│   └── users
│       ├── IUserRepository.go
│       ├── interactor
│       │   ├── delete.go
│       │   ├── get.go
│       │   ├── get_all.go
│       │   ├── register.go
│       │   └── update.go
│       └── ports
│           ├── delete.go
│           ├── get.go
│           ├── get_all.go
│           ├── register.go
│           ├── update.go
│           └── user.go
└── web
    ├── handlers
    │   └── users.go
    └── models
        └── users.go
```