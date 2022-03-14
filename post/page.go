package post

import "math"

type PageInfo struct {
	Total    int64
	Prev     int
	Next     int
	Current  int
	Start    int
	End      int
	LIndexes []int
	RIndexes []int
}

const (
	column_max int = 10
	search_max int = 10
)

func NewPageInfo() *PageInfo {
	pageInfo := &PageInfo{}

	return pageInfo
}

func SetPageInfo(pageInfo *PageInfo) error {
	if pageInfo.Total == 0 {
		return nil
	}

	pageInfo.Start = pageInfo.Current
	pageInfo.End = (int)(math.Ceil((float64(pageInfo.Total) / float64(search_max))))
	pageInfo.Prev = pageInfo.Current - 1
	pageInfo.Next = pageInfo.Current + 1

	pageMax := func(current int, end int) int {
		if end-current > search_max {
			return current + search_max
		}
		return end
	}(pageInfo.Current, pageInfo.End)

	pageMin := func(current int, end int) int {
		if end-current < search_max {
			return end - search_max
		}
		return pageInfo.Start
	}(pageInfo.Current, pageInfo.End)

	for i := pageMin; i < pageInfo.Current; i++ {
		pageInfo.LIndexes = append(pageInfo.LIndexes, i)
	}

	for i := pageInfo.Current + 1; i < pageMax; i++ {
		pageInfo.RIndexes = append(pageInfo.RIndexes, i)
	}

	if pageInfo.Current == pageInfo.End {
		pageInfo.Next = 0
	}

	return nil
}
