cat > README.md << 'EOL'
# Feast Friends API

A Go backend API for a social food sharing application.

## Project Structure

# 🍽️ Feast Friends API - Partner Onboarding Guide

Welcome to the Feast Friends API development team! This guide will get you up and running quickly.

## 👥 Team & Roles

### **Luiz** - Foundation & Infrastructure Lead
**Your Focus Areas:**
- 🔐 Authentication system (JWT, login/logout)
- ⚙️ Project configuration and environment setup
- 🛡️ Middleware (CORS, logging, auth validation)
- 🔧 Utilities (JWT handling, validation, response helpers)
- 🐳 Docker setup and deployment infrastructure
- 📊 Logging system and error handling

### **Gus** - Core Features & Social Lead  
**Your Focus Areas:**
- 📊 Data models (User, Post, Event, Message structures)
- 🎯 API handlers (posts, social features, events, messaging)
- 💼 Business logic services (post creation, social interactions)
- 🗄️ Database repository layer (data access patterns)
- 🖼️ Image upload and handling
- 🔍 Social features (likes, follows, comments)

---

## 🚀 Quick Start (5 minutes)

### 1. Clone the Repository
```bash
git clone https://github.com/YOUR_ORG/feast-friends-api.git
cd feast-friends-api
```

### 2. Set Up Your Development Environment
```bash
# Install Go (if not installed): https://golang.org/dl/
go version  # Should show Go 1.21+

# Install dependencies
make install

# Copy environment template
cp .env.example .env
```

