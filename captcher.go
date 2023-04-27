package gocaptcher

import (
	"github.com/joho/godotenv"
	"math/rand"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type Captcha struct {
	Length      int
	Variants    []string
	TrueVariant string
}

func (c *Captcha) GenerateCaptcha() {
	path, _ := os.Getwd()
	envPath := filepath.Join(path, "/config.env")
	godotenv.Load(envPath)
	env, _ := os.LookupEnv("EMOJIS")
	emojis := strings.Split(env, ",")
	min := 0
	max := len(emojis)
	count := c.Length

	rand.Seed(time.Now().UnixNano())

	nums := make(map[int]bool)
	for len(nums) < count {
		n := rand.Intn(max-min+1) + min
		nums[n] = true
	}

	for n := range nums {
		c.Variants = append(c.Variants, emojis[n])
	}

	c.TrueVariant = c.Variants[rand.Intn((count-1)-0+1)+0]
}

func (c *Captcha) VerifyCaptcha(v string) bool {
	if v == c.TrueVariant {
		return true
	}
	return false
}
