package password

import (
	"fmt"
	"time"
)

/*
Задача найти комбинаций пороля из 10 чисел 4-значное
*/
func main() {
	binaryFindPass()
	findForLoops()
}

func binaryFindPass() {
	now := time.Now()
	var password [][2]int
	var password2 [][2]int

	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			password = append(password, [2]int{i, j})
		}
	}

	for a := 0; a < 10; a++ {
		for b := 0; b < 10; b++ {
			password2 = append(password2, [2]int{a, b})
		}
	}

	var merge [][4]int

	for i := 0; i < len(password); i++ {
		for j := 0; j < len(password2); j++ {
			merge = append(merge, [4]int{
				password[i][0], password[i][1],
				password2[j][0], password2[j][1],
			})
		}
	}
	fmt.Println("variant-1", time.Since(now))

}

func findForLoops() {
	now := time.Now()

	var allCombination [][4]int
	for a := 0; a < 10; a++ {
		for b := 0; b < 10; b++ {
			for c := 0; c < 10; c++ {
				for d := 0; d < 10; d++ {
					allCombination = append(allCombination, [4]int{a, b, c, d})
				}
			}
		}
	}

	fmt.Println("variant-2", time.Since(now))
}
