package request

func Zen(token string) (string, error) {
	resp, err := GetRequester("https://api.github.com/zen", token)
	if err != nil {
		return "", err
	}
	return string(resp), nil
}
