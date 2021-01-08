package test

import (
	"fmt"
	"testing"
	"time"
)

func TestTimeToStr (t *testing.T){
	nt := time.Now()
	fmt.Println(nt)
	const base_format = "2006-01-02 15:04:05"
	str_time := nt.Format(base_format)
	fmt.Println(str_time)
	parse_str_time, _ := time.Parse(base_format, str_time)
	fmt.Println(parse_str_time)
	fmt.Println(parse_str_time.Format(base_format))

	dd, _ := time.ParseDuration("24h")
	dd1 := parse_str_time.Add(dd)
	fmt.Println(dd1)
}