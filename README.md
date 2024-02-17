# Fiber API

<div align="center">

![GitHub go.mod Go version (subdirectory of monorepo)](https://img.shields.io/github/go-mod/go-version/granitebps/fiber-api)
[![Go Report Card](https://goreportcard.com/badge/github.com/granitebps/fiber-api)](https://goreportcard.com/report/github.com/granitebps/fiber-api)
![GitHub Workflow Status (with event)](https://img.shields.io/github/actions/workflow/status/granitebps/fiber-api/test.yml)
![GitHub](https://img.shields.io/github/license/granitebps/fiber-api)

</div>

## Quick Start
1. Clone the repo.
2. Copy `.env.example` file to `.env` file.
3. Fill the `.env` file.
4. Install docker (optional) and some tools:
   - https://github.com/swaggo/swag for auto-generating Swagger API docs
   - https://github.com/securego/gosec for checking Go security issues
   - https://github.com/golangci/golangci-lint for checking Go linter issues
   - https://github.com/cosmtrek/air for run the app with hot reload
   - https://github.com/google/wire for generate dependency injection file
6. Run `make air` to run the app using hot reload or you can use docker.
7. Go to API Docs page (Swagger): http://localhost:8000/swagger/index.html
8. Check `Makefile` file to see available command.

## Template Structure

### ./assets
**Folder with app assets.** You can store your app assets in this folder. Such as logo, image, favicon, etc.

### ./config
**Folder with app configuration.** This folder contains all app and package config file.

### ./docs
**Folder with API Documentation.** This directory contains config files for auto-generated API Docs by Swagger.

### ./pkg
**Folder with reusable code.** This folder hold code that will be used everywhere. Code in here will not depend with other code (maybe a little bit).
- `./pkg/constants` folder for app const so no hardcode variable
- `./pkg/core` folder for app core utility function, like setup database, setup cache, setup log, setup http client, and setup newrelic
- `./pkg/utils` folder for custom utility function

### ./scheduler
This folder hold task scheduling code. If you want to use task scheduling like crontab but using the app, you can write you code in here and register the task.

### ./src
**Folder with application functionality code.** This folder will be the heart of the code. This folder similar to `internal` folder in other go project but can be used in other place if needed. This little bit follow the SOLID principal.
- `./src/handler` folder for functional handlers (used in routes)
- `./src/middleware` folder for app middleware
- `./src/model` folder for define your database table structure
- `./src/repository` folder for querying to database
- `./src/request` folder for define api request payload
- `./src/route` folder for describe routes of your project
- `./src/service` folder for business logic
- `./src/transformer` folder for define api response

### ./storage
**Folder with storing app misc files.** App log will be stored in this folder under `logs` folder. You can also store `file-based-storage` file in here.

## References
- [Standard Go Project Layout](https://github.com/golang-standards/project-layout?tab=readme-ov-file)
- [Go - The Ultimate Folder Structure](https://gist.github.com/ayoubzulfiqar/9f1a34049332711fddd4d4b2bfd46096)

## Notes
- Use `<Ctx>.UserContext()` to get context from fiber and parsing it down the line. That context will be used in all code that need or use context and registered it to newrelic.
- Use `fibersentry.GetHubFromContext(<Ctx>)` to get sentry instance and use that instance to send exception or message to sentry `<*sentry.Hub>.CaptureException(<error>)` . See https://docs.gofiber.io/contrib/fibersentry_v1.x.x/fibersentry/#usage
- Use https://github.com/ansel1/merry to wrap error so you can pass http error code and error message.

## TODO
- [x] Add JWT
- [x] Add scheduler
- [ ] Add event listener/queue
- [ ] Add test
- [x] Add github action
- [x] Add example
- [x] Update readme
- [x] Add validation
- [x] Add error wrapping
- [ ] Add migration
- [ ] Add cmd
- [ ] Add command