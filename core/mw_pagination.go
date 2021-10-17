package core

import (
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
//	Schema: TIMESTAMP_OBJECTID
//
func ExtractValuesFromContinuationToken(token string) (int64, primitive.ObjectID) {
	split := strings.Split(token, "_")
	i, err := strconv.ParseInt(split[0], 10, 64)
	if err != nil {
		return 0, primitive.NilObjectID
	}
	if len(split) < 2 {
		return i, primitive.NilObjectID
	}
	j, err := primitive.ObjectIDFromHex(split[1])
	if err != nil {
		return 0, primitive.NilObjectID
	}
	return i, j
}
