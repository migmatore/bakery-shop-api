package sort

import (
	"fmt"
	"strings"
)

type Option struct {
	Column string
	Order  string
}

func GetSortOptions(queryParams map[string]string) (sortOption Option) {

	//sortOption := SortOption{}
	//for key, value := range queryParams {
	//	if key == "sort_by" || key == "sort_order" {
	//		sortOptions = append(sortOptions, SortOption{Column: key, Order: value})
	//	}
	//}
	if col, _ := queryParams["sort_by"]; col != "" {
		if order, _ := queryParams["sort_order"]; order == "asc" || order == "desc" {
			sortOption.Column = col
			sortOption.Order = strings.ToUpper(order)
		} else {
			sortOption.Column = col
			sortOption.Order = "ASC"
		}
	}

	return
}

//func GetSortOptions1(queryParams []utils.QueryParam) (sortOption SortOption) {
//
//	//sortOption := SortOption{}
//	//for key, value := range queryParams {
//	//	if key == "sort_by" || key == "sort_order" {
//	//		sortOptions = append(sortOptions, SortOption{Column: key, Order: value})
//	//	}
//	//}
//	for _, param := range queryParams {
//		if param.Name == "sort_by" && param.Value != "" {
//			sortOption.Column = param.Value
//		}
//
//		if param.Name == "sort_order" && (param.Value == "asc" || param.Value == "desc") {
//			sortOption.Order = strings.ToUpper(param.Value)
//		} else {
//			sortOption.Order = "ASC"
//		}
//	}
//
//	return
//}

func EnrichQueryWithSort(query string, option Option) string {
	if option.Column != "" {
		query = fmt.Sprintf("%s ORDER BY %s %s", query, option.Column, option.Order)
	}
	return query
}
