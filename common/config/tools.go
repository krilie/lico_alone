package config

func GetInt(key string) int {
	ok := v.IsSet(key)
	if !ok {
		return 0
	}
	return v.GetInt(key)
}
func GetString(key string) string {
	ok := v.IsSet(key)
	if !ok {
		return ""
	}
	return v.GetString(key)
}
