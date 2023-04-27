# Go captcher 
Go module for generate and verify captcha

## How install
```bash
    go get github.com/Emingy/go-captcher
```
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