package helpers

import (
	"strconv"
	"strings"

	"github.com/aldisaputra17/dapur-fresh-id/entities"
	"github.com/gin-gonic/gin"
)

type PaginationResult struct {
	Result interface{}
	Error  error
}

func GeneratePaginationRequest(context *gin.Context) *entities.Pagination {
	// default limit, page & sort parameter
	limit := 10
	page := 1
	sort := "name asc"

	var searchs []entities.Search

	query := context.Request.URL.Query()

	for key, value := range query {
		queryValue := value[len(value)-1]

		switch key {
		case "limit":
			limit, _ = strconv.Atoi(queryValue)
			break
		case "page":
			page, _ = strconv.Atoi(queryValue)
			break
		case "sort":
			sort = queryValue
			break
		}

		// check if query parameter key contains dot
		if strings.Contains(key, ".") {
			// split query parameter key by dot
			searchKeys := strings.Split(key, ".")

			// create search object
			search := entities.Search{Column: searchKeys[0], Action: searchKeys[1], Query: queryValue}

			// add search object to searchs array
			searchs = append(searchs, search)
		}
	}

	return &entities.Pagination{Limit: limit, Page: page, Sort: sort, Searchs: searchs}
}
