package repository

import (
	"context"
	"encoding/json"
	"feast-friends-api/internal/models"
	"feast-friends-api/internal/utils"
	"feast-friends-api/pkg/logger"
	"time"
)

// 1. CreatePost(post) function
// 2. GetPostByID(postID) function
// 3. GetFeedPosts(userID, limit, offset) function
// 4. GetUserPosts(userID, limit, offset) function
// 5. UpdatePost(postID, data) function
// 6. DeletePost(postID) function
// 7. LikePost(userID, postID) function
// 8. UnlikePost(userID, postID) function
// 9. CheckIfLiked(userID, postID) function
func CreatePost(post *models.Post) error {
	err := post.Validate()
	if err != nil {
		logger.Error("failed to validate post %s", err)
		return err
		
	}

	if post.CreatedAt.IsZero(){
		post.CreatedAt = time.Now()
	}

	recipeJSON , err :=json.Marshal(post.Recipe)
	if err != nil {
		logger.Error("Failed to encode recipe %w" , err)
		return err
	}

	query := `
        INSERT INTO posts (user_id, title, description, image_url, recipe, likes_count, comments_count, created_at)
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
        RETURNING id
    `
	rows , err := utils.DB.Query(context.Background(),query,
		post.UserID,
        post.Title,
        post.Description,
        post.ImageURL,
        recipeJSON,
        post.LikesCount,
        post.CommentsCount,
        post.CreatedAt,
	)
	if err != nil {
		logger.Error("failed to insert post: %w",err)
		return err
	}

	defer rows.Close()

	if rows.Next(){
		if err := rows.Scan(&post.ID); err != nil {
			logger.Error("failed to retrive post id: %w" , err)
			return err
		}
	}

	return nil
}

