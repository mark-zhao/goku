package v1

import (
	"errors"
	jwtgo "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"goku/component"
	modelv1 "goku/model/v1"
	"goku/utils/except"
	"goku/utils/logging"
	"goku/utils/options"
	"net/http"
	"time"
)
type loginRequest struct {
	Username	string	`form:"username" json:"username"`
	Password	string  `form:"password" json:"password"`
}
//注册路由
//func LoginRouter() {
//	APIs["/login"] = map[UriInterface]interface{}{
//		NewUri("POST", ""): (&LoginResource{}).Login,
//	}
//}
type LoginResource struct {
}

/*func (l *loginResource)Login(c *gin.Context) {
	var user loginRequest
	if err := c.ShouldBindJSON(&user); err != nil {
		logging.Error("invalid json provided err: ",err)
		c.String(http.StatusUnprocessableEntity, "invalid json provided")
		return
	}
	claims :=make(jwt.MapClaims)
	if v, ok := options.Conf.UserList[user.Username];ok && v == user.Password{
	//if len(user.Username) != 0 && len(user.Password) != 0 {
		claims["username"] = user.Username
		claims["iat"] = time.Now().Unix()
		claims["exp"] = time.Now().Add(time.Minute * 15).Unix()
		token, err := component.CreateToken(claims)

		if err != nil {
			logging.Error(err.Error())
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}
		c.JSON(http.StatusOK, gin.H{"accessToken":token})
		return
	}
	logging.Info("username or password error")
	c.String(http.StatusBadRequest,"username or password error")
}*/

// 注册信息
type RegistInfo struct {
	// 手机号
	Name string `json:"name"`
	// 密码
	Pwd string `json:"pwd"`
}

// LoginResult 登录结果结构
type LoginResult struct {
	Token string `json:"token"`
	options.User `json:"user"`
}

func (l *LoginResource)Login(c *gin.Context) {
	var loginReq modelv1.LoginReq
	if c.Bind(&loginReq) == nil {
		isPass, user := modelv1.LoginCheck(loginReq)
		if isPass {
			generateToken(c, user)
			logging.Info("生成token 成功")
		} else {
			resp.Render(c,except.ERROR_AUTH_USER,nil,errors.New(except.GetMsg(except.ERROR_AUTH_USER)))
			logging.Error(except.ERROR_AUTH_USER)
			return
		}
	} else {
		resp.Render(c,except.INVALID_PARAMS,nil,errors.New(except.GetMsg(except.INVALID_PARAMS)))
		logging.Error(except.INVALID_PARAMS)
		return
	}
}


// 生成令牌
func generateToken(c *gin.Context, user options.User) {
	j := &component.JWT{
		[]byte("Woshinibaba"),
	}
	claims := component.CustomClaims{
		user.Id,
		user.Name,
		user.Phone,
		user.Permission,
		jwtgo.StandardClaims{
			NotBefore: int64(time.Now().Unix() - 1000), // 签名生效时间
			ExpiresAt: int64(time.Now().Unix() + 3600), // 过期时间 一小时
			Issuer:    "mark",                   //签名的发行者
		},
	}

	token, err := j.CreateToken(claims)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": -1,
			"msg":    err.Error(),
		})
		return
	}
	user.Pwd=""
	data := LoginResult{
		User:  user,
		Token: token,
	}
	resp.Render(c,except.SUCCESS,data,nil)
	return
}