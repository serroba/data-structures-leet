package training

func NumOfUnplacedFruits(fruits []int, baskets []int) int {
	placed := 0
	for i := range fruits {
		for j := 0; j < len(baskets); j++ {
			if fruits[i] <= baskets[j] {
				baskets[j] = -1
				placed++
				break
			} else {
				continue
			}
		}
	}
	return len(baskets) - placed
}
