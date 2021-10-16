package metrics

import (
	"time"

	"github.com/gin-gonic/gin"
)

const Key_Metrics_ReceivedOn = "Key_Metrics_ReceivedOn"

//
//	Should tag all incoming requests with a timestamp
//	Responders can use this timestamp before they respond to measure the time taken.
//	Could be used by devops for throttling etc...
//
func RequestResponseTimerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Set(Key_Metrics_ReceivedOn, time.Now())
		c.Next()
	}
}
