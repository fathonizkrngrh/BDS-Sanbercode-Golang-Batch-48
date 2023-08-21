package utils

import (
	"net/url"
	"quiz3/forms"
	"strconv"
	"strings"
)

func ParseInt(s string) int {
	val, _ := strconv.ParseInt(s, 10, 32)
	return int(val)
}

func IsValidURL(urlString string) bool {
	u, err := url.Parse(urlString)
	return err == nil && u.Scheme != "" && u.Host != ""
}

func CheckThickness(total_page int) string {
	var thickness string
	if total_page <= 100 {
		thickness = "tipis"
	} else if total_page <= 200 {
		thickness = "sedang"
	} else {
		thickness = "tebal"
	}

	return thickness
}

func BookFilter(queryText string, filter forms.BookFilter) string {
	if filter.Title != "" {
		queryText += " AND LOWER(title) LIKE '%" + strings.ToLower(filter.Title) + "%'"
	}
	if filter.MinYear != "" {
		queryText += " AND release_year >= " + filter.MinYear
	}
	if filter.MaxYear != "" {
		queryText += " AND release_year <= " + filter.MaxYear
	}
	if filter.MinPage != "" {
		queryText += " AND total_page >= " + filter.MinPage
	}
	if filter.MaxPage != "" {
		queryText += " AND total_page <= " + filter.MaxPage
	}
	if filter.SortByTitle == "asc" {
		queryText += " ORDER BY title ASC"
	} else if filter.SortByTitle == "desc" {
		queryText += " ORDER BY title DESC"
	}

	return queryText
}