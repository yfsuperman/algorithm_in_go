package regularExpressionMatching

// Given an input string (s) and a pattern (p), implement regular expression matching with support for '.' and '*'.

// '.' Matches any single character.
// '*' Matches zero or more of the preceding element.
// The matching should cover the entire input string (not partial).

// Note:

// s could be empty and contains only lowercase letters a-z.
// p could be empty and contains only lowercase letters a-z, and characters like . or *.

func isMatch(s string, p string) bool {
    lens, lenp := len(s), len(p)
    
    dp := make([][]bool, lens + 1)
    for i := range dp {
        dp[i] = make([]bool, lenp + 1)
    }
    
    dp[0][0] = true
    for i := 1; i < lenp; i++ {
        if p[i] == '*' && dp[0][i-1] { dp[0][i+1] = true }
    } 
    
    for i := 0; i < lens; i++ {
        for j := 0; j < lenp; j++ {
            if s[i] == p[j] || p[j] == '.' {
                dp[i+1][j+1] = dp[i][j]
            }
            if p[j] == '*' {
                if p[j-1] != s[i] && p[j-1] != '.' {
                    dp[i+1][j+1] = dp[i+1][j-1]
                } else {
                    dp[i+1][j+1] = dp[i+1][j] || dp[i+1][j-1] || dp[i][j+1]
                }
            }
        }
    }
    
    return dp[lens][lenp]
    
}