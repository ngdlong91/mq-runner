// Package stats
package stats

import (
	"github.com/gin-gonic/gin"
	"github.com/ngdlong91/mq-runner/klog"
	"github.com/ngdlong91/mq-runner/middleware"
)

type StatisticData struct {

}

type stats struct {
	log klog.Logger

}

func (s *stats) doProcess(data StatisticData)  {

}

func (s *stats) Gin() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Convert from gin context to internal payload

		data := StatisticData{}
		s.doProcess(data)
	}
}

func (s *stats) MQ() middleware.MQHandle {
	return func(payload interface{}) {
		// Convert from mq payload to internal payload
		data := StatisticData{}
		s.doProcess(data)
	}
}

func NewStatsMiddleware() *stats {
	return &stats{
		log:klog.NewFileLogger(),
	}
}
