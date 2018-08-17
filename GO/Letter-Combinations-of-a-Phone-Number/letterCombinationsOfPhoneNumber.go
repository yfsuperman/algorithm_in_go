package letterCombinationsOfPhoneNumber

// Given a string containing digits from 2-9 inclusive, return all possible letter combinations that the number could represent.

// A mapping of digit to letters (just like on the telephone buttons) is given below. Note that 1 does not map to any letters.

// Example:

// Input: "23"
// Output: ["ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"].

func letterCombinations(digits string) []string {
    var res []string
    if len(digits) == 0 { return res }

	dict := map[string][]string{
        "1": []string{},
		"2": []string{"a", "b", "c"},
		"3": []string{"d", "e", "f"},
		"4": []string{"g", "h", "i"},
		"5": []string{"j", "k", "l"},
		"6": []string{"m", "n", "o"},
		"7": []string{"p", "q", "r", "s"},
		"8": []string{"t", "u", "v"},
		"9": []string{"w", "x", "y", "z"},
	}
    
    prev := letterCombinations(digits[1:])
    for _, i := range dict[string(digits[0])] {
        if len(prev) > 0 {
            for _, pre := range prev {
                res = append(res, string(i) + pre)
            }
        } else {
            res = append(res, string(i))
        }
    }
    
    return res
}