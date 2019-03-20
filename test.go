package testmod

import "fmt"

func Test(key, value string) string {
	return fmt.Sprintf("input key:%s, value:%s, in module test", key, value)
}
