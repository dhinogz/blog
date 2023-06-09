# Personal Blog
A personal website and blog built with Go. 

## Installation

#### Pre-requisites
- Docker
- Docker Compose
- [Migrate library](https://github.com/golang-migrate/migrate)
- [LazyDocker](https://github.com/jesseduffield/lazydocker) (Optional)


#### 1. Copy .env.example file to .env file
```bash
$ cp .env.example .env
```

#### 2. Build Docker images
```bash
$ docker compose build
```

#### 3. Start Docker containers
```bash
$ docker compose up -d
```

#### 4. Make migrations
```bash
$ export BLOG_DB_DSN=(dsn can be found in .env.example)
$ migrate -path=./migrations -database=$BLOG_DB_DSN up
```

Go to localhost:4000 and it should be working.

#### Manage containers
```bash
$ lazydocker
```

## TODO
- [ ] Authentication
  - [ ] Find out best way to go through these requirements
    - [ ] Common users cannot find login url
    - [ ] Bots scraping the internet cannot find login url
    - [ ] Only I, The Lord Supreme, can access this view.
- [ ] Create a POST articles views. User needs to be authenticated to view
- [ ] Middleware
- [ ] Add more projects with more detail
- [ ] Deploy to Google Cloud