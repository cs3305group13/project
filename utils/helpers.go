package utils

// Function removes string value from slice at position index s
func Remove(slice []string, s int) []string {
    return append(slice[:s], slice[s+1:]...)
}

// Function checks if the string e is inside array
func ArrayContains(array []string, e string) bool {
    for _, s := range array {
        if e == s {
            return true
        }
    }
    return false
}
