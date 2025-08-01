
```
go-kafka
├─ .air.toml
├─ Makefile
├─ README.md
├─ cmd
│  └─ server
│     └─ main.go
├─ deploy
│  └─ docker
│     ├─ Dockerfile
│     └─ docker-compose.yml
├─ go.mod
├─ go.sum
├─ internal
│  ├─ config
│  │  └─ config.go
│  ├─ dto
│  │  └─ user.go
│  ├─ factory
│  │  ├─ auth_factory.go
│  │  └─ factory.go
│  ├─ handler
│  │  └─ auth.go
│  ├─ infrastructure
│  │  └─ db
│  │     └─ db.go
│  ├─ middleware
│  ├─ model
│  │  └─ user.go
│  ├─ repository
│  │  └─ user.go
│  ├─ routes
│  │  ├─ auth.go
│  │  └─ routes.go
│  └─ usecase
│     └─ auth.go
├─ migrations
│  ├─ 000001_create_users_table.down.sql
│  └─ 000001_create_users_table.up.sql
└─ pkg
   ├─ graceful
   │  └─ graceful.go
   ├─ logger
   │  ├─ logfile.go
   │  └─ logger.go
   ├─ response
   │  └─ response.go
   └─ utils
      ├─ hash_password.go
      └─ jwt.go

```