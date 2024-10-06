package util

import (
	"github.com/gin-gonic/gin"
	"strings"
	"unicode"
)

func If[T any](condition bool, a, b T) T {
	if condition {
		return a
	}
	return b
}

func IfNotNil[T any](t *T, def T) T {
	if t != nil {
		return *t
	}
	return def
}

func ToSnakeCase(string string) string {
	builder := strings.Builder{}
	builder.Grow(len(string) + 1)
	for i, r := range string {
		if unicode.IsUpper(r) {
			if i != 0 {
				builder.WriteRune('_')
			}
			builder.WriteRune(unicode.ToLower(r))
		} else {
			builder.WriteRune(r)
		}
	}
	return builder.String()
}

var IP_HEADERS = [5]string{
	"X-Forwarded-For",
	"Proxy-Client-IP",
	"WL-Proxy-Client-IP",
	"HTTP_CLIENT_IP",
	"X-Real-IP",
}

func GetRequestIp(context *gin.Context) string {
	for _, header := range IP_HEADERS {
		if ip := context.Request.Header.Get(header); ip != "" {
			return ip
		}
	}
	return context.ClientIP()
}

func GetRequestUa(context *gin.Context) string {
	return context.Request.UserAgent()
}
