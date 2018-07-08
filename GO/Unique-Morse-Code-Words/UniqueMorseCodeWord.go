package UniqueMorseCodeWord

// International Morse Code defines a standard encoding where each letter is mapped to a series of dots and dashes, as follows: "a" maps to ".-", "b" maps to "-...", "c" maps to "-.-.", and so on.

// For convenience, the full table for the 26 letters of the English alphabet is given below:

// [".-","-...","-.-.","-..",".","..-.","--.","....","..",".---","-.-",".-..","--","-.","---",".--.","--.-",".-.","...","-","..-","...-",".--","-..-","-.--","--.."]
// Now, given a list of words, each word can be written as a concatenation of the Morse code of each letter. For example, "cab" can be written as "-.-.-....-", (which is the concatenation "-.-." + "-..." + ".-"). We'll call such a concatenation, the transformation of a word.

// Return the number of different transformations among all words we have.

import (
    "bytes"
)

func uniqueMorseRepresentations(words []string) int {
    
	morseCode := [26]string{".-", "-...", "-.-.", "-..", ".", "..-.", "--.", 
		"....", "..", ".---", "-.-", ".-..", "--", "-.", "---", ".--.", "--.-",
		 ".-.", "...", "-", "..-", "...-", ".--", "-..-", "-.--", "--.."}
    var null struct{}
    
    m := make(map[rune]string)
    dict := make(map[string]interface{})
    cnt := 0
    
    for i, mc := range morseCode {
        m[toChar(i)] = mc
    }
    
    var tmp string
    for _, word := range words {
        tmp = getMorse(word, m)
        if _, ok := dict[tmp]; !ok {
            cnt ++
            dict[tmp] = null
        }
    }

    return cnt
}

func toChar(i int) rune {
    return rune('a' + i)
}

func getMorse(word string, m map[rune]string) string {
    var buffer bytes.Buffer
    for _, w := range word {
        buffer.WriteString(m[w])
    }
    return buffer.String()
}