package ys

import (
	"crypto/md5"
	"fmt"
	"github.com/tidwall/sjson"
	"math"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func getDS(q string, b map[string]any) string {
	br := ""
	if b != nil {
		for k, v := range b {
			br, _ = sjson.Set(br, k, v)
		}
	}
	s := "xV8v4Qu54lUKrEYFZkJhB8cuOh9Asafs"
	t := time.Now().Unix()
	rand.Seed(time.Now().UnixNano())
	r := rand.Intn(100001) + 100000
	c := md5.New()
	c.Write([]byte(fmt.Sprintf("salt=%s&t=%d&r=%d&b=%s&q=%s", s, t, r, br, q)))
	return fmt.Sprintf("%d,%d,%x", t, r, c.Sum(nil))
}

func getOldDS(mysbbs bool) string {
	n := "h8w582wxwgqvahcdkpvdhbh2w9casgfl"
	if mysbbs {
		n = "dWCcD2FsOUXEstC5f9xubswZxEeoBOTc"
	}
	t := time.Now().Unix()
	r := randStr(6)
	c := md5.New()
	c.Write([]byte(fmt.Sprintf("salt=%s&t=%d&r=%s", n, t, r)))
	return fmt.Sprintf("%d,%s,%x", t, r, c.Sum(nil))
}

func randStr(n int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyz0123456789")
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func getServer(uid int64) string {
	if strconv.FormatInt(uid, 10)[0:1] == "5" {
		return "cn_qd01"
	} else {
		return "cn_gf01"
	}
}

func getRandomHex(length int) string {
	rand.Seed(time.Now().UnixNano())
	r := rand.Int63n(int64(math.Pow(2, float64(length))))
	result := strings.ToUpper(strconv.FormatInt(r, 16))
	if len(result) < length {
		result = strings.Repeat("0", length-len(result)) + result
	}
	return result
}

func resolveTime(str string) string {
	s, _ := strconv.Atoi(str)
	m := s / 60
	h := m / 60
	return fmt.Sprintf("%s:%s:%s", completionTime(h), completionTime(m-h*60), completionTime(s-m*60))
}

func completionTime(t int) string {
	if t < 10 {
		return fmt.Sprintf("0%d", t)
	}
	return fmt.Sprintf("%d", t)
}
