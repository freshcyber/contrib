package regexp

import (
	"regexp"
)

// IsUserID IsUserID
func IsUserID(userid string) bool {
	if userid == "" {
		return false
	}
	ok, _ := regexp.MatchString(`^[a-z0-9_-]{3,16}$`, userid)
	return ok
}

// IsIP 验证一个输入是不是IP地址
func IsIP(ip string) bool {
	if m, _ := regexp.MatchString("^[0-9]{1,3}\\.[0-9]{1,3}\\.[0-9]{1,3}\\.[0-9]{1,3}$", ip); !m {
		return false
	}
	return true
}

// IsStrongPwd 是否是合格的密码
func IsStrongPwd(pwd string) (bool, string) {
	if pwd == "" {
		return false, "密码不能为空"
	}
	if len(pwd) < 6 {
		return false, "密码至少6位"
	}
	// 弱密码 ^(.{6,}).*
	ok, _ := regexp.MatchString(`^[a-zA-Z0-9_-]{4,16}$`, pwd)
	if !ok {
		return false, "密码强度不足"
	}
	return true, "密码合规"
}

// IsEmail 是否是email
func IsEmail(email string) bool {
	if email == "" {
		return false
	}
	ok, _ := regexp.MatchString(`^([a-zA-Z0-9]+[_|\_|\.|\-]?)*[_a-z\-A-Z0-9]+@([a-zA-Z0-9]+[_|\_|\.|\-]?)*[a-zA-Z0-9\-]+\.[0-9a-zA-Z]{2,6}$`, email)
	return ok
}

// IsUsername 是否只包含数字, 字母 -, _
func IsUsername(username string) bool {
	if username == "" {
		return false
	}
	ok, _ := regexp.MatchString(`[^0-9a-zA-Z_\-]`, username)
	return !ok
}

const (
	regularMobile = "^1([38][0-9]|14[57]|5[^4])\\d{8}$"
)

// IsMobile check if mobile
func IsMobile(mobileNum string) bool {
	reg := regexp.MustCompile(regularMobile)
	return reg.MatchString(mobileNum)
}

// MobileReplaceRepl MobileReplaceRepl 手机号脱敏
func MobileReplaceRepl(str string) string {
	re, _ := regexp.Compile("(\\d{3})(\\d{5})(\\d{3})")
	return re.ReplaceAllString(str, "$1*****$3")
}
