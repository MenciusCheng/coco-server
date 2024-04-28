package request

const (
	DefaultPageSize = int(20)
	DefaultOffset   = int(0)
)

func FormatPage(page, pageSize int) (int, int) {
	offset, size := DefaultOffset, DefaultPageSize
	if pageSize > 0 {
		size = pageSize
	}
	if page > 1 {
		offset = (page - 1) * size
	}
	return offset, size
}
