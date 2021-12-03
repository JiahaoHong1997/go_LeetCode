package leetcode139

func wordBreak(s string, wordDict []string) bool {
    m := make(map[string]bool,0)
    for _,v := range wordDict {
        m[v] = true
    }

    dp := make([]bool,len(s)+1)
    dp[0] = true

    for i:=1; i<=len(s); i++ {
        for j:=0; j<i; j++ {
            if dp[j] && m[s[j:i]] {
                
                dp[i] = true
                break
            }
        }
    }

    return dp[len(s)]
}