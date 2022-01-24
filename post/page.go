package post

type PageInfo struct {
	Total   int64
	Prev    bool
	Next    bool
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
	pageInfo := &PageInfo{0, false, false, 0, 1, 0, nil, max}

	return pageInfo
}

func SetPageInfo(pageInfo *PageInfo) error {
	pageInfo.Start = pageInfo.Current
	pageInfo.End = (int(pageInfo.Total) / pageInfo.Max) + 1

	if pageInfo.Current > 1 {
		pageInfo.Prev = true
	}

	if pageInfo.Current < pageInfo.End {
		pageInfo.Next = true
	}

	for i := pageInfo.Start; i < pageInfo.End; i++ {
		pageInfo.Indexes = append(pageInfo.Indexes, i)
	}

	return nil
}
