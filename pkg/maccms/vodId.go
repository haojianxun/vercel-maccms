package maccms

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"testing"
	"time"
)

// 打乱字符集
func shuffleCharset(charset string) string {
	runes := []rune(charset)
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(runes), func(i, j int) {
		runes[i], runes[j] = runes[j], runes[i]
	})
	return string(runes)
}

// 随机 [A-Z]+[a-z]+[0-9]
const charset = "TRzMb1CKONuAyFnj58XP4EDUri3owWBHpLfv79x6Yq0sJlt2SkhVZeIcaQgmGd"

// 加密ID为6位字符串

func EncryptID(id int) string {
	if id < 1 || id > 999999 {
		return "" // 超出范围
	}

	// 将ID转换为字符串
	idStr := strconv.Itoa(id)
	// 如果不足6位，在前面补0
	for len(idStr) < 6 {
		idStr = "0" + idStr
	}

	encrypted := ""
	for i := 0; i < 6; i++ {
		index := idStr[i] - '0' // 将字符转换为对应的数字
		encrypted += string(charset[index])
	}
	return encrypted
}

// 解密6位字符串还原成ID

func DecryptID(encryptedID string) int {
	if len(encryptedID) != 6 {
		panic("invalid encrypted ID length")
	}
	idStr := ""
	for _, char := range encryptedID {
		index := strings.IndexRune(charset, char)
		if index == -1 {
			panic("invalid character in encrypted ID")
		}
		idStr += strconv.Itoa(index)
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		panic(err)
	}
	return id
}

func _Test(t *testing.T) {
	fmt.Println(shuffleCharset(charset))
	// 加密ID
	id := 1
	encryptedID := EncryptID(id)
	fmt.Println("加密后的ID:", encryptedID)

	// 解密字符串
	decryptedID := DecryptID(encryptedID)

	fmt.Println("解密后的ID:", decryptedID)
}
