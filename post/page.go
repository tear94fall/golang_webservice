package post

import "math"

type PageInfo struct {
	Total    int64
	First    int
	Last     int
	Prev     int
	Next     int
	Current  int
	Start    int
	End      int
	LIndexes []int
	RIndexes []int
}

const (
	col_max int = 10
	row_max int = 10
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
	pageInfo.End = (int)(math.Ceil((float64(pageInfo.Total) / float64(row_max))))
	pageInfo.First = 1
	pageInfo.Last = pageInfo.End
	pageInfo.Prev = pageInfo.Current - 10
	pageInfo.Next = pageInfo.Current + 10

	if pageInfo.Current-row_max < 1 {
		pageInfo.Prev = 1
	} else if pageInfo.Current+row_max > pageInfo.End {
		pageInfo.Next = pageInfo.End
	}

	if pageInfo.Current == 1 {
		pageInfo.First = 0
		pageInfo.Prev = 0
	} else if pageInfo.Current == pageInfo.End {
		pageInfo.Last = 0
		pageInfo.Next = 0
	}

	startIndex := pageInfo.Current - (pageInfo.Current % 10) + 1
	endIndex := startIndex + 10

	if endIndex > pageInfo.End {
		endIndex = pageInfo.End + 1
	}

	for i := startIndex; i < pageInfo.Current; i++ {
		pageInfo.LIndexes = append(pageInfo.LIndexes, i)
	}

	for i := pageInfo.Current + 1; i < endIndex; i++ {
		pageInfo.RIndexes = append(pageInfo.RIndexes, i)
	}

	return nil
}
