func miniMaxSum(arr []int32) {
	// Write your code here
	sort.SliceStable(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})
	var sum int64
	for i := 0; i < len(arr); i++ {
		sum = sum + int64(arr[i])
	}
	var min int64
	var max int64
	for i := 0; i < len(arr); i++ {
		if sum-int64(arr[i]) <= min || min == 0 {
			min = sum - int64(arr[i])
		}
		if sum-int64(arr[i]) > max || max == 0 {
			max = sum - int64(arr[i])
		}
	}
	fmt.Println(min, max)

}