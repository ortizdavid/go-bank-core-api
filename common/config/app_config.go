package config

import "github.com/ortizdavid/go-nopain/conversion"


func ListenAddr() string {
	return GetEnv("APP_HOST") +":"+ GetEnv("APP_PORT")
}

func ShutdownTimeout() int {
	return conversion.StringToInt(GetEnv("SHUTDOWN_TIMEOUT"))
}

//------- Uploads
func UploadImagePath() string {
	return GetEnv("UPLOAD_IMAGE_PATH")
}

func UploadDocumentPath() string {
	return GetEnv("UPLOAD_DOCUMENT_PATH")
}

// ---- Requests
func RequestsPerSecond() int {
	return conversion.StringToInt(GetEnv("REQUESTS_PER_SECONDS"))
}

func RequestsExpiration() int {
	return conversion.StringToInt(GetEnv("REQUESTS_EXPIRATION"))
}

