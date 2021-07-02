package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	sts "github.com/tencentyun/qcloud-cos-sts-sdk/go"
	"net/http"
	"osaka/global"
	"osaka/model"
	"osaka/repository"
	"osaka/routes/JSONmap"
	"strconv"
	"time"
)

// HandOutCOSCredential reference: https://github.com/tencentyun/qcloud-cos-sts-sdk/tree/master/go
func HandOutCOSCredential(context *gin.Context) {
	var err error
	var res *sts.CredentialResult
	tx := global.DB.Begin()
	defer func() {
		var errMsg string
		if r := recover(); r != nil {
			errMsg = fmt.Sprint(r)
			tx.Rollback()
			context.JSON(http.StatusOK, gin.H{
				"success": false,
				"exc":     errMsg,
			})
		} else {
			tx.Commit()
			context.JSON(http.StatusOK, res)
		}
		return
	}()

	c := sts.NewClient(
		global.CosSecretId,
		global.CosSecretKey,
		nil,
	)
	opt := &sts.CredentialOptions{
		DurationSeconds: int64(time.Hour.Seconds()),
		Region:          "ap-chengdu",
		Policy: &sts.CredentialPolicy{
			Statement: []sts.CredentialPolicyStatement{
				{
					Action: []string{
						"name/cos:*",
					},
					Effect: "allow",
					Resource: []string{
						//这里改成允许的路径前缀，可以根据自己网站的用户登录态判断允许上传的具体路径，例子： a.jpg 或者 a/* 或者 * (使用通配符*存在重大安全风险, 请谨慎评估使用)
						"*",
					},
				},
			},
		},
	}
	res, err = c.GetCredential(opt)
	if err != nil {
		panic(err)
	}
	//fmt.Printf("%+v\n", res)
	//fmt.Printf("%+v\n", res.Credentials)
}

func RegisterFile(context *gin.Context) {
	var err error
	var fileId uint
	tx := global.DB.Begin()
	defer func() {
		var errMsg string
		if r := recover(); r != nil {
			errMsg = fmt.Sprint(r)
			tx.Rollback()
			context.JSON(http.StatusOK, gin.H{
				"success": false,
				"exc":     errMsg,
			})
		} else {
			tx.Commit()
			context.JSON(http.StatusOK, gin.H{
				"success": true,
				"exc":     errMsg,
				"fileId":  fileId,
			})
		}
		return
	}()

	var json JSONmap.RegisterFile
	err = context.BindJSON(&json)
	if err != nil {
		panic(err)
	}

	file := new(model.ResourceFile)
	file.Url = json.Url
	file.Name = json.Name
	file.Type = json.Type

	err = tx.Save(file).Error
	if err != nil {
		panic(err)
	}
	fileId = file.ID
}

func GetFileInfo(context *gin.Context) {
	type fileDto struct {
		Id   uint   `json:"id"`
		Url  string `json:"url"`
		Name string `json:"name"`
		Type uint   `json:"type"`
	}
	var err error
	fileInfo := new(fileDto)
	tx := global.DB.Begin()
	defer func() {
		var errMsg string
		if r := recover(); r != nil {
			errMsg = fmt.Sprint(r)
			tx.Rollback()
			context.JSON(http.StatusOK, gin.H{
				"success": false,
				"exc":     errMsg,
			})
		} else {
			tx.Commit()
			context.JSON(http.StatusOK, gin.H{
				"success": true,
				"file":    fileInfo,
			})
		}
		return
	}()

	fileId, err := strconv.ParseUint(context.Query("id"), 10, 64)

	file, err := repository.FindFileById(uint(fileId))
	if err != nil {
		panic(err)
	}

	// set file dto
	fileInfo.Id = file.ID
	fileInfo.Url = file.Url
	fileInfo.Name = file.Name
	fileInfo.Type = file.Type
}
