func maximumStrongPairXor(nums []int) int {
	max := 0
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if int(math.Abs(float64(nums[i]-nums[j]))) <= min(nums[i], nums[j]) {
				c := xor(nums[i], nums[j])
                fmt.Println(nums[i], nums[j],c)
				if max < c {
					max = c
				}
			}
		}
	}

	return max
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}

func xor(i, j int) int {
	xi := convertToBinary(i)
	xj := convertToBinary(j)
	left := len(xi) - 1
	right := len(xj) - 1
	ln := left
	if ln < right {
		ln = right
	}
	arr := make([]int, ln+1)

	for left >= 0 || right >= 0 {
		x := 0
		if left >= 0 {
			x = xi[left]
		}
		y := 0
		if right >= 0 {
			y = xj[right]
		}
		arr[ln] = 0
		if x != y {
			arr[ln] = 1
		}
		ln--
		left--
		right--

	}
    fmt.Println(arr)
	return convertToInt(arr)
}


func convertToBinary(i int) []int {
	if i == 0 {
		return []int{0}
	}

	arr := make([]int, 0)
	d := i
	for d > 0 {
		r := d % 2 
		arr = append(arr, r)
		d = d / 2 
	}

	
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}

	return arr
}
func convertToInt(arr []int) int {
	sum := 0
	for i, bit := range arr {
		sum += bit * int(math.Pow(2, float64(len(arr)-1-i)))
	}


	return sum
}