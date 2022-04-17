package weekCamp

func waysToBuyPensPencils(total int, cost1 int, cost2 int) int64 {
    dp1 := make([]int64, total+1)
    dp2 := make([]int64, total+1)
    
    for i:=0; i<=total; i++ {
        if i < cost1 {
            dp1[i] = 0
        } else {
            dp1[i] = dp1[i-cost1]+1
        }
        
        if i < cost2 {
            dp2[i] = 0
        } else {
            dp2[i] = dp2[i-cost2]+1
        }
    }
    
    var ret int64
    for i:=0; i<=int(dp1[total]); i++ {
        rest := total-i*cost1
        ret += dp2[rest]+1
    }
    
    
    return ret
}