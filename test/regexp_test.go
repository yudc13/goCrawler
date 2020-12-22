package test

import (
	"fmt"
	"regexp"
	"testing"
)

func TestRegexp(t *testing.T) {
	const text = `
my email is yudachao@qq.com
tom email is tome@gmail.com
`
	re := regexp.MustCompile(`([a-zA-Z0-9]+)@[a-zA-Z0-9]+\.[a-zA-Z0-9]+`)
	//match := re.FindAllString(text, -1)
	match := re.FindAllStringSubmatch(text, -1)
	fmt.Println(match)
}
