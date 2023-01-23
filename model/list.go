package model

import "math"

type models interface {
	User | Comment | Post
}

type List[T models] struct {
	Data  []*T `json:"data"`
	Total int  `json:"total"`
	Page  int  `json:"page"`
	Limit int  `json:"limit"`
}

func (l *List[T]) TotalPage() int {
	return int(math.Ceil(float64(l.Total/l.Limit))) + 1
}

func EmptyList[T models]() *List[T] {
	return &List[T]{Data: []*T{}, Total: 0, Page: 0, Limit: 0}
}
