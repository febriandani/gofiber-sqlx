[![forthebadge](https://forthebadge.com/images/badges/made-with-go.svg)](https://forthebadge.com)
[![forthebadge](https://forthebadge.com/images/badges/built-with-love.svg)](https://forthebadge.com)
[![forthebadge](https://forthebadge.com/images/badges/makes-people-smile.svg)](https://forthebadge.com)

# User CRUD App

This is a simple User CRUD Application built using Golang, Fiber, and SQLx.

# Features

- Create a new User
- Retrive a User by ID
- Update an Existing User
- Delete a User by ID

# Requirements

- GO (1.16+)
- Postgresql Database
- Fiber Framework
- SQLx
- Editor
  - [Vs Code](https://code.visualstudio.com/download)
  - [IntelliJ IDEA](https://www.jetbrains.com/idea/)
  - [Goland](https://www.jetbrains.com/go/)
  - [Atom](https://atom.io/)
  - [NotePad++](https://notepad-plus-plus.org/downloads/)

# User CRUD App

This is a simple User CRUD Application built using Golang, Fiber, and SQLx.

# Features

- Create a new User
- Retrive a User by ID
- Update an Existing User
- Delete a User by ID

# Requirements

- GO (1.16+)
- Postgresql Database
- Fiber Framework
- SQLx

# Installation

1. Clone the repository:

```bash
  git clone https://github.com/febriandani/gofiber-sqlx.git
```

2. Install or Update Dependencies:

```bash
go mod tidy
```

3. Run the application:

```bash
go run cmd/main.go
```

## API Endpoints Documentation

#### Create User

```http
  POST /users
```

| Body    | Type     | Description   |
| :------ | :------- | :------------ |
| `name`  | `string` | **Required**. |
| `email` | `string` | **Required**. |

#### Get User by ID

```http
  GET /users/:id
```

| params | Type      | Description   |
| :----- | :-------- | :------------ |
| `id`   | `integer` | **Required**. |

#### Update User

```http
  PUT /users
```

| Body    | Type     | Description   |
| :------ | :------- | :------------ |
| `name`  | `string` | **Required**. |
| `email` | `string` | **Required**. |

#### Delete User by ID

```http
  DELETE /users/:id
```

| params | Type      | Description   |
| :----- | :-------- | :------------ |
| `id`   | `integer` | **Required**. |

## Authors

- [@febriandani\_](https://www.linkedin.com/in/mhmmdfebriandani/)
