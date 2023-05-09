package utils

import (
	"math"
)

type Filter struct {
	Page     int
	PageSize int
}
type Metadata struct {
	CurrentPage  int
	PageSize     int
	FirstPage    int
	LastPage     int
	TotalRecords int
}

func (f *Filter) Limit() int {
	return f.PageSize
}
func (f *Filter) Offset() int {
	return int(math.Abs(float64(f.Page-1) * float64(f.PageSize)))
}
func ComputerMetadata(totalRecords, currentPage, pageSize *int) Metadata {
	if *totalRecords == 0 {
		return Metadata{}
	}
	return Metadata{
		CurrentPage: *currentPage,
		PageSize:    *pageSize,
		FirstPage:   1,
		LastPage:    int(math.Ceil(float64(*totalRecords) / float64(*pageSize))),
		TotalRecords: *totalRecords,
	}
}
