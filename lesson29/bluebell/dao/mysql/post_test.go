/*
   @Author: StudentCWZ
   @Description:
   @File: post_test
   @Software: GoLand
   @Project: GoWeb
   @Date: 2022/3/25 15:46
*/

package mysql

import (
	"GoWeb/lesson29/bluebell/models"
	"GoWeb/lesson29/bluebell/settings"
	"testing"
)

func init() {
	dbConfig := settings.MySQLConfig{ // 测试的数据库
		Host:              "127.0.0.1",
		User:              "root",
		Password:          "root1234",
		DB:                "bluebell",
		MaxOpenConnection: 10,
		MaxIdleConnection: 10,
	}
	err := Init(&dbConfig)
	if err != nil {
		panic(err)
	}
}

func TestCreatePost(t *testing.T) {
	p := &models.Post{
		ID:          10,
		AuthorID:    123,
		CommunityID: 1,
		Title:       "test",
		Content:     "just a test",
	}
	err := CreatePost(p)
	if err != nil {
		t.Fatalf("CreatePost insert record into mysql failed, err: %v\n", err)
	}
	t.Logf("CreatePost insert record into mysql success")
}
