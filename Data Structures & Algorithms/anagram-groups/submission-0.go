import "slices"

func sortString(s string) string {
	runes := []rune(s)

	sort.Slice(runes, func(i, j int) bool {
		return runes[i] < runes[j]
	})

	return string(runes)
}


func groupAnagrams(strs []string) [][]string {
	op:=make(map[string][]string)
	var result [][]string
	for _,i:=range strs{
		op[sortString(i)]=append(op[sortString(i)],i)
	}

	for _, v := range op {
		result = append(result, v)
	}

	slices.SortFunc(result, func(a, b []string) int {
		return len(a) - len(b)
	})

return result

}
