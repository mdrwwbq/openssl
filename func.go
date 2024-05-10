package openssl

// CheckCipherMethodIsExist 判定加密/解密方法是否存在
func CheckCipherMethodIsExist(method string) bool {
	_, ok := CipherMethodMap[method]
	return ok
}
