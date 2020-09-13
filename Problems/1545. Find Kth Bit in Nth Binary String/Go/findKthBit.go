package findKthBit

func FindKthBit(n int, k int) byte {
    /**********************************
     binary k | ans | y | z |
    ----------|-----|---|---|
     xxxxxx01 | '0' | 0 | 0 |
     xxxxxx11 | '1' | 1 | 0 |
     x010...0 | '1' | 0 | 1 |
     x110...0 | '0' | 1 | 1 |

    #define  x = the first bit 1 of k from right side
    #define  y = the bit left then x
    #define  z = {
        1, if k is even
        0, if k is odd
    }
    ans = y ^ z
    **********************************/
    if (k & -k) << 1 & k > 0 {
        return '0' + byte((k & 1) ^ 0)
    } else {
        return '0' + byte((k & 1) ^ 1)
    }
}
