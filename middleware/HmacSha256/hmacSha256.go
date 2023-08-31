package HmacSha256

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
)

var HmacKey = "XingTong_Official"

// HmacSha256 计算HmacSha256
func HmacSha256(data string) []byte {
	mac := hmac.New(sha256.New, []byte(HmacKey))
	mac.Write([]byte(data))

	return mac.Sum(nil)
}

// HmacSha256ToHex 将加密后的二进制转16进制字符串
func HmacSha256ToHex(message string) string {
	return hex.EncodeToString(HmacSha256(message))
}

// HmacSha256ToBase64 将加密后的二进制转Base64字符串
func HmacSha256ToBase64(message string) string {
	return base64.URLEncoding.EncodeToString(HmacSha256(message))
}
