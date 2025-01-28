package setting

var serverUrl = "https://api.circle.com/v1/w3s"

// GetServerURL returns the server URL
func GetServerURL() string {
	return serverUrl
}

// SetServerURL sets the server URL
func SetServerURL(input string) string {
	serverUrl = input
	return serverUrl
}
