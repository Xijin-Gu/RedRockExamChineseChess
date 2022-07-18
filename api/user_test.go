/**
* @Author: gxj
* @Data: 2022/7/18-15:05
* @DESC:
**/

package api

import (
	"encoding/json"
	"github.com/go-playground/assert/v2"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestRouter(t *testing.T) {
	tests := []struct{
		name string
		param string
		expect string
	}{
		{"bad",`{"name" : "","password" : ""}`,"亲，用户名或密码不能为空哦"},
		{"bad",`{"name" : "123456789123456","password" : "1234567"}`,"亲，用户名最长为15位哦"},
		{"bad",`{"name" : "123","password" : "1234"}`,"亲，密码最短为6位哦"},
		{"bad",`{"name" : "123","password" : "1234567"}`,"邮箱格式不对哟"},
	}

	r := Router(SendActivationCode)

	for _,v := range tests{
		t.Run(v.name, func(t *testing.T) {
			// mock一个HTTP请求
			req := httptest.NewRequest(
				"POST",
				"/test",
				strings.NewReader(v.param),
				)
			// mock一个响应记录器
			w := httptest.NewRecorder()


			// 让server端处理mock请求并记录返回的响应内容
			r.ServeHTTP(w,req)


			// 校验状态码是否符合预期
			assert.Equal(t,http.StatusOK,w.Code)


			// 解析并检验响应内容是否复合预期
			var resp map[string]string
			err := json.Unmarshal([]byte(w.Body.String()),&resp)
			assert.Equal(t,nil,err)
			assert.Equal(t,resp["message"],v.expect)
		})
	}


}


