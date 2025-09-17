// TODO - Gus
// 1. GET /posts handler (feed with pagination):
//    - Get posts from followed users + own posts
//    - Order by creation date (newest first)
//    - Include pagination (limit, offset)
//    - Include user data and like status
// 2. GET /posts/:id handler (post details):
//    - Fetch specific post with user info
//    - Include like status for current user
// 3. POST /posts handler (create post):
//    - Validate required fields
//    - Handle image upload
//    - Create post record
//    - Update user's posts count
// 4. PUT /posts/:id handler (update post):
//    - Verify post ownership
//    - Update post data
// 5. DELETE /posts/:id handler:
//    - Verify ownership and delete