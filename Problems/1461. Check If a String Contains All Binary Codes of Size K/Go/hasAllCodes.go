func hasAllCodes(s string, k int) bool {
    if len(s) < (1 << k) + (k - 1) {
        return false
    }

    allCodes := make([]bool, 1 << k)
    n := 0
    for i := 0; i < k; i++ {
        n <<= 1
        if s[i] == '1' {
            n |= 1
        }
    }
    allCodes[n] = true

    c := len(allCodes) - 1
    m := (1 << k) - 1
    for i := k; i < len(s); i++ {
        n <<= 1
        if s[i] == '1' {
            n |= 1
        }
        n &= m
        if allCodes[n] == false {
            allCodes[n] = true
            c--
        }
        if c == 0 {
            return true
        } else if c > len(s) - i {
            return false
        }
    }

    if c == 0 {
        return true
    } else {
        return false
    }
}
