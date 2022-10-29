package filter

import (
	"fmt"
	"regexp"
	"strings"
)

type Option struct {
	Column   string
	Operator string
	Value    string
}

func GetFilterOptions(queryParams map[string]string) []Option {
	filterOptions := make([]Option, 0)

	for key, value := range queryParams {
		if key != "sort_by" && key != "sort_order" {
			if value != "" {
				//opAndVal := regexp.MustCompile("^[a-z]{2}:").Split(value, 2)
				opAndVal := strings.SplitN(value, ":", 2)
				filterOptions = append(filterOptions, Option{Column: key, Operator: opAndVal[0], Value: opAndVal[1]})
			}
		}
	}

	return filterOptions
}

//func GetFilterOptions1(queryParams []utils.QueryParam) []FilterOption {
//	filterOptions := make([]FilterOption, 0)
//
//	for _, param := range queryParams {
//		if param.Name != "sort_by" && param.Name != "sort_order" {
//			if param.Value != "" {
//				//opAndVal := regexp.MustCompile("^[a-z]{2}:").Split(value, 2)
//				opAndVal := strings.SplitN(param.Value, ":", 2)
//				filterOptions = append(filterOptions, FilterOption{Column: param.Name, Operator: opAndVal[0], Value: opAndVal[1]})
//			}
//		}
//	}
//
//	return filterOptions
//}

func getOperator(operator string) (op string) {
	switch operator {
	case "eq":
		op = "="
	case "lt":
		op = "<"
	case "gt":
		op = ">"
	case "neq":
		op = "<>"
	case "lk":
		op = "LIKE"
	}

	return
}

// TODO check , in queries
func enrichQueryWIthStringOperator(operator string, value string) (q string) {
	switch operator {
	case "eq":
		q = "= " + value
	case "lt":
		q = "< " + value
	case "gt":
		q = "> " + value
	case "neq":
		q = "<> " + value
	case "lk":
		q = "LIKE '" + value + "'"
	case "in": // TODO in with strings
		q = "IN (" + value + ")"
	case "nin":
		q = "NOT IN (" + value + ")"
	}

	return
}

// TODO check , in queries
func enrichQueryWIthIntegerOperator(operator string, value string) (q string) {
	switch operator {
	case "eq":
		q = "= " + value
	case "lt":
		q = "< " + value
	case "gt":
		q = "> " + value
	case "lteq":
		q = "<= " + value
	case "gteq":
		q = ">= " + value
	case "neq":
		q = "<> " + value
	case "in":
		q = "IN (" + value + ")"
	case "nin":
		q = "NOT IN (" + value + ")"
	}

	return
}

func EnrichQueryWithFilter(query string, option []Option) string {
	reg := regexp.MustCompile("^[0-9,]+$")

	if len(option) != 0 {
		query = fmt.Sprintf("%s WHERE", query)

		for i, opt := range option {
			if opt.Value != "" {
				// check if not string
				if reg.MatchString(opt.Value) {
					query = fmt.Sprintf("%s %s %s", query, opt.Column, enrichQueryWIthIntegerOperator(opt.Operator, opt.Value))
				} else {
					query = fmt.Sprintf("%s %s %s", query, opt.Column, enrichQueryWIthStringOperator(opt.Operator, opt.Value))
				}

				if i != len(option)-1 {
					query = fmt.Sprintf("%s AND", query)
				}
			}
		}
	}

	return query
}
