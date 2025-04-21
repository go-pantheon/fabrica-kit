// Package profile provides environment profile and configuration utilities
// for application setup and runtime behavior control.
package profile

const (
	// MaxPageSize defines the maximum number of items per page for pagination.
	MaxPageSize = 200
	// DefaultPageSize defines the default number of items per page if not specified.
	DefaultPageSize = 20
	// FirstPage defines the first page number for pagination.
	FirstPage = 1
)

// PageStartLimit calculates the start index and limit for pagination.
// It returns the start offset and the number of items to fetch based on the page number and size.
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
