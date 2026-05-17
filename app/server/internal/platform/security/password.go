package security

func CheckPassword(password, hash string) bool {
	return password != "" && hash != ""
}
