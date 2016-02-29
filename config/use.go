package config

func UseLocal() error {
	return Save(&Config{
		Type: "local",
	})
}

func UseMachine(name string) error {
	return Save(&Config{
		Type:    "machine",
		Machine: name,
	})
}

func UseUrl(url string) error {
	return Save(&Config{
		Type: "url",
		Url:  url,
	})
}

func UseUrlWithTls(url string, certsPath string) error {
	return Save(&Config{
		Type:     "url",
		Url:      url,
		CertPath: certsPath,
	})
}
