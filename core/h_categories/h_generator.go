package h_categories

import (
	`crypto/md5`
	`encoding/hex`
	`math/rand`
	`time`
)

func GenUid() string {
	str := "0123456789"
	bytes := []byte(str)
	var result []byte
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for {
		for i := 0; i < 7; i++ {
			result = append(result, bytes[r.Intn(len(bytes))])
		}
		if IsUid(string(result)) {
			break
		}
	}
	return string(result)
}

func GenGoodsId() string {
	var goodsId string
	for {
		goodsId = GenRandomString(64)
		if IsGoodsId(goodsId) {
			break
		}
	}
	return goodsId
}

func GenPassword(password string) string {
	return Reverse(GenMD5String(password))
}

func GenRandomString(length int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	var result []byte
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < length; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

func GenMD5String(str string) string {
	data := []byte(str)
	hash := md5.New()
	hash.Write(data)
	sum := hash.Sum(nil)
	encodeToString := hex.EncodeToString(sum)
	return encodeToString
}

func GenRandomInt(min int, max int) int {
	for {
		maxI := rand.Intn(max)
		if maxI > min {
			return maxI
		}
	}
}
