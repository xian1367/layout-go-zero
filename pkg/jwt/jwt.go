package jwt

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"github.com/golang-module/carbon/v2"
	"github.com/xian1367/layout-go-zero/config"
	"net/http"
	"strings"
)

var (
	ErrTokenExpired           = errors.New("令牌已过期")
	ErrTokenExpiredMaxRefresh = errors.New("令牌已过最大刷新时间")
	ErrTokenMalformed         = errors.New("请求令牌格式有误")
	ErrTokenInvalid           = errors.New("请求令牌无效")
	ErrHeaderEmpty            = errors.New("需要认证才能访问！")
	ErrHeaderMalformed        = errors.New("请求头中 Authorization 格式有误")
)

// CustomClaims 自定义载荷
type CustomClaims struct {
	UserID string `json:"user_id"`

	// RegisteredClaims
	// JWT 规定了7个官方字段，提供使用:
	// - iss (issuer)：发布者
	// - sub (subject)：主题
	// - iat (Issued At)：生成签名的时间
	// - exp (expiration time)：签名过期时间
	// - aud (audience)：观众，相当于接受者
	// - nbf (Not Before)：生效时间
	// - jti (JWT ID)：编号
	jwt.RegisteredClaims
}

func GenerateToken(userID string) (tokenString string, expireString string, err error) {
	expireAt := carbon.Now().AddMinutes(config.Get().Http.JwtAuth.AccessExpire)
	claims := CustomClaims{
		userID,
		jwt.RegisteredClaims{
			NotBefore: jwt.NewNumericDate(carbon.Now().StdTime()), // 签名生效时间
			IssuedAt:  jwt.NewNumericDate(carbon.Now().StdTime()), // 首次签名时间（后续刷新 Token 不会更新）
			ExpiresAt: jwt.NewNumericDate(expireAt.StdTime()),     // 签名过期时间
			Issuer:    config.Get().Http.Name,                     // 签名颁发者
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err = token.SignedString(config.Get().Http.JwtAuth.AccessSecret)
	expireString = expireAt.ToDateTimeString()
	return
}

// ParserToken 解析 Token，中间件中调用
func ParserToken(r *http.Request) (*CustomClaims, error) {
	tokenString, err := getTokenFromHeader(r)
	if err != nil {
		return nil, err
	}

	// 1. 调用 jwt 库解析用户传参的 Token
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return config.Get().Http.JwtAuth.AccessSecret, nil
	})

	if err != nil {
		if errors.Is(err, jwt.ErrTokenMalformed) {
			return nil, ErrTokenMalformed
		}
		if errors.Is(err, jwt.ErrTokenExpired) {
			return nil, ErrTokenExpired
		}
		return nil, ErrTokenInvalid
	}

	// 2. 将 token 中的 claims 信息解析出来和 CustomClaims 数据结构进行校验
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, ErrTokenInvalid
}

// RefreshToken 更新 Token，用以提供 refresh token 接口
func RefreshToken(r *http.Request) (tokenNewString string, expireString string, err error) {
	// 1. 调用 jwt 库解析用户传参的 Token
	claims, err := ParserToken(r)

	// 2. 解析出错，未报错证明是合法的 Token（甚至未到过期时间）
	if err != nil {
		return
	}

	// 3. 检查是否过了『最大允许刷新的时间』
	if carbon.CreateFromStdTime(claims.IssuedAt.Time).Gt(carbon.Now().SubMinutes(config.Get().Http.JwtAuth.AccessMaxRefreshExpire)) {
		// 修改过期时间
		expireAt := carbon.Now().AddMinutes(config.Get().Http.JwtAuth.AccessExpire)
		claims.RegisteredClaims.ExpiresAt = jwt.NewNumericDate(expireAt.StdTime())

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		tokenNewString, err = token.SignedString(config.Get().Http.JwtAuth.AccessSecret)
		expireString = expireAt.ToDateTimeString()
		return
	}

	err = ErrTokenExpiredMaxRefresh
	return
}

// Authorization:Bearer xxx
func getTokenFromHeader(r *http.Request) (string, error) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return "", ErrHeaderEmpty
	}
	// 按空格分割
	parts := strings.SplitN(authHeader, " ", 2)
	if !(len(parts) == 2 && parts[0] == "Bearer") {
		return "", ErrHeaderMalformed
	}
	return parts[1], nil
}
