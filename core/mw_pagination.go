package core

import (
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func PaginationMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		pageSizeForm := c.Request.FormValue("pageSize")
		pageSize, err := strconv.Atoi(pageSizeForm)
		if err != nil {
			pageSize = 100
		}
		//TODO: Currently set minimum pageSize to 2 to avoid infinite continuationToken generation edge-case
		if pageSize > 100 || pageSize < 2 {
			pageSize = 100
		} // limit to 100 elements per page

		c.Set("pageSize", pageSize)

		continuationToken := c.Request.FormValue("continuationToken")
		c.Set("continuationToken", continuationToken)

		c.Next()
	}
}

//
//	Extracts the values from a continuationToken like this one "0_61645e12fa3136ac261913dd"
//	Schema: TIMESTAMP_ID
//
func ExtractValuesFromContinuationToken(token string) (int64, string) {
	split := strings.Split(token, "_")
	i, err := strconv.ParseInt(split[0], 10, 64)
	if err != nil {
		return 0, ""
	}
	if len(split) < 2 {
		return i, ""
	}
	j := split[1]
	return i, j
}
