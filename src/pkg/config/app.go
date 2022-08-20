package config

func GetLocale() string {
	return "fa"
}
func isFa() bool {
	if GetLocale() == "fa" {
		return true
	}
	return false
}
