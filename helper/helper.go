package helper

import (
	"crypto/md5"
	"fmt"
	uuid "github.com/satori/go.uuid"
	"time"
)

// GetMd5 md5加密
func GetMd5(s string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(s)))
}

func GetUUID() string {
	return uuid.NewV4().String()
}

func GetCurrentTime() int64 {
	return time.Now().Unix()
}
