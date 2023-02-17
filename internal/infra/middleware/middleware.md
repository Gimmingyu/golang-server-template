# MIDDLEWARE

아래 예시를 참조하여 커스텀 미들웨어를 파일로 구분하여 구현합니다.

```go
func DefaultMiddleware(config ...fiber.Config) fiber.Handler {
    return func(c *fiber.Ctx) error {
        return c.Next()
    }
}
```

각 미들웨어는 라우터 단위로 적용합니다. 