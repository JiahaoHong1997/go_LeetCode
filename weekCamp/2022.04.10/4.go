package weekCamp

import "sort"

func maximumBeauty(flowers []int, newFlowers int64, target int, full int, partial int) int64 {
    sort.Ints(flowers)
    n := len(flowers)
   
    // count:已经完善的个数
    count := 0
    // 计算 >= target数目, 退出循环后 flowers[0, idx] < target
    idx := n - 1
    for idx >= 0 && flowers[idx] >= target {
        count++
        idx--
    }
    //fmt.Println("idx:", idx, "count:", count)
    
    // 全部都是完善的
    if idx == -1 {
        return int64(n * full)
    }
    
    sum := make([]int64, n + 1)
    for i := 1; i <= n; i++ {
        sum[i] = sum[i - 1] + int64(flowers[i - 1])
    }
    
    check := func(left, right, remain, t int) bool {
        // 1.找到值 < t的区间的右端点
        for left < right {
            mid := left + (right - left + 1) / 2
            if flowers[mid] < t {
                left = mid 
            } else {
                right = mid - 1
            }
        }
        // 2.检查区间[0, left] 是否可以完善为t
        return int64(remain) >= int64(t * (left + 1)) - sum[left + 1]
        
    }
    
    var res int64
    
    // cnt: 新的变成完善的个数, 从右往左填
    for cnt := 0; cnt <= idx + 1; cnt++ {
        // 区间：[idx - cnt + 1,idx]完善, [0, idx - cnt]不完善，并计算[0, idx - cnt]中的最小值的最大值
        // 填充完区间[idx - cnt + 1, idx]后剩余可种的花的数目
        remain := newFlowers - (int64(target * cnt) - (sum[idx + 1] - sum[idx - cnt + 1]))
       
        //fmt.Println("remain:", remain)
        if remain < 0 {
            break
        } 
        
        // 计算 [0, i - 1]中的最小值, 直接二分答案。
        left, right := flowers[0], target - 1
        
        for left < right {
            mid := left + (right - left + 1) / 2
            // 检查是否能够将区间[0, idx -cnt]中，值 < mid的区间（假设是区间[0, j]）全部完善为mid
            if check(0, idx - cnt, int(remain), mid) {
                left = mid
            } else {
                right = mid - 1
            }
            
        }
        
        //fmt.Println(i, count + idx - i + 1, left)
       
        // 如果全部完善， 则只有 完善数目 * full
        if count + cnt == n {
            res = max(res, int64((count + cnt) * full))
        } else {
            // 完善的数量：count + cnt, 未完善中最小值：left
            res = max(res, int64((count + cnt) * full + left * partial))
        }
        
    }
    
    return res   
   
}

func max(a, b int64) int64 {
    if a >= b {
        return a
    }
    return b
}