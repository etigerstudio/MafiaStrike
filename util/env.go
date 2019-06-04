package util

import (
	"github.com/gin-gonic/gin"
	"mafia-strike/consts"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

func InitRandSeed()  {
	rand.Seed(time.Now().UTC().UnixNano())
}

func MustGetQuery(key string, c *gin.Context) string {
	if param, ok := c.GetQuery(key); ok {
		return param
	}

	Warnln("failed to get query:", key)
	ErrorMessageUniversal(c)

	c.Abort()
	panic(nil)
}

func MustGetPostForm(key string, c *gin.Context) string {
	if param, ok := c.GetPostForm(key); ok {
		return param
	}

	Warnln("failed to get post form:", key)
	ErrorMessageUniversal(c)

	c.Abort()
	panic(nil)
}

func MustGetParamInt(key string, c *gin.Context) int {
	if param, err := strconv.ParseInt(c.Param(key), 10, 64); err == nil {
		// TODO: prevent potential overflow
		return int(param)
	}

	Warnln("failed to get int parameter:", key)
	ErrorMessageUniversal(c)

	c.Abort()
	panic(nil)
}

func ErrorMessage(desc string, code int, c *gin.Context)  {
	c.String(code, desc)
}

func ErrorMessageUniversal(c *gin.Context) {
	ErrorMessage(consts.ErrorUniversal, http.StatusForbidden, c)
}

func RandomByteFrom(bytes string) byte {
	return bytes[rand.Intn(len(bytes))]
}

func GenerateRandomShortID() string {
	const idLetters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789-_"
	const idLength = 16

	filename := make([]byte, idLength)
	for i := range filename {
		filename[i] = RandomByteFrom(idLetters)
	}
	return string(filename)
}

func RemoveSliceElement(s []string, i int) []string {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}