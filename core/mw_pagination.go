package core

import (
	"strconv"

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
