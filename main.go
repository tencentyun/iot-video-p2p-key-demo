package main

import (
	"crypto/sha256"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	guuid "github.com/google/uuid"
)

func validateConfigPath(path string) error {
	s, err := os.Stat(path)
	if err != nil {
		return err
	}
	if s.IsDir() {
		return fmt.Errorf("'%s' is a directory, not a normal file", path)
	}
	return nil
}

func parseFlags() (string, error) {
	var configPath string
	flag.StringVar(&configPath, "config", "./config.yml", "path to config file")
	flag.Parse()
	if err := validateConfigPath(configPath); err != nil {
		return "", err
	}

	return configPath, nil
}

func main() {
	configPath, err := parseFlags()
	if err != nil {
		panic(err)
	}
	config, err := NewConfig(configPath)
	if err != nil {
		panic(err)
	}

	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	client := redis.NewClient(&redis.Options{
		Addr:     config.Redis.Addr,
		Password: config.Redis.Pass,
	})

	router.GET("/code/:code", func(c *gin.Context) {
		code := c.Param("code")
		key, err := client.Get(c, code).Result()
		if err == redis.Nil {
			c.AbortWithStatus(404)
			return
		} else if err != nil {
			c.AbortWithStatus(500)
			return
		}
		c.JSON(200, gin.H{"key": key})
	})

	router.POST("/code", func(c *gin.Context) {
		code := guuid.New().String()
		h := sha256.New()
		h.Write([]byte(code + fmt.Sprint(time.Now().UnixNano()) + fmt.Sprint(rand.Float64())))
		key := fmt.Sprintf("%x", h.Sum(nil))
		expire := time.Second * time.Duration(config.Redis.Expire)
		err := client.Set(c, code, key, expire).Err()
		if err != nil {
			c.AbortWithStatus(500)
			return
		}
		c.JSON(200, gin.H{
			"code":   code,
			"key":    key,
			"expire": time.Now().Add(expire).Unix(),
		})
	})

	router.Run(config.Service.Addr)
}
