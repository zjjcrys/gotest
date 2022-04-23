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
//topic 227 边界条件
func calculate(s string) int {
	stack:=make([]int,0)
	res:=0
	num:=0
	var op byte
	op='+'
	for i:=0;i<len(s);i++ {
		if s[i]>='0' {
			num=num*10+int(s[i]-'0')
		}
		if s[i]!=' '&&s[i]<'0' || i==len(s)-1{
			if op=='+' {
				stack=append(stack,num)
			} else if op=='-' {
				stack=append(stack,0-num)
			} else if op=='*' {
				top:=stack[len(stack)-1]
				stack[len(stack)-1]=top*num
			} else if op=='/' {
				top:=stack[len(stack)-1]
				stack[len(stack)-1]=top/num
			}
			num=0
			op=s[i]
		}

	}
	for i:=0;i<len(stack);i++ {
		res+=stack[i]
	}
	return res
}