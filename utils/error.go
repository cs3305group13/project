package utils

// panics err if err != nil
func CheckError( err error ) {
	if err != nil {
		panic(err)
	}
}