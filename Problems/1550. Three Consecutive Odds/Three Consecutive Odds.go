func threeConsecutiveOdds(arr []int) bool {
    x := 0
    for _, n := range arr {
        if n & 1 == 1  {
            x++
        } else {
            x = 0
        }
        if x == 3 {
            return true
        }
    }
    return false
}
