package utils

import (
	"app-download/dto"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GeneratePaginationFromRequest(c *gin.Context) dto.Pagination {
	pagesize := 10
	page := 1
	sort := "created_at asc"
	query := c.Request.URL.Query()
	for key, value := range query {
		queryValue := value[len(value)-1]
		switch key {
		case "pagesize":
			pagesize, _ = strconv.Atoi(queryValue)
		case "page":
			page, _ = strconv.Atoi(queryValue)
		case "sort":
			sort = queryValue
		}
	}
	return dto.Pagination{
		Limit: pagesize,
		Page:  page,
		Sort:  sort,
	}

}
