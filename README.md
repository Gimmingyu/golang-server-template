# golang-server-template

## Directory

```shell
├── Dockerfile
├── Makefile
├── build
    └── build.md
├── cmd
    ├── cmd.md
    └── main.go
├── internal
    ├── config
        ├── aws.go
        ├── config.go
        ├── database.go
        └── env.go
    ├── domain
        ├── auth
            ├── auth.go
            ├── auth.handler.go
            └── auth.router.go
        ├── item
            ├── item.go
            ├── item.handler.go
            └── item.router.go
        └── user
            ├── user.go
            ├── user.handler.go
            └── user.router.go
    ├── infra
        ├── filter
            └── filter.go
        ├── interceptor
            └── interceptor.go
        ├── lifecycle
            └── lifecycle.go
        └── middleware
            └── middleware.go
    └── pkg
└── pkg

15 directories, 31 files
