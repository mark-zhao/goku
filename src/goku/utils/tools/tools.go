package tools

import (
	"fmt"
	"goku/component"
	"goku/utils/logging"
	"math/rand"
	"reflect"

	"github.com/gin-gonic/gin"
)

func NewRandString() string {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		logging.Error(err)
	}
	return fmt.Sprintf("%x-%x", b[0:4], b[4:6])
}

func IsExistItem(value interface{}, array interface{}) bool {
	switch reflect.TypeOf(array).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(array)
		for i := 0; i < s.Len(); i++ {
			if reflect.DeepEqual(value, s.Index(i).Interface()) {
				return true
			}
		}
	}
	return false
}

func FunAuth(c *gin.Context, method string) (string, bool) {
	logging.Debug(method)
	if claims, T := c.Get("claims"); T == false {
		return "", false
	} else {
		if user, ok := claims.(*component.CustomClaims); ok {
			if IsExistItem(method, user.Role) || IsExistItem("admin", user.Role) {
				return user.Name, true
			}
		} else {
			logging.Debug("解析失败")
			return "", false
		}
	}
	return "", false
}
