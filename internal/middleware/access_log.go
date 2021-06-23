package middleware

import (
	"blog/global"
	"blog/pkg/logger"
	"bytes"
	"github.com/gin-gonic/gin"
	"time"
)

type AccessLogWriters struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w AccessLogWriters) Write(p []byte) (int, error) {
	if n, err := w.body.Write(p); err != nil {
		return n, err
	}
	return w.ResponseWriter.Write(p)
}

func AccessLog() gin.HandlerFunc {
	return func(c *gin.Context) {
		bodyWriter := &AccessLogWriters{body: bytes.NewBufferString(""),
			ResponseWriter: c.Writer}
		c.Writer = bodyWriter

		beginTime := time.Now().Unix()
		c.Next()
		endTime := time.Now().Unix()

		fields := logger.Fields{
			"request" : c.Request.PostForm.Encode(),
			"response" : bodyWriter.body.String(),
		}
		global.Logger.WithFields(fields).Infof(c, "access log: method: %s, status_code: %d,begin_time:%d, end_time:%d",
			c.Request.Method,
			bodyWriter.Status(),
			beginTime,
			endTime,
			)
	}
}

//gin自带的Recovery()
//func Recovery() gin.HandlerFunc {
//	return func(c *gin.Context) {
//		defer func() {
//			if err := recover(); err != nil {
//				global.Logger.WithCallersFrames().Errorf("panic recover err : %v", err)
//				app.NewResponse(c).ToErrorResponse(errcode.ServerError)
//				c.Abort()
//			}
//		}()
//		c.Next()
//	}
//}