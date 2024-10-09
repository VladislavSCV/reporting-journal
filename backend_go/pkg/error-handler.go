package pkg

// CError checks if the error is not nil and logs it. This is a convenience
// function for cases where you want to log an error but don't want to
// explicitly check if the error is not nil.
func CError(err error) {
	if err != nil {
		Log(err)
	}
}
