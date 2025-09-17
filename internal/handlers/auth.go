// TODO - Luiz
// 1. POST /auth/register handler:
//    - Validate email, username, password
//    - Check if user already exists
//    - Hash password with bcrypt
//    - Create user in Supabase Auth
//    - Return success response
// 2. POST /auth/login handler:
//    - Validate credentials with Supabase Auth
//    - Generate JWT token
//    - Return user data + token
// 3. POST /auth/logout handler:
//    - Invalidate JWT token (add to blacklist)
//    - Return success response
// 4. GET /auth/me handler:
//    - Return current authenticated user data
//    - Require authentication middleware