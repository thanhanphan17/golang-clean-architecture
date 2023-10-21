package paging

import (
	"math"

	"gorm.io/gorm"
)

// GetOffset calculates the offset based on the current page
// and the limit per page.
func (p *Pagination) GetOffset() int {
	return (p.GetPage() - 1) * p.GetLimit()
}

// GetLimit returns the limit value for pagination.
func (p *Pagination) GetLimit() int {
	// If the limit is less than or equal to 0, set it to 10.
	if p.Limit <= 0 {
		p.Limit = 10
	}
	return p.Limit
}

// GetPage returns the value of the Page field in the Pagination struct.
// If the Page field is less than or equal to 0, it is set to 1.
// Returns:
//   - int: The value of the Page field.
func (p *Pagination) GetPage() int {
	if p.Page <= 0 {
		p.Page = 1
	}
	return p.Page
}

// GetSort returns the sort string for pagination.
// If the sort string is empty, it defaults to "id desc".
func (p *Pagination) GetSort() string {
	if p.Sort == "" {
		p.Sort = "created_at desc"
	}
	return p.Sort
}

// Paginate takes a value, pagination parameters, and a database connection and returns a function that performs pagination on the given database query.
func Paginate(value interface{}, pagination *Pagination, db *gorm.DB) func(db *gorm.DB) *gorm.DB {
	// Get the total number of rows
	var totalRows int64
	db.Model(value).Count(&totalRows)

	// Calculate the total number of pages
	totalPages := int(math.Ceil(float64(totalRows) / float64(pagination.GetLimit())))

	// Update the pagination object with total pages and total rows
	pagination.TotalPages = totalPages
	pagination.TotalRows = totalRows

	// Return a function that performs pagination on the database query
	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(pagination.GetOffset()).
			Limit(pagination.GetLimit()).
			Order(pagination.GetSort())
	}
}
