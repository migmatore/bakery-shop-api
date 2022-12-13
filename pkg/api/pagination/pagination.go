package pagination

import (
	"fmt"
	"strconv"
)

type Pagination struct {
	Page    int
	PerPage int
}

func GetPaginationOptions(queryParams map[string]string) (pagination Pagination) {
	if col, _ := queryParams["page"]; col != "" {
		pagination.Page, _ = strconv.Atoi(col)
	} else {
		pagination.Page = 1
	}

	if col, _ := queryParams["per_page"]; col != "" {
		pagination.PerPage, _ = strconv.Atoi(col)
	} else {
		pagination.PerPage = 10
	}

	return
}

func EnrichQueryWithPagination(query string, pagination Pagination) string {
	offset := (pagination.Page - 1) * pagination.PerPage
	query = fmt.Sprintf("%s LIMIT %d OFFSET %d", query, pagination.PerPage, offset)

	return query
}
