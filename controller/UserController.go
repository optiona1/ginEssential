package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
	"log"
	"strconv"
	"zd.xyz/ginessential/model"
	"zd.xyz/ginessential/util"
	"zd.xyz/ginessential/common"
)

func Register(ctx *gin.Context) {
	DB := common.DB
	// 获取参数
	name := ctx.PostForm("name")
	telephone := ctx.PostForm("telephone")
	password := ctx.PostForm("password")

	// 数据验证
	if len(telephone) != 11 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 422,
			"msg":  "手机号必须为 11 位",
		})
		return
	}

	_, err := strconv.Atoi(telephone)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 422,
			"msg":  "手机号必须为数字",
		})
		return
	}

	if len(password) < 6 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 422,
			"msg":  "密码不能少于 6 位",
		})
		return
	}

	// 如果没有传入名称，生成一个 10 位的随机字符串
	if len(name) == 0 {
		name = util.RandomString(10)
	}

	log.Println(name, telephone, password)
	// 判断手机号是否存在
	if isTelephoneExist(DB, telephone) {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 422,
			"msg":  "用户已经存在",
		})
		return
	}

	// 创建用户
	newUser := model.User{
		Name:      name,
		Telephone: telephone,
		Password:  password,
	}
	DB.Create(&newUser)

	// 返回结果
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "注册成功",
	})
}

func isTelephoneExist(db *gorm.DB, telephone string) bool {
        var user model.User
        db.Where("telephone = ?", telephone).First(&user)
        if user.ID != 0 {
                return true
        }
        return false
}

