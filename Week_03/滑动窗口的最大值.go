func maxSlidingWindow(nums []int, k int) []int {
    if nums == nil || k == 0{
        return nil
    }
    queue := []int{}
    res := []int{}
    for i := 0; i < k; i++ {
    	for len(queue) > 0 && queue[len(queue)-1] < nums[i] {
			queue = queue[0:len(queue)-1]
    	}
    	queue = append(queue, nums[i])
    }
    res = append(res, queue[0])
    for i := k; i < len(nums); i++ {
    	for len(queue) > 0 && queue[len(queue)-1] < nums[i] {
			queue = queue[0:len(queue)-1]
    	}
    	queue = append(queue, nums[i])
    	if nums[i-k] == queue[0] {
    		queue = queue[1:]
    	}
    	res = append(res, queue[0])
    }
    return res
}