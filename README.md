# api-go

## Description

Web API for my apps *written in Go* for learning purposes.

[» Here is the API in PHP «](https://github.com/necodeus/api)

## Requirements

- [Go 1.21.4](https://go.dev/doc/install)

## GraphQL API Documentation

### Posts

Get all blog posts

```graphql
query {
  posts {
    id
    name
    content
  }
}
```

## REST API Documentation

### Posts

Get all blog posts

```
http://blog-api.localhost/posts
```

### Images

Get specific image by id

```
http://images.localhost/{uuidv4_id}
```
