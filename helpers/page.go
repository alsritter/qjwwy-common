package helpers

const (
	pageDefaultPageIndex = 1
	pageDefaultPageSize  = 10
)

type (
	Pager interface {
		Offset() int
		SetTotal(int)
		Total() int
		PageSize() int
		NeedAll() bool
	}
	pager struct {
		pageIndex int
		pageSize  int
		total     int
		needAll   bool
	}
)

func NewPager(pageIndex, pageSize int, needAll bool) Pager {
	if pageIndex < 1 {
		pageIndex = pageDefaultPageIndex
	}
	if pageSize <= 0 {
		pageSize = pageDefaultPageSize
	}
	p := pager{
		pageSize:  pageSize,
		pageIndex: pageIndex,
		needAll:   needAll,
	}
	return &p
}

func (p *pager) Offset() int {
	return (p.pageIndex - 1) * p.pageSize
}

func (p *pager) SetTotal(total int) {
	p.total = total
}

func (p *pager) Total() int {
	return p.total
}

func (p *pager) PageSize() int {
	return p.pageSize
}

func (p *pager) NeedAll() bool {
	return p.needAll
}

type (
	Int32Pager interface {
		Offset() int32
		SetTotal(int32)
		Total() int32
		PageSize() int32
		NeedAll() bool
	}
	int32Pager struct {
		pageIndex int32
		pageSize  int32
		total     int32
		needAll   bool
	}
)

func NewInt32Pager(pageIndex, pageSize int32, needAll bool) Int32Pager {
	if pageIndex < 1 {
		pageIndex = pageDefaultPageIndex
	}
	if pageSize <= 0 {
		pageSize = pageDefaultPageSize
	}
	return &int32Pager{
		pageIndex: pageIndex,
		pageSize:  pageSize,
		total:     0,
		needAll:   needAll,
	}
}

func (p *int32Pager) Offset() int32 {
	return (p.pageIndex - 1) * p.pageSize
}

func (p *int32Pager) SetTotal(total int32) {
	p.total = total
}

func (p *int32Pager) Total() int32 {
	return p.total
}

func (p *int32Pager) PageSize() int32 {
	return p.pageSize
}

func (p *int32Pager) NeedAll() bool {
	return p.needAll
}

type (
	Int64Pager interface {
		Offset() int64
		SetTotal(int64)
		Total() int64
		PageSize() int64
		NeedAll() bool
	}
	int64Pager struct {
		pageIndex int64
		pageSize  int64
		total     int64
		needAll   bool
	}
)

func NewInt64Pager(pageIndex, pageSize int64, needAll bool) Int64Pager {
	if pageIndex < 1 {
		pageIndex = pageDefaultPageIndex
	}
	if pageSize <= 0 {
		pageSize = pageDefaultPageSize
	}
	return &int64Pager{
		pageIndex: pageIndex,
		pageSize:  pageSize,
		total:     0,
		needAll:   needAll,
	}
}

func (p *int64Pager) Offset() int64 {
	return (p.pageIndex - 1) * p.pageSize
}

func (p *int64Pager) SetTotal(total int64) {
	p.total = total
}

func (p *int64Pager) Total() int64 {
	return p.total
}

func (p *int64Pager) PageSize() int64 {
	return p.pageSize
}

func (p *int64Pager) NeedAll() bool {
	return p.needAll
}
