package ci

func minArray(numbers []int) int {
    l, r := 0, len(numbers)-1
    for l < r {
        if numbers[l] == numbers[r] {
            r--
        }
        mid := l + (r-l)/2
        if numbers[mid] == numbers[r] {
            r--
        } else if numbers[mid] < numbers[r] {
            r = mid
        } else {
            l = mid+1
        }
    }
    return numbers[l]
}