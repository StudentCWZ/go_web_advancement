/*
   @Author: StudentCWZ
   @Description:
   @File: community
   @Software: GoLand
   @Project: GoWeb
   @Date: 2022/3/22 15:21
*/

package mysql

import (
	"GoWeb/lesson29/bluebell/models"
	"database/sql"
	"go.uber.org/zap"
)

func GetCommunityList() (communityList []*models.Community, err error) {
	sqlStr := "select community_id, community_name from community"
	err = db.Select(&communityList, sqlStr)
	if err != nil {
		if err == sql.ErrNoRows {
			zap.L().Warn("there is no community in db")
			err = nil
		}
		zap.L().Error("query data failed", zap.Error(err))
	}
	return
}

// GetCommunityDetailByID 根据 ID 查询社区详情
func GetCommunityDetailByID(id int64) (communityDetail *models.CommunityDetail, err error) {
	communityDetail = new(models.CommunityDetail)
	sqlStr := `select community_id, community_name, introduction, create_time from community where community_id = ?`
	err = db.Get(communityDetail, sqlStr, id)
	if err != nil {
		if err == sql.ErrNoRows {
			zap.L().Warn("there is no communityDetail in db")
			err = ErrorInvalidID
		}
		zap.L().Error("query data failed", zap.Error(err))
	}
	return
}
