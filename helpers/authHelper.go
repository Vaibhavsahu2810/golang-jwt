package helper

import (
	"errors"

	"github.com/gin-gonic/gin"
)

func CheckUserType(c *gin.Context, roll string) (err error) {
	//check user type
	err = nil
	userType := c.GetString("user_type")
	if userType != roll {
		err = errors.New("user type is not match")
	}
	return err
}

func MatchUserTypeToUid(c *gin.Context, userId string) (err error) {
	userType := c.GetString("user_type")

	id := c.GetString("uid")
	err = nil
	if userType == "USER" && id != userId {
		return errors.New("user_type is USER but uid is not match")
	}
	err = CheckUserType(c, userType)
	if err != nil {
		return err
	}
}
