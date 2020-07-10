// Copyright 2020 CleverGo. All rights reserved.
// Use of this source code is governed by a MIT style license that can be found
// in the LICENSE file.

package pagination

import (
	"net/http"
	"strconv"

	"clevergo.tech/clevergo"
)

var (
	// PageParam defines the name of page parameter.
	PageParam = "page"
	// LimitParam defines the name of limit parameter.
	LimitParam = "limit"
	// MaxLimit defines a maximum number of limitation.
	MaxLimit int64 = 1000
	// DefaultPage defines the default value of page.
	DefaultPage int64 = 1
	// DefaultLimit defines the default value of limit.
	DefaultLimit int64 = 20
)

// Pagination represents a paginated list of data items.
type Pagination struct {
	Page  int64       `json:"page"`
	Limit int64       `json:"limit"`
	Total int64       `json:"total"`
	Items interface{} `json:"items"`
}

// New returns a pagination with the given page and limit.
func New(page, limit int64) *Pagination {
	return &Pagination{
		Page:  page,
		Limit: limit,
	}
}

// NewFromRequest returns a pagination from the given HTTP request.
func NewFromRequest(req *http.Request) *Pagination {
	query := req.URL.Query()
	return New(
		parsePage(query.Get(PageParam)),
		parseLimit(query.Get(LimitParam)),
	)
}

// NewFromContext returns a pagination from the given context.
func NewFromContext(ctx *clevergo.Context) *Pagination {
	return New(
		parsePage(ctx.QueryParam(PageParam)),
		parseLimit(ctx.QueryParam(LimitParam)),
	)
}

// Upage returns the unsigned page.
func (p *Pagination) Upage() uint64 {
	return uint64(p.Page)
}

// Ulimit returns the unsigned limit.
func (p *Pagination) Ulimit() uint64 {
	return uint64(p.Limit)
}

// Offset returns the offset.
func (p *Pagination) Offset() int64 {
	return (p.Page - 1) * p.Limit
}

// Uoffset returns the unsigned offset.
func (p *Pagination) Uoffset() uint64 {
	return uint64(p.Offset())
}

// PageCount returns the number of pages.
func (p *Pagination) PageCount() int64 {
	return (p.Total + p.Limit - 1) / p.Limit
}

// Pages returns a set of page numbers for rendering pagination.
// Zero means dots.
func (p *Pagination) Pages() (pages []int64) {
	total := p.PageCount()
	min := p.Page - 2
	max := p.Page + 2
	dot := false
	for i := int64(1); i <= total; i++ {
		if i == 1 || i == total || (min <= i && i <= max) {
			pages = append(pages, i)
			dot = false
		} else if !dot {
			pages = append(pages, 0)
			dot = true
		}
	}
	return
}

func parsePage(s string) int64 {
	v, err := strconv.ParseInt(s, 10, 64)
	if err == nil && v > 0 {
		return v
	}

	return DefaultPage
}

func parseLimit(s string) int64 {
	v, err := strconv.ParseInt(s, 10, 64)
	if err == nil && v > 0 {
		if v > MaxLimit {
			v = MaxLimit
		}
		return v
	}

	return DefaultLimit
}
