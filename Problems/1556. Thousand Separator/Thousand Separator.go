func thousandSeparator(n int) string {
    if n == 0 {
        return "0"
    }
    s := make([]byte, 13)
    s[9] = '.'
    s[5] = '.'
    s[1] = '.'
    i := 13
    for n > 0 {
        i--
        if (s[i] == '.') {
            i--
        }
        s[i] = byte(n % 10) + 0x30
        n /= 10
    }
    return string(s[i:])
}
