# Pagination
[![Build Status](https://travis-ci.org/clevergo/pagination.svg?branch=master)](https://travis-ci.org/clevergo/pagination)
[![Coverage Status](https://coveralls.io/repos/github/clevergo/pagination/badge.svg?branch=master)](https://coveralls.io/github/clevergo/pagination?branch=master)
[![GoDoc](https://img.shields.io/badge/godoc-reference-blue)](https://pkg.go.dev/github.com/clevergo/pagination)
[![Go Report Card](https://goreportcard.com/badge/github.com/clevergo/pagination)](https://goreportcard.com/report/github.com/clevergo/pagination)
[![Release](https://img.shields.io/github/release/clevergo/pagination.svg?style=flat-square)](https://github.com/clevergo/pagination/releases)

## Usage

```go
import "github.com/clevergo/pagination"
```

```go
// query entities from database.
func query(limit, offset int64) (total int64, es []entity, err error) {
    return
}

func index(w http.ResponseWriter, req *http.Request) {
    p := pagination.NewFromRequest(req)
    p.Total, p.Items, _ = query(p.Limit, p.Offset())
    data, _ := json.Marshal(p)
    w.Write(data)
}
```