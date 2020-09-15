package p392

func isSubsequence(s string, t string) bool {
	sl := len(s)
	tl := len(t)
	i := 0
	j := 0
	for ;i<sl && j < tl; {
		if s[i] == t[j] {
			i++
			j++
		}else{
			j++
		}
	}
	return i == sl
}