### 3. Configure Supabase (Database)
1. Go to [supabase.com](https://supabase.com) and create account
2. Create new project: "feast-friends"
3. Go to Settings → API → Copy your:
   - `Project URL`
   - `anon/public key`
4. Update your `.env` file:
```env
SUPABASE_URL=your_project_url_here
SUPABASE_KEY=your_anon_key_here
JWT_SECRET=change-this-to-something-secure
```

### 4. Install VS Code + Extensions
Download [VS Code](https://code.visualstudio.com/) and install these extensions:
- **Go** (by Google) - Essential for Go development
- **REST Client** - Test APIs directly in VS Code
- **GitLens** - Enhanced Git features

### 5. Test Your Setup
```bash
# Run the application
make run

# Should see: "Starting server on port: 8080"
# Visit: http://localhost:8080/health
```

---

## 📁 Project Structure Explained

```
feast-friends-api/
├── cmd/server/main.go           # 🚀 App entry point (Luiz)
├── internal/
│   ├── config/config.go         # ⚙️ Environment setup (Luiz)
│   ├── handlers/                # 🎯 API endpoints (Gus)
│   │   ├── auth.go              # 🔐 Login/register (Luiz)
│   │   ├── posts.go             # 📝 Post CRUD (Gus)
│   │   ├── social.go            # ❤️ Likes/follows (Gus)
│   │   ├── events.go            # 📅 Events system (Gus)
│   │   ├── messages.go          # 💬 Direct messages (Gus)
│   │   └── users.go             # 👤 User profiles (Gus)
│   ├── middleware/              # 🛡️ Request processing (Luiz)
│   ├── models/                  # 📊 Data structures (Gus)
│   ├── services/                # 💼 Business logic (Gus)
│   ├── repository/              # 🗄️ Database access (Gus)
│   └── utils/                   # 🔧 Helper functions (Luiz)
└── ...
```

---

## 🔄 Daily Workflow

### Starting Your Work Day
```bash
# 1. Switch to develop branch and get latest code
git checkout develop
git pull origin develop

# 2. Create your feature branch
git checkout -b feature/[your-name]-[description]
# Examples:
git checkout -b feature/luiz-jwt-middleware
git checkout -b feature/gus-post-models

# 3. Start coding!
```

### During Development
```bash
# Check what you've changed
git status

# Add your changes
git add .

# Commit with descriptive message
git commit -m "feat: implement user authentication middleware"
git commit -m "fix: resolve post creation validation bug"
git commit -m "docs: update API endpoint documentation"

# Push to your branch
git push
```

### When Feature is Complete
1. **Push final changes**: `git push`
2. **Create Pull Request** on GitHub:
   - Go to repository → "New Pull Request"
   - Base: `develop` ← Compare: `feature/your-branch`
   - Fill in the PR template
   - Request review from your partner
3. **Address feedback** if needed
4. **Merge** after approval

---

## 🧪 Testing Your Code

### Manual Testing with REST Client (VS Code)
Create a file `api-tests.http`:
```http
### Health Check
GET http://localhost:8080/health

### Register User
POST http://localhost:8080/api/auth/register
Content-Type: application/json

{
  "email": "test@example.com",
  "username": "testuser",
  "password": "password123",
  "full_name": "Test User"
}

### Login
POST http://localhost:8080/api/auth/login
Content-Type: application/json

{
  "email": "test@example.com",
  "password": "password123"
}
```

### Running Automated Tests
```bash
# Run all tests
make test

# Run tests with coverage
make test-coverage

# Check code formatting
make fmt
```

---

## 📋 Your TODO Lists

### Luiz's Initial Tasks (Foundation)
- [ ] **Week 1**: Complete `cmd/server/main.go` - app startup and routing
- [ ] **Week 1**: Implement `internal/config/config.go` - environment configuration
- [ ] **Week 1**: Build `internal/middleware/auth.go` - JWT validation
- [ ] **Week 1**: Create `internal/utils/jwt.go` - token generation/validation
- [ ] **Week 2**: Finish authentication handlers in `internal/handlers/auth.go`
- [ ] **Week 2**: Add CORS and logging middleware
- [ ] **Week 3**: Docker setup and deployment preparation
- [ ] **Week 4**: Performance optimization and production readiness

### Gus's Initial Tasks (Core Features)
- [ ] **Week 1**: Define all data models in `internal/models/` 
- [ ] **Week 1**: Create database schema in `migrations/`
- [ ] **Week 2**: Implement post handlers in `internal/handlers/posts.go`
- [ ] **Week 2**: Build social features in `internal/handlers/social.go`
- [ ] **Week 3**: Add events system in `internal/handlers/events.go`
- [ ] **Week 3**: Create messaging system in `internal/handlers/messages.go`
- [ ] **Week 4**: Image upload service and final integrations

---

## 🆘 Getting Help

### Common Issues & Solutions

**"go: cannot find main module"**
```bash
# Make sure you're in the project root directory
cd feast-friends-api
go mod tidy
```

**"Supabase connection failed"**
- Check your `.env` file has correct `SUPABASE_URL` and `SUPABASE_KEY`
- Verify Supabase project is active

**"Port 8080 already in use"**
```bash
# Kill process using port 8080
sudo lsof -i :8080
sudo kill -9 [PID]
```

**Git conflicts when pulling**
```bash
# Stash your changes, pull, then reapply
git stash
git pull origin develop
git stash pop
```

### Communication Channels
- **Daily Standup**: Share what you worked on, what you're working on, any blockers
- **Code Reviews**: Always review each other's PRs thoroughly
- **Slack/Discord**: For quick questions and updates

---

## 📖 Learning Resources

### Go Development
- [Go Tour](https://tour.golang.org/) - Interactive Go tutorial
- [Gin Framework Docs](https://gin-gonic.com/docs/) - Our web framework
- [Go by Example](https://gobyexample.com/) - Practical Go examples

### API Development
- [REST API Best Practices](https://restfulapi.net/)
- [HTTP Status Codes](https://httpstatuses.com/)
- [Postman Learning Center](https://learning.postman.com/)

### Database & Supabase
- [Supabase Documentation](https://supabase.com/docs)
- [PostgreSQL Tutorial](https://www.postgresqltutorial.com/)

### Git & GitHub
- [Git Handbook](https://guides.github.com/introduction/git-handbook/)
- [GitHub Flow](https://guides.github.com/introduction/flow/)

---

## 🎯 Project Goals & Timeline

### Week 1-2: Foundation
- ✅ Project setup complete
- ✅ Authentication system working
- ✅ Basic user registration/login
- ✅ Database models defined

### Week 3-4: Core Features
- ✅ Posts creation and retrieval
- ✅ Social features (likes, follows)
- ✅ User profiles
- ✅ Basic API testing complete

### Week 5-6: Advanced Features
- ✅ Events system
- ✅ Direct messaging
- ✅ Image upload
- ✅ Integration testing

### Week 7-8: Polish & Deploy
- ✅ Performance optimization
- ✅ Security hardening
- ✅ Production deployment
- ✅ API documentation complete

---

## 🚨 Important Notes

### Security Reminders
- **Never commit secrets**: Check `.env` is in `.gitignore`
- **Validate all inputs**: Use our validation utilities
- **Sanitize database queries**: Prevent SQL injection
- **Use HTTPS in production**: Configure proper TLS

### Code Quality Standards
- **Go formatting**: Always run `go fmt ./...`
- **Descriptive commits**: Use conventional commit format
- **Test your code**: Write unit tests for your functions
- **Document APIs**: Add comments for exported functions

### Collaboration Best Practices
- **Communicate early**: Don't wait if you're blocked
- **Review thoroughly**: Check your partner's PRs carefully
- **Stay organized**: Update TODO comments as you complete tasks
- **Backup your work**: Commit and push frequently

---

## 🎉 Welcome to the Team!

You're now ready to start building an amazing social food app! Remember:

1. **Ask questions** - There are no stupid questions
2. **Commit often** - Small, frequent commits are better than large ones
3. **Test thoroughly** - Always test your endpoints before pushing
4. **Communicate actively** - Keep your partner informed of progress
5. **Have fun!** - We're building something cool together

### Quick Commands Reference
```bash
make run          # Start development server
make test         # Run tests
make fmt          # Format code
git status        # Check what's changed
git pull          # Get latest code
git push          # Upload your changes
```

**Happy coding! 🚀**

---

*Need help? Check the troubleshooting section above or reach out to your partner. We're in this together!*