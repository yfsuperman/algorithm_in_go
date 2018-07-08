package JewelsAndStones

// You're given strings J representing the types of stones that are jewels, and S representing the stones you have.  Each character in S is a type of stone you have.  You want to know how many of the stones you have are also jewels.
// The letters in J are guaranteed distinct, and all characters in J and S are letters. Letters are case sensitive, so "a" is considered a different type of stone from "A".

// Example 1:
// Input: J = "aA", S = "aAAbbbb"
// Output: 3

func numJewelsInStones(J string, S string) int {
    
    var null struct{}
	m := make(map[rune]interface{})
	cnt := 0
    
    for _, j := range J {
        m[j] = null
    }
    
    for _, s := range S {
        if _, ok := m[s]; ok {
            cnt ++
        }
    }
    
    return cnt
}