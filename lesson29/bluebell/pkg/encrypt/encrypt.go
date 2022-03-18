/*
   @Author: StudentCWZ
   @Description:
   @File: encrypt
   @Software: GoLand
   @Project: GoWeb
   @Date: 2022/3/18 10:56
*/

package encrypt

import (
	"GoWeb/lesson29/bluebell/settings"
	"crypto/md5"
	"encoding/hex"
)

func EncryptPassword(opassword string) string {
	h := md5.New()
	h.Write([]byte(settings.Conf.EncryptConfig.SecretKey))
	return hex.EncodeToString(h.Sum([]byte(opassword)))
}
