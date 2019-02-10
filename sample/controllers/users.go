package controllers

import (
	"github.com/gin-gonic/gin"
	"../models"
)

// Post Modelを返す
func NewPost() Post {
    return Post{}
}

// idに合致する記事を取得
func (m PostRepository) GetByPostID(id int) *Post {
	var post Post
	db.Where(Post{ID: id}).Find(&amp;post)
	return &amp;post
}

// 記事を全検索
func (m PostRepository) GetAllPost() []Post {
	var post []Post
	db.Select("*").Find(&amp;post)
	return post
}

// 記事を投稿する
func (m PostRepository) CreatePost(header string, body string, author string) bool {
	post := Post {
			Header: header,
			Body:   body,
	}
	db.Create(&amp;post)
	return true
}

// 記事の件名と本文を更新する
func (m PostRepository) UpdatePost(id int, header string, body string) interface{} {
	post := Post{ID: id}
	db.Model(&amp;post).Updates(map[string]interface{}{
			"Header":header, 
			"Body":  body,
			})
	return true
}

// 記事を削除する
func (m PostRepository) DeletePost(id int) interface{} {
	post := Post{ID: id}
	db.Delete(post)
	return nil
}