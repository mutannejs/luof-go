package lerror

func Equals(a, b error) bool {
    if (a == nil || b == nil) {
        return false
    }
    return a.Error() == b.Error()
}
