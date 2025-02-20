package slidingtechnique

import (
	"strings"
)

// Problem:  Given a string s, find the length of the longest substring without repeating characters.

func LongestSubstring(s string) string {
  characters := strings.Split(s, "")
  alreadySeem := make(map[string]bool)

  answer := ""
  tempAnswer := ""
  left := 0
  right := 0

  for right < len(characters){
    _, ok := alreadySeem[characters[right]]
    // there exist the element
    if ok {
      if len(tempAnswer) > len(answer) {
        answer = tempAnswer
      }
      left++
      tempAnswer = ""
      alreadySeem = make(map[string]bool)
      right = left
    }else{
      alreadySeem[characters[right]] = true
      tempAnswer += string(characters[right])
    }
    right++
  }
  return answer
}
