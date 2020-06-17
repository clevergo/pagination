// Copyright 2020 CleverGo. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found
// in the LICENSE file.

package pagination

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"clevergo.tech/clevergo"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	cases := []struct {
		page  int64
		limit int64
	}{
		{1, 10},
		{2, 10},
		{1, 20},
	}
	for _, test := range cases {
		p := New(test.page, test.limit)
		assert.Equal(t, test.page, p.Page)
		assert.Equal(t, test.limit, p.Limit)
	}
}

type requestTestCase struct {
	req   *http.Request
	page  int64
	limit int64
}

func requestTestCases() []requestTestCase {
	return []requestTestCase{
		{httptest.NewRequest(http.MethodGet, "/", nil), DefaultPage, DefaultLimit},
		{httptest.NewRequest(http.MethodGet, "/?page=1", nil), 1, DefaultLimit},
		{httptest.NewRequest(http.MethodGet, "/?limit=10", nil), DefaultPage, 10},
		{httptest.NewRequest(http.MethodGet, "/?page=2&limit=10", nil), 2, 10},
		{httptest.NewRequest(http.MethodGet, "/?page=-1&limit=10", nil), DefaultPage, 10},
		{httptest.NewRequest(http.MethodGet, "/?page=0&limit=10", nil), DefaultPage, 10},
		{httptest.NewRequest(http.MethodGet, "/?page=1&limit=-1", nil), 1, DefaultLimit},
		{httptest.NewRequest(http.MethodGet, fmt.Sprintf("/?page=1&limit=%d", MaxLimit+1), nil), 1, MaxLimit},
	}
}

func TestNewFromRequest(t *testing.T) {
	for _, test := range requestTestCases() {
		p := NewFromRequest(test.req)
		assert.Equal(t, test.page, p.Page)
		assert.Equal(t, test.limit, p.Limit)
	}
}

func TestNewFromContext(t *testing.T) {
	for _, test := range requestTestCases() {
		p := NewFromContext(&clevergo.Context{Request: test.req})
		assert.Equal(t, test.page, p.Page)
		assert.Equal(t, test.limit, p.Limit)
	}
}

func TestPaginationPageCount(t *testing.T) {
	cases := []struct {
		total     int64
		limit     int64
		pageCount int64
	}{
		{0, 10, 0},
		{1, 10, 1},
		{9, 10, 1},
		{10, 10, 1},
		{11, 10, 2},
		{19, 10, 2},
		{20, 10, 2},
		{21, 10, 3},
	}
	for _, test := range cases {
		p := Pagination{
			Total: test.total,
			Limit: test.limit,
		}
		assert.Equal(t, test.pageCount, p.PageCount())
	}
}

func TestPaginationOffset(t *testing.T) {
	cases := []struct {
		page   int64
		limit  int64
		offset int64
	}{
		{1, 10, 0},
		{2, 10, 10},
		{2, 15, 15},
		{3, 15, 30},
	}
	for _, test := range cases {
		p := Pagination{
			Page:  test.page,
			Limit: test.limit,
		}
		assert.Equal(t, test.offset, p.Offset())
	}
}
