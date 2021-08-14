package main

import (
	"fmt"
	"net/url"
)

func main() {
	str := "http://weibo.com"
	u, _ := url.Parse(str)
	u.Query().Add("yangfan", "ab")
	fmt.Println(u.String())
}
