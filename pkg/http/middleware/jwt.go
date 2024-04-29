package middleware

import (
	"context"
	server "github.com/xian1367/layout-go-zero/pkg/http"
	"github.com/xian1367/layout-go-zero/pkg/jwt"
	"net/http"
	"strings"
)

func Jwt(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tokenString := getTokenFromHeader(r)

		claims, err := jwt.ParserToken(tokenString)
		// JWT 解析失败，有错误发生
		if err != nil {
			server.ValidationError(w, err)
			return
		}
		ctx := context.WithValue(r.Context(), "user_id", claims.UserID)

		next(w, r.WithContext(ctx))
	}
}

// Authorization:Bearer xxx
func getTokenFromHeader(r *http.Request) string {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return ""
	}
	// 按空格分割
	parts := strings.SplitN(authHeader, " ", 2)
	if !(len(parts) == 2 && parts[0] == "Bearer") {
		return ""
	}
	return parts[1]
}
