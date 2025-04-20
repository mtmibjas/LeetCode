// func moveZeroes(nums []int) {
// 	arr := make([]int, len(nums))
// 	index := 0
// 	for _, n := range nums {
// 		if n > 0 {
// 			arr[index] = n
// 			index++
// 		}
// 	}
//     for i, n := range arr {
//         nums[i] = n
//     }
// }

func moveZeroes(nums []int) {
    index := 0
    for _, n := range nums {
        if n != 0 {
            nums[index] = n
            index++
        }
    }
    for index < len(nums) {
        nums[index] = 0
        index++
    }
}