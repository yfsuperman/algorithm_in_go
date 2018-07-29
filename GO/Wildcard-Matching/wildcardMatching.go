package wildcardMatching

// Given an input string (s) and a pattern (p), implement wildcard pattern matching with support for '?' and '*'.

// '?' Matches any single character.
// '*' Matches any sequence of characters (including the empty sequence).
// The matching should cover the entire input string (not partial).

// Note:

// s could be empty and contains only lowercase letters a-z.
// p could be empty and contains only lowercase letters a-z, and characters like ? or *.

func isMatch(s string, p string) bool {
    si, pi, star, match := 0, 0, -1, 0
    for si < len(s) {
        if pi < len(p) && (s[si] == p[pi] || p[pi] == '?') {
            si, pi = si + 1, pi + 1
        } else if pi < len(p) && p[pi] == '*' {
            star, pi, match = pi, pi + 1, si
        } else if (star != -1) {
            pi, match, si = star + 1, match + 1, match +1
        } else {
            return false
        }
    }
    for pi < len(p) && p[pi] == '*' { pi ++ }
    
    return pi == len(p)
    
}