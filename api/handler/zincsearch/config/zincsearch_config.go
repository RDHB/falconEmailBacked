package config

// GetConfig function that we can export config needed for consume zincsearch api
func GetConfig() map[string]string {
	zincsearchConfig := map[string]string{}
	// zincsearchConfig["user_id"] = os.Getenv("ZINCSEARCH_USER_ID")
	// zincsearchConfig["password"] = os.Getenv("ZINCSEARCH_PASSWORD")
	// zincsearchConfig["zincHost"] = os.Getenv("ZINCSEARCH_ZINCHOST")
	zincsearchConfig["zincURL"] = "%s/api/%s/_search"
	return zincsearchConfig
}
