package smallest_string_with_given_numeric_value_1663

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}
func getSmallestString(n int, k int) string {
    out := make([]byte, n)
    for i := 0; i < n; i++{
        out[i] = byte('a')
    }
    k = k - n
    pos := n - 1
    for k > 0 {
        tmp := min(k, 25)
        help := int(out[pos])        
        out[pos] = byte(help + tmp)
        pos--
        k = k - tmp
    }
    return string(out)
}