package main

import (
	"math/rand"
	"time"
)

type Captcher interface {
	generateCaptcha()
	verifyCaptcha()
}

type captcha struct {
	length      int
	variants    []string
	trueVariant string
}

func (c *captcha) generateCaptcha() {
	config := Config{}
	config.initConfig()
	min := 0
	max := len(config.emojis)
	count := c.length

	rand.Seed(time.Now().UnixNano())

	nums := make(map[int]bool)
	for len(nums) < count {
		n := rand.Intn(max-min+1) + min
		nums[n] = true
	}

	for n := range nums {
		c.variants = append(c.variants, config.emojis[n])
	}

	c.trueVariant = c.variants[rand.Intn((count-1)-0+1)+0]
}

func (c *captcha) verifyCaptcha(v string) bool {
	if v == c.trueVariant {
		return true
	}
	return false
}
