package utils

import (
	"io"
	"fmt"
	"crypto/md5"
)

func DoubleMd5(pw string) string{
	h:=md5.New()
	io.WriteString(h,pw)
	pwmd5:=fmt.Sprintf("%x",h.Sum(nil))
	h2:=md5.New()
	io.WriteString(h2,pwmd5)
	return fmt.Sprintf("%x",h2.Sum(nil))
}
