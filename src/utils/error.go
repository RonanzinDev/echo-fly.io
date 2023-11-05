package utils

func ErrorResponse(err string) map[string]string {
	return map[string]string{"error": err}
}
