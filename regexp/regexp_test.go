package regexp

import (
	"fmt"
	"testing"
)

func TestIsUserID(t *testing.T) {

	_is := IsUserID("a_1")
	if !_is {
		fmt.Println("错误")
	}

}
func TestIsEmail(t *testing.T) {

	_is := IsEmail("ministor@126.com")

	if !_is {
		t.Errorf("Isn't Mail")
	}

	_is = IsEmail("ministor11@126.cn.com")

	if !_is {
		t.Errorf("Isn't Mail")
	}

	_is = IsEmail("minis@tor11@126.cn.com")

	if _is {
		t.Errorf("Isn't Mail")
	}
}

func TestIsMobile(t *testing.T) {

	_is := IsMobile("13811234616")

	if !_is {
		t.Errorf("Isn't Mobile")
	}

}

func TestIsStrongPwd(t *testing.T) {

	_is, _msg := IsStrongPwd("a11111")
	fmt.Println(_is)
	if !_is {
		fmt.Println(_msg)
	}

}

func TestIsIP(t *testing.T) {

	_is := IsIP("172.0.0.1")
	fmt.Println(_is)

}

func TestMobileReplaceRepl(t *testing.T) {

	_str := MobileReplaceRepl("13811234616")

	fmt.Println(_str)
	if _str != "138*****616" {
		t.Errorf("MobileReplaceRepl")
	}
}
