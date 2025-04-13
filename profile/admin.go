package profile

const (
	MaxPageSize     = 200
	DefaultPageSize = 20
	FirstPage       = 1
)

func PageStartLimit(page, size int64) (start, limit int64) {
	if page <= 1 {
		page = FirstPage
	}
	if limit = size; limit <= 0 {
		limit = DefaultPageSize
	}
	start = (page - 1) * limit
	return
}
