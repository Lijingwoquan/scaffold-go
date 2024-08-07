package jwt

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
	"scaffold/models"
	"time"
)

var mySecret = []byte("老子就是喜欢看美女怎么滴")

// MyClaims 自定义声明结构体并内嵌jwt.StandardClaims
// jwt包自带的jwt.StandardClaims只包含了官方字段
// 我们这里需要额外记录username和userID字段，所以要自定义结构体
// 如果想要保存更多信息，都可以添加到这个结构体中
type MyClaims struct {
	UserID int64 `json:"user_id"`
	jwt.StandardClaims
}

// GenToken 生成JWT
func GenToken(u *models.User) (string, error) {
	// 创建一个我们自己的声明的数据
	c := MyClaims{
		u.Id,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add( //当下时间加上一周时间
				time.Duration(viper.GetInt("auth.jwt_expire")) * time.Hour).Unix(), // 过期时间
			Issuer: "liuzihao-blog", // 签发人
		},
	}
	// 使用指定的签名方法创建签名对象 --> 得到一个结构体
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	var mySecret = []byte(viper.GetString("auth.secret"))
	// 使用指定的secret签名并获得完整的编码后的字符串token
	return token.SignedString(mySecret)
}

// ParseToken 解析JWT
func ParseToken(tokenString string) (*MyClaims, error) {
	// 解析token
	var mc = new(MyClaims)
	token, err := jwt.ParseWithClaims(tokenString, mc, func(token *jwt.Token) (i interface{}, err error) {
		return mySecret, nil
	})
	if err != nil {
		return nil, err
	}
	if token.Valid { // 校验token
		return mc, nil
	}

	return nil, errors.New("invalid token")
}
