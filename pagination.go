// Copyright 2020 CleverGo. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found
// in the LICENSE file.

package pagination

import (
	"net/http"
	"strconv"
)

var (
	PageParam          = "page"
	LimitParam         = "limit"
	MaxLimit     int64 = 1000
	DefaultPage  int64 = 1
	DefaultLimit int64 = 20
)

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

// Offset returns offset.
func (p *Pagination) Offset() int64 {
	return (p.Page - 1) * p.Limit
}

// PageCount returns the number of pages.
func (p *Pagination) PageCount() int64 {
	return (p.Total + p.Limit - 1) / p.Limit
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
