# GoRest

## Description

Go REST is a simple boilerplate for rest API's using Fiber.

## Requirements

* [golangci-lint](https://golangci-lint.run)
* [goimports](https://pkg.go.dev/golang.org/x/tools/cmd/goimports)
* [goverreport](https://github.com/mcubik/goverreport)
* [pre-commit](https://pre-commit.com)
* [swag](https://github.com/swaggo/swag)

## Run

Before you run the project you need to install all deps mentioned above by running:
```bash
make install_deps
```

After installing all deps just run:
```bash
make run
```

The server will run on port 8080, http://localhost:8080/swagger/index.html.

## Swagger

```bash
make swagger
```
