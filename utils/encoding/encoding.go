package encoding

import (
	"bytes"
	"crypto/md5"
	crand "crypto/rand"
	"crypto/sha1"
	"encoding/base64"
	"encoding/hex"
	"math/big"
	"math/rand"
	"strconv"
	"strings"
)

var seeds = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz1234567890")

const (
	base64Seeds = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"
)

// 根据长度创建随机字符串
func CreateRandomString(len int) string {
	var container string
	var str = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
	b := bytes.NewBufferString(str)
	length := b.Len()
	bigInt := big.NewInt(int64(length))
	for i := 0; i < len; i++ {
		randomInt, _ := crand.Int(crand.Reader, bigInt)
		container += string(str[randomInt.Int64()])
	}
	return container
}

// 创建uuid
func UUID() string {
	var buffer = make([]rune, 30)
	for i := 0; i < 30; i++ {
		index := rand.Int31n(30)
		buffer[i] = seeds[index]
	}
	return string(buffer)
}

// 将字符串编码为ascii
func EncodeASCII(input string) string {
	ret := strconv.QuoteToASCII(input)
	return ret[1 : len(ret)-1]
}

// 解码EncodeASCII()
func DecodeASCII(input string) (string, error) {
	ret, err := strconv.Unquote("\"" + input + "\"")
	return ret, err
}

// 将字符串编码为Base64
func EncodeBase64(input string) string {
	var coder = base64.NewEncoding(base64Seeds)
	_input := []byte(input)
	return coder.EncodeToString(_input)
}

// 解码Base64
func DecodeBase64(input string) (string, error) {
	var coder = base64.NewEncoding(base64Seeds)
	ret, err := coder.DecodeString(input)
	return string(ret), err
}

// md5加密
func EncryptMd5(input string) string {
	md5Ctx := md5.New()
	md5Ctx.Write([]byte(input))
	cipherStr := md5Ctx.Sum(nil)
	encryptedData := hex.EncodeToString(cipherStr)
	return encryptedData
}

// 截取字符串中间部分
func GetBetweenStr(str, start, end string) string {
	n := strings.Index(str, start)
	if n == -1 {
		n = 0
	} else {
		n = n + len(start)
	}
	str = string([]byte(str)[n:])
	m := strings.Index(str, end)
	if m == -1 {
		m = len(str)
	}
	str = string([]byte(str)[:m])
	return str
}

func SHA1(s string) string {
	o := sha1.New()
	o.Write([]byte(s))
	return hex.EncodeToString(o.Sum(nil))
}
