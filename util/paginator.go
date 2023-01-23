package util

import "strconv"

type Paginator struct {
	Page     string
	PageSize string
}

func ToPaginator(page int, size int) *Paginator {
	return &Paginator{Page: strconv.Itoa(page), PageSize: strconv.Itoa(size)}
}
