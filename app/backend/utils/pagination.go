package utils

import (
	"net/http"
	"strconv"
)

type Pagination struct {
	Limit int `json:"limit"`
	Page  int `json:"page"`
}

func GeneratePagination(r *http.Request) Pagination {
	page := 1
	limit := 20

	query := r.URL.Query()
	for key, value := range query {
		queryValue := value[len(value)-1]
		switch key {
		case "limit":
			limit, _ = strconv.Atoi(queryValue)
		case "page":
			page, _ = strconv.Atoi(queryValue)
		}
	}

	return Pagination{
		Limit: limit,
		Page:  page,
	}
}
