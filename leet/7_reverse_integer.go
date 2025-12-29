package leet

func reverse(x int) int {
	const (
		Max = 1<<31 - 1
		Min = -1 << 31
	)

	rev := 0

	for x != 0 {
		d := x % 10
		x /= 10

		if rev > Max/10 || (rev == Max/10 && d > 7) {
			return 0
		}

		if rev < Min/10 || (rev == Min/10 && d < -8) {
			return 0
		}

		rev = rev*10 + d
	}

	return rev
}
