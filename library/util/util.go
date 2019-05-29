package util

import "github.com/vsiryxm/gin-api/library/setting"

// Setup Initialize the util
func Setup() {
	jwtSecret = []byte(setting.AppSetting.JwtSecret)
}