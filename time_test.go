package box

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

// go的time时间根据模版转换为字符串日期的测试
// 同时可以当作example
func TestDateT(t *testing.T) {
	now := time.Now()
	got := DateT(now, "YYYY-MM-DD")

	if got == "" {
		t.Errorf("expected: , got:%v", got)
	}
}

// unix的int64时间滚句模版转换为日期字符串的测试
// 同时可以当作example
func TestDateI(t *testing.T) {
	unixTime := int64(1646275103)
	want := "2022-03-03"
	got := DateI(unixTime, "YYYY-MM-DD")

	if got != want {
		t.Errorf("expected:%v, got:%v", want, got)
	}
}

// 注意：这个方法的日期模版是根据【datePatterns】里定义的，与其他日期转换方法不同！
func TestDateParse(t *testing.T) {
	now := "2020-03-03 11:03:20"
	got, err := DateParse(now, "Y-m-d H:i:s")

	if err != nil {
		t.Errorf("expected: , got:%v", got)
	}
}

func ExampleDateT() {
	t := time.Now()
	// 模版可参考【DateT】的注释
	t1 := DateT(t, "YYYY-MM-DD")
	t2 := DateT(t, "YYYY-MM-DD HH:mm:ss")
	fmt.Println(t1)
	fmt.Println(t2)

}

func ExampleDateI() {
	unixt := time.Now().Unix()

	// 模版可参考【DateT】的注释
	t1 := DateI(unixt, "YYYY-MM-DD")
	t2 := DateI(unixt, "YYYY-MM-DD HH:mm:ss")
	fmt.Println(t1)
	fmt.Println(t2)

}

func ExampleDateS() {
	unixt := time.Now().Unix()
	unixts := strconv.FormatInt(unixt, 10)

	// 模版可参考【DateT】的注释
	t1 := DateS(unixts, "YYYY-MM-DD")
	t2 := DateS(unixts, "YYYY-MM-DD HH:mm:ss")
	fmt.Println(t1)
	fmt.Println(t2)

}

func ExampleDateParse() {
	ts := "2020-07-01 17:21:30"
	p := "Y-m-d H:i:s"

	t, err := DateParse(ts, p)

	if err != nil {
		fmt.Errorf("err msg: %s", err)
	}

	fmt.Println(t)
}
