package leedcode
//pid 42 单调栈 高度
func trap(height []int) int {
	res:=0
	if len(height)<3 {
		return res
	}
	stack:=make([]int,0)
	stack=append(stack,0)
	i:=1
	for i<len(height) {
		if len(stack)<1 || height[i]<height[stack[len(stack)-1]] {
			stack=append(stack,i)
			i++
			continue
		}
		top:=height[stack[len(stack)-1]]
		stack=stack[:len(stack)-1]
		if len(stack)<1 {
			continue
		}
		left:=stack[len(stack)-1]
		res+=(min(height[i],height[left])-top)*(i-left-1)
	}
	return res
}

func trap2(height []int) int {
	res:=0
	if len(height)<3 {
		return res
	}
	stack:=make([]int,0)
	index:=0
	for index<len(height) {
		if len(stack)<1 || height[index]<height[stack[len(stack)-1]] {
			stack=append(stack,index)
			index++
			continue
		}
		top:=height[stack[len(stack)-1]]
		stack=stack[:len(stack)-1]
		if len(stack)<1 {
			continue
		}
		left:=stack[len(stack)-1]
		res+=(min(height[index],height[left])-top)*(index-left-1)
	}

	return res
}
