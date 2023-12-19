# Go Gin

## Project Structure

- https://stackoverflow.com/questions/57024470/folder-structure-and-package-naming-convention-for-a-rest-api-develop-in-gin-fra
- https://github.com/RezaOptic/gin-project-structure
- https://medium.com/@poonnadapattarapadya_93995/%E0%B8%AA%E0%B8%A3%E0%B9%89%E0%B8%B2%E0%B8%87-restfulapi-%E0%B8%94%E0%B9%89%E0%B8%A7%E0%B8%A2-golang-gin-grom-postgresql-by-heroku-d8ce149497d8
- https://medium.com/thipwriteblog/golang-live-reload-4e1f96648b80
- Use air for live reload: https://github.com/cosmtrek/air

## Go "package"

- import `"gin-app/routers"`: means
- this project (see: `go.mod`) > `routers` folder

By convention, `package name` == `folder name` (we import `path/to/folder`)

```go
// something/hello.go
package something

// define Hello function
```

To use the `something` package, we need to import it

```go
import "path/to/something"

something.Hello()
```
