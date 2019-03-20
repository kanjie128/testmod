package testmod

import "fmt"

func Test(s string) string {
	return fmt.Sprintf("input %s in module test", s)
}
