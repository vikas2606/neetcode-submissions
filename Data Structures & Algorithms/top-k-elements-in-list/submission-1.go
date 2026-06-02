func topKFrequent(nums []int, k int) []int {
 freq := make(map[int]int)

for _, num := range nums {
	freq[num]++
}

type Pair struct {
	Key   int
	Value int
}

pairs := make([]Pair, 0, len(freq))

for k, v := range freq {
	pairs = append(pairs, Pair{k, v})
}

sort.Slice(pairs, func(i, j int) bool {
	return pairs[i].Value > pairs[j].Value
})

res := make([]int, 0, k)
for i := 0; i < k; i++ {
	res = append(res, pairs[i].Key)
}

return res

}
