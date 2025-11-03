# SDE Intern Progress Report - Go Learning

**Date:** Nov 3, 2025  
**Intern:** Prem Modha  
**Focus:** Go API Development  
**Repository:** https://github.com/PremModhaOfficial/go-bookstore

## Summary
Built two RESTful APIs in Go, mastering concurrency, HTTP handling, and ORM setup.

## Projects Completed

### 1. Bookstore API
- **Tech:** Go 1.22.2, Gorilla Mux, GORM, SQLite
- **Features:**
  - CRUD endpoints: POST/GET/PUT/DELETE /book
  - Book model: ID, Name, Author, Publication, timestamps
  - Soft deletes, JSON handling, error responses
- **Architecture:**
  - Packages: models, controllers, routes, config, utils
  - Thread-safe in-memory storage with RWMutex

### 2. Movie CRUD Web Server
- **Tech:** Go HTTP server
- **Features:**
  - Movie CRUD operations
  - Web server setup
- **Deliverables:**
  - Progress reports (HTML/MD)
  - Project documentation

## Technical Learnings
- Go basics: structs, methods, maps, concurrency
- HTTP: routing, parsing, responses
- Tools: jj VCS, dependency management
- Code: modular packages, clean structure

## Today's Commits
1. 15:45 - "web server": Movie API + docs
2. 19:34 - "bookstore": Full API implementation
3. 19:43 - Naming fixes, updates

## Challenges Solved
- Thread-safe operations
- HTTP error handling
- Package organization