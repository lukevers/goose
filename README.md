# goose

### Quick Start

```
go get github.com/go-martini/martini
go get github.com/russross/meddler
```

```
go run server.go
```

**NOTE:** if you are trying to just build it, running `go build` will result in errors due to local packages. Local packages are used because this is intended for anyone to be able to use and not be dependant on packages that will be changed locally for each project. Running `go build server.go` will give you the results you are looking for.
