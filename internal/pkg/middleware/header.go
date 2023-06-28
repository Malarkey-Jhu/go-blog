// Copyright 2023 Alan Jhu (朱倫君) <alan61109@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package middleware

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func NoCahe(c *gin.Context) {
	c.Header("Cache-Control", "no-cache, no-store, max-age=0, must-revalidate, value")
	c.Header("Expires", "Thu, 01 Jan 1970 00:00:00 GMT")
	c.Header("Last-Modified", time.Now().Format(http.TimeFormat))
	c.Next()
}

func Cors(c *gin.Context) {
	if c.Request.Method != "OPTIONS" {
		c.Next()
	} else {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, PATCH")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, X-Requested-With, X-Request-ID, Authorization, Origin, Content-Type, Accept")
		c.Header("Allow", "HEAD, GET, POST, PUT, PATCH, DELETE, OPTIONS")
		c.Header("Content-Type", "application/json")
		c.AbortWithStatus(http.StatusOK)
	}
}

func Secure(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("X-Frame-Options", "DENY")
	c.Header("X-Content-Type-Options", "nosniff")
	c.Header("X-XSS-Protection", "1; mode=block")
	if c.Request.TLS != nil {
		c.Header("Strict-Transport-Security", "max-age=31536000")
	}
}
