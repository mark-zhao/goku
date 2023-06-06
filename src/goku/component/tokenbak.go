package component

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"goku/utils/logging"
)

const AccessSecret  =  "ACCESS_SECRET_KKKK"
type Header struct {
	AccessToken	string	`form:"accessToken" json:"accessToken"`
}

//通过用户信息获取token
func CreateToken(info jwt.MapClaims) (string, error) {

	var err error
	//os.Setenv("ACCESS_SECRET", "KKKK")
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, info)
	//token, err := at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
	token, err := at.SignedString([]byte(AccessSecret))
	if err != nil {
		return "", err
	}
	if err := info["username"];err == nil  {
		logging.Info(info["username"],"token is :", token)
	}
	return token, nil
}

//获取jwt原信息
func ExtractTokenMetaData(r string) (username string, error error) {
	token, err := VerifyToken(r)
	if err != nil {
		logging.Error(err)
		return "", err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		username, ok := claims["username"].(string)
		if !ok {
			return "", err
		}
		return username, nil
	}
	return "", err
}


//校验token
func VerifyToken(r string) (*jwt.Token, error) {
	token, err := jwt.Parse(r, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method :%v", token.Header["accessToken"])
		}
		return []byte(AccessSecret), nil
	})

	if err != nil {
		logging.Error("unexpected signing method :%v", token.Header["accessToken"])
		return nil, err
	}
	return token, nil
}

func GetUsername(c *gin.Context) (string, error){
	h := Header{}
	if err := c.ShouldBindHeader(&h);err != nil {
		logging.Error("invalid header provided,err: ",err)
		return "", err
	}
	logging.Debug(h)
	username, err := ExtractTokenMetaData(h.AccessToken)
	if err != nil {
		logging.Error("header has no accessToken, err:",err)
		return "", err
	}
	logging.Info("username: ",username)
	return  username ,nil
}