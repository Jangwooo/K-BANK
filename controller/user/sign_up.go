package user

import (
	"database/sql"
	"net/http"
	"path/filepath"

	"K-BANK/lib"
	"K-BANK/model"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type SignupRequest struct {
	ID          string `form:"id"  binding:"required" json:"id"`
	Pwd         string `form:"pwd" binding:"required" json:"pwd"`
	SimplePwd   string `form:"simple_pwd" binding:"required" json:"simple_pwd"`
	PhoneNumber string `form:"phone_number" binding:"required" json:"phone_number"`
	SSN         string `form:"ssn" binding:"required" json:"ssn"`
	Name        string `form:"name" binding:"required" json:"name"`
	Nickname    string `form:"nickname" json:"nickname"`
	Agree       string `form:"agree" binding:"required" json:"agree"`
}

func SignUpHandler(c *gin.Context) {
	type Response struct {
		Msg string `json:"msg,omitempty"`
	}
	res := Response{}
	req := new(SignupRequest)
	err := c.Bind(req)
	if err != nil {
		res.Msg = "리퀘스트 형식이 잘못되었습니다"
		c.JSON(http.StatusBadRequest, res)
		return
	}

	result := lib.DuplicateCheck("id", req.ID)
	if result == false {
		res.Msg = "ID 중복됨"
		c.JSON(http.StatusBadRequest, res)
		return
	}

	result = lib.DuplicateCheck("ssn", req.SSN)
	if result == false {
		res.Msg = "한명당 하나의 ID만 가질 수 있습니다"
		c.JSON(http.StatusBadRequest, res)
		return
	}

	pwd, err := bcrypt.GenerateFromPassword([]byte(req.Pwd), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	ssn, err := lib.Cipher.Encrypt(req.SSN)
	if err != nil {
		panic(err)
	}

	simplePwd, err := bcrypt.GenerateFromPassword([]byte(req.SimplePwd), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}

	var n sql.NullString

	if req.Nickname == "" {
		n.Valid = false
	} else {
		n.String = req.Nickname
		n.Valid = true
	}

	var uploadPath string
	file, err := c.FormFile("profile")
	if err != nil {
		res.Msg = "프로필 사진을 반드시 등록해야 합니다!"
		c.JSON(http.StatusBadRequest, res)
		return
	}

	ext := filepath.Ext(file.Filename)
	uploadPath = "./images/profile/" + req.ID + ext

	if err := c.SaveUploadedFile(file, uploadPath); err != nil {
		panic(err)
	}

	u := model.User{
		ID:          req.ID,
		Password:    string(pwd),
		PhoneNumber: req.PhoneNumber,
		SSN:         ssn,
		Name:        req.Name,
		NickName:    n,
		UserType:    "normal",
		Agree:       req.Agree,
		ProfilePic: &model.ProfilePic{
			UserID: req.ID,
			Path:   uploadPath,
		},
		SimplePwd: &model.SimplePwd{
			UserID: req.ID,
			Pwd:    string(simplePwd),
		},
		CheckingAccount: nil,
	}

	err = model.DB.Create(&u).Error
	if err != nil {
		panic(err)
	}

	res.Msg = "성공"
	c.JSON(http.StatusOK, res)
}
