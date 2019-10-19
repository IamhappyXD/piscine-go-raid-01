package piscine



func Raid1d(x, y int) {

	if x > 0 && y > 0 {

		if x == 1 && y == 1 {
			z01.PrintRune('A')
			z01.PrintRune('\n')
			return
		} else if x == 1 {
			z01.PrintRune('A')
			for t := 1; t < y-1; t++ { //stolko raz skolko y - 2
				if t <= 0 {
					break
				} else {
					z01.PrintRune('\n')
					z01.PrintRune('B')

				}

			}
			z01.PrintRune('\n')
			z01.PrintRune('A')
			z01.PrintRune('\n')

		} else if y == 1 {
			z01.PrintRune('A')
			for i := 1; i < x-1; i++ {
				z01.PrintRune('B') // 3 dash iz 5
			}
			z01.PrintRune('C') // is the top-right corner if y=1
			z01.PrintRune('\n')
			return
		} else {
			z01.PrintRune('A')
			for i := 1; i < x-1; i++ {
				z01.PrintRune('B') // 3 dash iz 5
			}
			z01.PrintRune('C') // changing position on right top corner
			z01.PrintRune('\n')

			for t := 1; t < y-1; t++ { //stolko raz skolko y - 2
				z01.PrintRune('B')

				for j := 1; j < x-1; j++ {
					z01.PrintRune(' ') // 3 probela iz 5
				}
				z01.PrintRune('B')
				z01.PrintRune('\n')

			}

			z01.PrintRune('A') // is the low-left

			for i := 1; i < x-1; i++ {
				z01.PrintRune('B') // 3 dash iz 5
			}
			z01.PrintRune('C')
			z01.PrintRune('\n')

		}
	}
}
