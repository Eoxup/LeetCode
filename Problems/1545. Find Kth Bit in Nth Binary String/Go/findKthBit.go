func findKthBit(n int, k int) byte {
    if k & 1 > 0 {
        if k >> 1 & 1 == 0 {
            return '0'
        } else {
            return '1'
        }
    }
    if k & ((k ^ (k - 1)) + 1) == 0 {
        return '1'
    } else {
        return '0'
    }
}