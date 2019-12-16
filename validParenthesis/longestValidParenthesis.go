package leetcode

func isValid(sub string) bool {
	if len(sub) == 0 {
		return false
	}
	stack := 0
	for i := 0; i < len(sub); i++ {
		if stack == 0 && sub[i] == byte(')') {
			return false
		}
		if sub[i] == byte('(') {
			stack++
		} else {
			stack--
		}

	}
	return stack == 0
}
func longestValidParentheses(s string) int {
	max := 0
	for i := 0; i < len(s); i++ {
		for j := len(s); j > i; j-- {
			//fmt.Printf("evaluationg %s on range i=%d & j=%d and isValid=%v \n",s[i:j],i,j, isValid(s[i:j]))

			if s[i] == byte('(') && s[j-1] == byte(')') && (j-i)%2 == 0 && (j-i) > max {
				if isValid(s[i:j]) {
					max = j - i
				}
			}
		}
	}
	return max
}
