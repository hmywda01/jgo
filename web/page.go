package web

import (
	"fmt"
	"strings"
)

const (
	DefaultPageSize = 25
)

var pageSize = DefaultPageSize

type Page struct {
	Page   uint
	Offset uint
	Url    string
	ItmCnt int
	Params map[string]string
}

func (p Page) Prev() uint {
	if p.Page <= 1 {
		return 1
	}
	return p.Page - 1
}

func (p Page) Next() uint {
	return p.Page + 1
}

func (p Page) GetUrl(params map[string]string) string {
	var allParams = make(map[string]string)
	for k, v := range p.Params {
		allParams[k] = v
	}
	for k, v := range params {
		allParams[k] = v
	}
	if len(allParams) == 0 {
		return p.Url
	}
	var parts []string
	for k, v := range allParams {
		if v == "" {
			parts = append(parts, fmt.Sprintf("%s", k))
		} else {
			parts = append(parts, fmt.Sprintf("%s=%s", k, v))
		}
	}
	return fmt.Sprintf("%s?%s", p.Url, strings.Join(parts, "&"))
}

func (p Page) GetPageUrl(page uint) string {
	return p.GetUrl(map[string]string{"page": fmt.Sprintf("%d", page)})
}

func (p Page) IsFirstPage() bool {
	return p.Page == 1
}

func (p Page) IsLastPage() bool {
	return p.ItmCnt == 0
}

func (p Page) GetPrevUrl() string {
	if p.Prev() > 1 {
		return p.GetPageUrl(p.Prev())
	}
	return p.GetUrl(nil)
}

func (p Page) GetNextUrl() string {
	return p.GetPageUrl(p.Next())
}

func SetPageSize(size int) {
	pageSize = size
}

func GetPage(r *Response, url string) *Page {
	return GetPageWithSize(r, url, pageSize)
}

func GetPageWithSize(r *Response, url string, size int) *Page {
	pageInput := r.Request.GetUrlParameterUInt("page")
	if pageInput < 1 {
		pageInput = 1
	}
	offset := (pageInput - 1) * uint(size)
	page := Page{
		Page:   pageInput,
		Offset: offset,
		Url:    url,
		Params: map[string]string{},
	}
	r.Helper["Page"] = &page
	return &page
}