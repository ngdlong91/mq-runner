// Package middleware
package middleware

import "github.com/gin-gonic/gin"

type Middleware interface {
	Gin() gin.HandlerFunc
	MQ() MQHandle
}

type MQHandle func(payload interface{})
type MQHandleChain []MQHandle





