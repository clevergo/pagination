# Pagination
[![Build Status](https://img.shields.io/travis/clevergo/pagination?style=for-the-badge)](https://travis-ci.org/clevergo/pagination)
[![Coverage Status](https://img.shields.io/coveralls/github/clevergo/pagination?style=for-the-badge)](https://coveralls.io/github/clevergo/pagination?branch=master)
[![Go.Dev reference](https://img.shields.io/badge/go.dev-reference-blue?logo=go&logoColor=white&style=for-the-badge)](https://pkg.go.dev/clevergo.tech/pagination?tab=doc)
[![Go Report Card](https://goreportcard.com/badge/github.com/clevergo/pagination?style=for-the-badge)](https://goreportcard.com/report/github.com/clevergo/pagination)
[![Release](https://img.shields.io/github/release/clevergo/pagination.svg?style=for-the-badge)](https://github.com/clevergo/pagination/releases)
[![Downloads](https://img.shields.io/endpoint?url=https://pkg.clevergo.tech/api/badges/downloads/month/clevergo.tech/pagination&style=for-the-badge)](https://pkg.clevergo.tech/)

## Usage

```go
import "clevergo.tech/pagination"
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
