package post

type PageInfo struct {
	Total   int64
	Prev    int
	Next    int
	Current int
	Start   int
	End     int
	Indexes []int
	Max     int
}

const (
	max int = 10
)

func NewPageInfo() *PageInfo {
	pageInfo := &PageInfo{0, 0, 0, 0, 1, 0, nil, max}

	return pageInfo
}

func SetPageInfo(pageInfo *PageInfo) error {
	pageInfo.Start = pageInfo.Current
	pageInfo.End = (int(pageInfo.Total) / pageInfo.Max)

	if int(pageInfo.Total)%pageInfo.Max != 0 {
		pageInfo.End += 1
	}

	if pageInfo.Current > 1 {
		pageInfo.Prev = pageInfo.Current - 1
	}

	if pageInfo.Current < pageInfo.End && pageInfo.End >= max {
		pageInfo.Next = pageInfo.Current + 1
	}

	pageMax := func(pages int) int {
		if pages > max {
			return max
		}
		return pages
	}(pageInfo.End)

	for i := pageInfo.Start; i <= pageMax; i++ {
		pageInfo.Indexes = append(pageInfo.Indexes, i)
	}

	return nil
}
