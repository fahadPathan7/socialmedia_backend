# 🌐 Social Media Backend

This project is a backend service for a social media application. It provides APIs for user management, post creation, comments, and reactions. The backend is built using **Go**, **gRPC**, and **MongoDB**. The APIs are documented using **Swagger**. The project is structured into multiple microservices.

## ✨ Features

- 👤 **User Management**: Registration and authentication.
- 📝 **Post Management**: Create, retrieve, update, and delete posts.
- 💬 **Comment Management**: Add, view, update, and delete comments on posts.
- ❤️ **Reaction Management**: Add, view, update, and delete reactions to posts.
- 📜 **API Documentation**: Interactive documentation using **Swagger**.

---

## ⚙️ Installation

### ✅ Prerequisites

- 🐹 **Go**: Version 1.16 or later.
- 🍃 **MongoDB**: A running MongoDB instance.
- 🛠 **Protobuf Compiler**: `protoc` installed.
- 🔌 **Protobuf Plugins**:
  - `protoc-gen-go`
  - `protoc-gen-go-grpc`
  - `protoc-gen-grpc-gateway`
  - `protoc-gen-openapiv2`

---

### 🛠 Steps

1. **📂 Clone the repository**

    ```sh
    git clone https://github.com/yourusername/socialmedia_backend.git
    cd socialmedia_backend
    ```

2. **📦 Install dependencies**

    ```sh
    go mod tidy
    ```

3. **🔧 Install `protoc` plugins**

    ```sh
    go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
    go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
    go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest
    go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest
    ```

4. **🔧 Configure environment variables**

    Create a `.env` file in each microservice directory. Follow the `.env example` file for reference.

5. **▶️ Run the server**

    Run the `main.go` file of each microservice. Use run/debug option or run the following command:

    ```sh
    cd <microservice directory>
    go run main.go
    ```

---

## 📖 API Documentation

The API documentation is accessible at:

```
http://localhost:<port>/docs
```


Replace `<port>` with the port number of the running microservice.

---

## 🚀 Operations

### 👤 User Management

- **Register**: Create a new user.
- **Login**: Authenticate an existing user.

### 📝 Post Management

- **Create**: Add a new post.
- **Retrieve**: View a post by its ID.
- **Update**: Modify an existing post.
- **Delete**: Remove a post.

### 💬 Comment Management

- **Add Comment**: Attach a comment to a post.
- **View Comments**: List comments of a specific post.
- **Update Comment**: Edit a comment.
- **Delete Comment**: Remove a comment.

### ❤️ Reaction Management

- **Add Reaction**: React to a post.
- **View Reactions**: List reactions for a specific post.
- **Update Reaction**: Modify a reaction.
- **Delete Reaction**: Remove a reaction.

---