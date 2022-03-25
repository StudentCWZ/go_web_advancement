/*
   @Author: StudentCWZ
   @Description:
   @File: post_test
   @Software: GoLand
   @Project: GoWeb
   @Date: 2022/3/25 15:25
*/

package controller

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreatePostHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	url := "/api/v1/post"
	r.POST(url, CreatePostHandler)
	body := `{
		"community_id": 1,
		"title": "test",
		"content": "just a test"
	}`
	req, _ := http.NewRequest(http.MethodPost, url, bytes.NewReader([]byte(body)))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	// 判断响应的内容是不是按预期返回了需要登录的错误
	// 方法一：判断响应内容中是不是包含指定的字符串
	//assert.Contains(t, w.Body.String(), "需要登录")
	// 方法二：将响应的内容反序列化到 ResponseData ，然后怕奴蛋字段与预期是否是一致
	res := new(ResponseData)
	if err := json.Unmarshal(w.Body.Bytes(), res); err != nil {
		t.Fatalf("json.Unmarshal w.Body failed, err: %v", err)
	}
	assert.Equal(t, res.Code, CodeNeedLogin)

}
