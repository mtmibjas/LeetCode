func canPlaceFlowers(flowerbed []int, n int) bool {
	if n == 0 {
		return true
	} else if len(flowerbed) < 3 {
		if slices.Contains(flowerbed, 1) {
			return false
		} else {
            if n > 1 {
                return false
            }
			return true
		}

	} else {
		count := 0

		for i, j, k := 0, 1, 2; k < len(flowerbed); {
			fe, se, te := flowerbed[i], flowerbed[j], flowerbed[k]
			if i == 0 && fe == 0 && se == 0 {
                flowerbed[i] = 1
				count++
				i++
				j++
				k++
			//} else if k == len(flowerbed)-1 && se == 0 && te == 0 {
			//	count++
			//	i++
			//	j++
			//	k++
			} else if fe == 0 && se == 0 && te == 0 {
                flowerbed[j] = 1
				count++
				i += 2
				j += 2
				k += 2

			} else {
				i++
				j++
				k++
			}

			if count == n {
				return true
			}
           // fmt.Println(fe, se, te, count)
		}
        if flowerbed[len(flowerbed)-1] == 0 && flowerbed[len(flowerbed)-2] == 0 {
            count++
            if count == n {
				return true
			}
        }
	}

	return false
}