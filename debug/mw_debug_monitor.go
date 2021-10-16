package debug

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

//
//	Should tag all incoming requests with a timestamp
//	Responders can use this timestamp before they respond to measure the time taken.
//	Could be used by devops for throttling etc...
//
func DebugMonitorMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		fmt.Printf("Request URL :: %s \n", c.Request.URL)
		fmt.Printf("Request Params :: %v \n", c.Params)
		fmt.Printf("Request Headers :: %v \n", c.Request.Header)
		fmt.Printf("Request Body :: %v \n", c.Request.Body)

		c.Next()
	}
}
