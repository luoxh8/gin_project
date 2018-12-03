package h_categories

import (
	`encoding/base64`
	`fmt`
	`math/rand`
	`regexp`
	`strings`
	`time`
)

func GetIP(s string) string {
	if len(s) == 0 {
		return ""
	}
	ip := strings.Split(s, ":")
	return ip[0]
}

func base64Encode(src []byte) []byte {
	return []byte(base64.StdEncoding.EncodeToString(src))
}

func GetFormatCode() string {
	const shortForm = "2006-01-01 15:04:05"
	t := time.Now()
	temp := time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), time.Local)
	str := temp.Format(shortForm)
	var valid = regexp.MustCompile("[0-9]")
	timeStr := strings.Join(valid.FindAllString(str, -1), "")
	r := rand.New(rand.NewSource(t.UnixNano()))
	codeStr := fmt.Sprintf("%s%d", timeStr, r.Intn(100))
	return codeStr
}

func Reverse(s string) string {
	n := len(s)
	runes := make([]rune, n)
	for _, i := range s {
		n--
		runes[n] = i
	}
	return string(runes[n:])
}
