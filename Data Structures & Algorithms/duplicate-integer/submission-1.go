func hasDuplicate(nums []int) bool {
    x:=make(map[int]bool)
    for _,i := range nums{
		if _,exists:=x[i];exists{
			return true
		}
		x[i]=true
    }
    return false
}
