package paging

type Pagination struct {
	Limit      int         `json:"limit,omitempty" query:"limit" form:"limit"`
	Page       int         `json:"page,omitempty" query:"page" form:"page"`
	Sort       string      `json:"sort,omitempty" query:"sort" form:"sort"`
	TotalRows  int64       `json:"total_rows"`
	TotalPages int         `json:"total_pages"`
	Rows       interface{} `json:"rows"`
}
