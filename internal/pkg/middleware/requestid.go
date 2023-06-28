// Copyright 2023 Alan Jhu (朱倫君) <alan61109@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package middleware

import (
	"github.com/Malarkey-Jhu/miniblog/internal/pkg/known"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func RequestID() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestID := c.Request.Header.Get(known.XRequestIDKey)

		if requestID == "" {
			requestID = uuid.New().String()
		}

		c.Set(known.XRequestIDKey, requestID)
		c.Writer.Header().Set(known.XRequestIDKey, requestID)
		c.Next()
	}
}
