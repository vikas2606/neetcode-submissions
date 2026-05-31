func isAnagram(s string, t string) bool {
	if (len(s)!=len(t)){
		return false
	}
	seen := make (map[rune]int,len(s))
	for _,i:=range s{
		seen[i]++
	}
	for _,j:=range t{
		seen[j]--
	}

	for _, v := range seen {
    if v != 0 {
        return false
    }
}
return true
}
