/*
   @Author: StudentCWZ
   @Description:
   @File: post
   @Software: GoLand
   @Project: GoWeb
   @Date: 2022/3/23 09:43
*/

package mysql

import "GoWeb/lesson29/bluebell/models"

func CreatePost(p *models.Post) (err error) {
	sqlStr := `insert into post (post_id, title, content, author_id, community_id) values (?, ?, ?, ?, ?)`
	_, err = db.Exec(sqlStr, p.ID, p.Title, p.Content, p.AuthorID, p.CommunityID)
	return
}

func GetPostById(pid int64) (post *models.Post, err error) {
	post = new(models.Post)
	sqlStr := `select post_id, title, content, author_id, community_id, create_time from post where post_id = ?`
	err = db.Get(post, sqlStr, pid)
	return
}

func GetPostList(page, size int64) (posts []*models.Post, err error) {
	posts = make([]*models.Post, 0, size)
	sqlStr := `select post_id, title, content, author_id, community_id, create_time from post limit ?,?`
	err = db.Select(&posts, sqlStr, (page-1)*size, size)
	return
}
