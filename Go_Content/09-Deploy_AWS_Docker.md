# ✅ Module 9 — Docker & AWS Deployment (Go + MySQL + Nginx + HTTPS)

## 🎯 Learning Objectives

By the end of this module, students will be able to:

* Containerize a Go application using Docker
* Run MySQL in a container
* Use Docker Compose for orchestration
* Deploy containers on AWS EC2
* Use Nginx as a reverse proxy
* Secure the app with HTTPS (Let’s Encrypt)
* Connect domain name (optional)

---

## 9.1 Install Docker on EC2 (Amazon Linux 2)

```bash
sudo yum update -y
sudo yum install docker -y
sudo service docker start
sudo usermod -aG docker ec2-user
```

Log out and log back in.

Test:

```bash
docker --version
```

---

## 9.2 Create Dockerfile for Go API

`Dockerfile`:

```dockerfile
FROM golang:1.22-alpine

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN go build -o main

EXPOSE 8080

CMD ["./main"]
```

---

## 9.3 docker-compose.yml

```yaml
version: "3.8"

services:
  mysql:
    image: mysql:8
    restart: always
    environment:
      MYSQL_ROOT_PASSWORD: root
      MYSQL_DATABASE: contacts_db
    ports:
      - "3306:3306"
    volumes:
      - db-data:/var/lib/mysql

  app:
    build: .
    restart: always
    ports:
      - "8080:8080"
    depends_on:
      - mysql
    environment:
      DB_HOST: mysql
      DB_USER: root
      DB_PASS: root
      DB_NAME: contacts_db

volumes:
  db-data:
```

⚠️ Update your Go DSN:

```go
dsn := os.Getenv("DB_USER") + ":" + 
       os.Getenv("DB_PASS") + "@tcp(" +
       os.Getenv("DB_HOST") + ":3306)/" +
       os.Getenv("DB_NAME")
```

---

## 9.4 Build and Run

```bash
docker-compose up --build
```

Test:

```
http://EC2_PUBLIC_IP:8080
```

---

## 9.5 Install Nginx on EC2

```bash
sudo yum install nginx -y
sudo systemctl start nginx
sudo systemctl enable nginx
```

---

## 9.6 Configure Nginx as Reverse Proxy

Edit:

```bash
sudo nano /etc/nginx/nginx.conf
```

Add inside server block:

```nginx
location / {
  proxy_pass http://localhost:8080;
}
```

Restart Nginx:

```bash
sudo systemctl restart nginx
```

Now test:

```
http://EC2_PUBLIC_IP
```

---

## 9.7 Add HTTPS (Let’s Encrypt)

Only works with a domain name:

```bash
sudo amazon-linux-extras install epel -y
sudo yum install certbot python3-certbot-nginx -y

sudo certbot --nginx
```

Follow the prompts.

✅ Free SSL
✅ Auto renewal
✅ Secure hosting

---

# ✅ Final Architecture

```
User Browser
    ↓
 HTTPS (Nginx)
    ↓
 Go App Container (API + HTMX)
    ↓
 MySQL Container
```

---

## ✅ Capstone Student Project

Each student must:

✅ Deploy their app to AWS
✅ Secure with HTTPS
✅ Create test users
✅ Demonstrate CRUD
✅ Show Docker containers running
✅ Show database persistence

**Extra credit:**

* Add S3 for file uploads
* Add logging with Grafana
* Add Redis caching
* Add CI/CD with GitHub Actions

---
