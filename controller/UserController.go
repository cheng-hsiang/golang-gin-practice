package controller

import (
	"fmt"
	"gin_api/common"
	"gin_api/dto"
	"gin_api/model"
	"gin_api/response"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func Register(ctx *gin.Context) {
	DB := common.GetDB()

	// 獲取參數 1  map

	// var request_map = make(map[string]string)

	// request_map = Json.newDecoder(ctx.Request.body).Decode(&request_map)

	// 獲取參數 結構體
	var requestUser = model.User{}
	// json.NewDecoder(ctx.Request.Body).Decode(&requestUser)

	// bind函數
	ctx.Bind(&requestUser)
	fmt.Println("request user", requestUser.Name)

	// 獲取參數
	name := ctx.PostForm("name")
	telephone := ctx.PostForm("telephone")
	password := ctx.PostForm("password")

	fmt.Println(name, "name", len(name))
	fmt.Println(telephone, "telephone", len(telephone))
	fmt.Println(password, "password", len(password))
	//驗證

	if len(password) < 6 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "必須為6")
		return
	}

	//判斷是否存在
	if isTelephoneExist(DB, telephone) {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "用戶已註冊")
		return
	}

	//創建用戶 加密用戶密碼
	hasePassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		response.Response(ctx, http.StatusUnprocessableEntity, 500, nil, "加密失败")
		return
	}

	newUser := model.User{
		Name:      name,
		Telephone: telephone,
		Password:  string(hasePassword),
	}
	DB.Create(&newUser)

	token, err := common.ReleaseToken(newUser)
	if err != nil {
		response.Response(ctx, http.StatusUnprocessableEntity, 500, nil, "token 生成失敗")
		log.Printf("token generate error: %v", err)
		return
	}
	//返回結果
	response.Success(ctx, gin.H{"token": token}, "注册成功")
}

func Login(ctx *gin.Context) {
	DB := common.GetDB()

	// 獲取參數
	telephone := ctx.PostForm("telephone")
	password := ctx.PostForm("password")

	//驗證

	//是否存在
	var user model.User
	DB.Where("telephone = ?", telephone).First(&user)
	fmt.Print(user)
	if user.ID == 0 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "用戶不存在")
		return
	}

	//密碼正確
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		response.Response(ctx, http.StatusUnprocessableEntity, 400, nil, "密碼錯誤")
		return
	}

	//發放token
	token, err := common.ReleaseToken(user)
	if err != nil {
		response.Response(ctx, http.StatusUnprocessableEntity, 500, nil, "token 生成失敗")
		log.Printf("token generate error: %v", err)
		return
	}

	//返回結果
	response.Success(ctx, gin.H{"token": token}, "登陸成功")
}
func Info(ctx *gin.Context) {
	user, _ := ctx.Get("user")
	ctx.JSON(http.StatusOK, gin.H{"code": 200, "data": gin.H{"user": dto.ToUserDto(user.(model.User))}})
}

func isTelephoneExist(db *gorm.DB, telephone string) bool {
	var user model.User
	db.Where("telephone = ?", telephone).First(&user)
	if user.ID != 0 {
		return true
	}
	return false
}
