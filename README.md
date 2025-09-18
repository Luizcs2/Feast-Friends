cat > README.md << 'EOL'
# Feast Friends APP
# An app where you can share what you make socail events tailored to foodies and follow and like others along with ai features that calculate calories on what you cook based on the ingerdients you add into recipe

A Go backend API for a social food sharing application.

## Project Structure

# 🍽️ Feast Friends API - Partner Onboarding Guide

Welcome to the Feast Friends APP development team! This guide will get you up and running quickly.

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

[![CI/CD Pipeline Status](https://github.com/Luizcs2/Feast-Friends/actions/workflows/ci.yml/badge.svg)](https://github.com/Luizcs2/Feast-Friends/actions)

This project is configured with a complete CI/CD pipeline using GitHub Actions to automate testing, security scanning, and deployment.

---

## 🤖 CI/CD and Automation Overview

This repository contains a set of files that automate our development lifecycle. Here’s a quick guide to what each file does:

### Workflows (`.github/workflows/`)

* **`pr-check.yml`**: Runs automatically on every **pull request**. It lints the code and runs all unit tests to ensure code quality before merging.
* **`ci.yml`**: Runs automatically when code is **pushed to the `main` branch**. It serves as a final check to ensure the main branch is always stable.
* **`security.yml`**: Runs on pull requests and pushes to `main`. It uses GitHub CodeQL to scan the code for common security vulnerabilities.
* **`cd.yml`**: The Continuous Deployment workflow. It triggers automatically when a new **version tag** (e.g., `v1.0.0`) is pushed. It builds a production-ready Docker image and pushes it to a container registry.

---

### Dockerization (`/docker/`)


* **`Dockerfile.dev`**: The blueprint for building the Go application. It uses a multi-stage build to create a small, efficient, and secure production image.
* **`docker-compose.prod.yml`**: A file to easily run the entire application stack (the API and a PostgreSQL database) in a production-like environment.
* **`.dockerignore`**: Lists files and folders that should be ignored by Docker during the build process to keep the image clean and small.

---

### Scripts & Configuration
* **``** : hello

* **`scripts/smoke-tests.sh`**: A simple post-deployment script to check if the application is running and healthy.
* **`.golangci.yml`**: The configuration file for our linter, `golangci-lint`. It defines the rules for keeping our Go code consistent and clean.
* **`.github/` templates**: The markdown files in `ISSUE_TEMPLATE`, along with `PULL_REQUEST_TEMPLATE.md`, provide a consistent structure for creating issues and pull requests.

---

## 🚀 How the Workflow Works

1.  **Development**: A developer creates a pull request.
2.  **Code Review**: The `pr-check.yml` and `security.yml` workflows automatically run tests and security scans. GitHub requires these to pass before merging.
3.  **Merge**: Once approved, the code is merged into the `main` branch, triggering the `ci.yml` workflow.
4.  **Release**: To deploy a new version, a maintainer pushes a new tag (e.g., `git tag v1.0.1` && `git push origin v1.0.1`).
5.  **Deployment**: Pushing the tag triggers the `cd.yml` workflow, which builds the final Docker image and pushes it to the container registry, ready for deployment.


**Happy coding! 🚀**

---

*Need help? Check the troubleshooting section above or reach out to your partner. We're in this together!*
