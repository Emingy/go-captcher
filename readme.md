# Go captcher 
Go module for generate and verify captcha

## How to use
### Generate captcha
```go
    c := captcha{
        length      int
        variants    []string
    }
	c.generateCaptcha()
```

### Verify captcha
```go
    c := captcha{
        length      int
        variants    []string
    }
    c.generateCaptcha()
	value := "🐹"
	verified := c.verifyCaptcha(v)
```