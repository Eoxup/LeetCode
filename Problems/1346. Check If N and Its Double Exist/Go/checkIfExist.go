func checkIfExist(arr []int) bool {
    dbMap := make([]bool, 4001)
    for i := 0; i < len(arr); i++ {
        if dbMap[arr[i]*2 + 2000] {
            return true
        } else if arr[i] & 1 == 0 && dbMap[arr[i] / 2 + 2000] {
            return true
        } else {
            dbMap[arr[i] + 2000] = true
        }
    }
    return false
}
