package http

import (
	"fmt"
	"net/url"
	"strings"
	"testing"
	"time"
)

func TestPostForm(t *testing.T) {

	urlStr := "http://localhost:3000/post"
	form := make(url.Values)
	form.Set("email", "ministor@126.com")
	form.Set("pwd", "a11111")
	_string := form.Encode()
	fmt.Println(_string)
	_ok, _ := PostForm(urlStr, form)

	if _ok != "OK" {
		t.Errorf("PostForm error")
	}
}

func TestPostFormTimeout(t *testing.T) {

	urlStr := "https://baidu.com"
	form := make(url.Values)
	form.Set("email", "ministor@126.com")
	form.Set("pwd", "a11111")
	_string := form.Encode()
	fmt.Println(_string)
	_ok, err := httpPostThree(urlStr, form)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(_ok)
	}
}

func httpPostThree(url string, data url.Values) (string, error) {

	fmt.Println(url)

	fmt.Println("post first")
	resp, err := PostFormTimeout(url, data)
	if err == nil {
		return resp, nil
	}
	fmt.Println("httpPostThree 1", err.Error())

	time.Sleep(1 * time.Second)

	fmt.Println("post second")
	resp, err = PostFormTimeout(url, data)
	if err == nil {
		return resp, nil
	}
	fmt.Println("httpPostThree 2", err.Error())

	time.Sleep(1 * time.Second)

	fmt.Println("post three")
	resp, err = PostFormTimeout(url, data)
	if err == nil {
		return resp, nil
	}

	fmt.Println("httpPostThree 3", err.Error())
	return "", err
}

func TestPostFormTimeoutThree(t *testing.T) {

	urlStr := "https://baiduss.com"
	form := make(url.Values)
	form.Set("email", "ministor@126.com")
	form.Set("pwd", "a11111")
	_string := form.Encode()
	fmt.Println(_string)
	_ok, err := httpPostThree(urlStr, form)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(_ok)
	}
}

func TestPostJSONTimeout(t *testing.T) {

	urlStr := "http://127.0.0.1:1193/glut/tx/history"
	_string := `{
		"addr": "0x2a3b954c3f555c8d3903afe9422b0efd8f56fbff",
		"page": 0,
		"api_token": "28b7b572ffac749759643c2c268e4a98"
	}`
	fmt.Println(_string)
	_ok, err := httpPostJSONThree(urlStr, _string)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(_ok)
	}
}

func httpPostJSONThree(url string, jsonPost string) (string, error) {

	fmt.Println(url)

	fmt.Println("post first")
	resp, err := PostJSONTimeout(url, jsonPost)
	if err == nil {
		return resp, nil
	}
	fmt.Println("httpPostThree 1", err.Error())

	time.Sleep(1 * time.Second)

	fmt.Println("post second")
	resp, err = PostJSONTimeout(url, jsonPost)
	if err == nil {
		return resp, nil
	}
	fmt.Println("httpPostThree 2", err.Error())

	time.Sleep(1 * time.Second)

	fmt.Println("post three")
	resp, err = PostJSONTimeout(url, jsonPost)
	if err == nil {
		return resp, nil
	}

	fmt.Println("httpPostThree 3", err.Error())
	return "", err
}

func TestPost(t *testing.T) {

	url := "http://localhost:3000/post"
	bodyType := "application/x-www-form-urlencoded"
	body := strings.NewReader("email=ministor@126.com&pwd=b11111")

	_ok, _ := Post(url, bodyType, body)

	if _ok != "OK" {
		t.Errorf("Post error")
	}
}
