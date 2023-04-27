package gocaptcher

import (
	"math/rand"
	"time"
)

type Captcha struct {
	Length      int
	Variants    []string
	TrueVariant string
}

func (c *Captcha) GenerateCaptcha() {
	config := Config{}
	config.initConfig()
	min := 0
	max := len(config.emojis)
	count := c.Length

	rand.Seed(time.Now().UnixNano())

	nums := make(map[int]bool)
	for len(nums) < count {
		n := rand.Intn(max-min+1) + min
		nums[n] = true
	}

	for n := range nums {
		c.Variants = append(c.Variants, config.emojis[n])
	}

	c.TrueVariant = c.Variants[rand.Intn((count-1)-0+1)+0]
}

func (c *Captcha) VerifyCaptcha(v string) bool {
	if v == c.TrueVariant {
		return true
	}
	return false
}
