package config

// GetConfig function that we can export config needed for consume zincsearch api
func GetConfig() map[string]string {
	zincsearchConfig := map[string]string{}
	zincsearchConfig["zincURL"] = "%s/api/%s/_search"
	return zincsearchConfig
}
