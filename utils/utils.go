package utils

// InList 判断该图片类型是否存在白名单中
func InList(key string, list []string) bool {
	for _, s := range list {
		if key == s {
			return true
		}
	}
	return false
}
