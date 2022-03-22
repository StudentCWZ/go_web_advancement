/*
   @Author: StudentCWZ
   @Description:
   @File: community
   @Software: GoLand
   @Project: GoWeb
   @Date: 2022/3/22 15:19
*/

package logic

import (
	"GoWeb/lesson29/bluebell/dao/mysql"
	"GoWeb/lesson29/bluebell/models"
)

func GetCommunityList() ([]*models.Community, error) {
	// 查数据库 查找到所有的 community 并返回
	return mysql.GetCommunityList()
}

func GetCommunityDetail(id int64) (*models.CommunityDetail, error) {
	return mysql.GetCommunityDetailByID(id)
}
