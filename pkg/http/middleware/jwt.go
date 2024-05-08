package middleware

import (
	"context"
	server "github.com/xian1367/layout-go-zero/pkg/http"
	"github.com/xian1367/layout-go-zero/pkg/jwt"
	"net/http"
)

func Jwt(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		claims, err := jwt.ParserToken(r)
		// JWT 解析失败，有错误发生
		if err != nil {
			server.ValidationError(w, err)
			return
		}
		ctx := context.WithValue(r.Context(), "user_id", claims.UserID)

		next(w, r.WithContext(ctx))
	}
}
