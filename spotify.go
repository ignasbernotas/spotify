package spotify

func Login(username string, password string, deviceName string) (*Session, error) {
	return CoreLogin(username, password, deviceName)
}

func LoginSaved(username string, authData []byte, deviceName string) (*Session, error) {
	return CoreLoginSaved(username, authData, deviceName)
}

func LoginOAuth(deviceName string, clientId string, clientSecret string) (*Session, error) {
	return CoreLoginOAuth(deviceName, clientId, clientSecret)
}
