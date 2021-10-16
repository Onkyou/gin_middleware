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

//
//	Calculates processing time needed from request received (dated by middleware) to response sent
//	Should be called as late as possible before sending the response
//
func GetTimeConsumedFromContext(c *gin.Context) time.Duration {
	timeReceived := c.GetTime(Key_Metrics_ReceivedOn)
	timeResponse := time.Now()
	return timeResponse.Sub(timeReceived)
}
