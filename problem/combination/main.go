package main

import "fmt"

func combinationSum(candidates []int, target int) [][]int {
	var result [][]int
	var backtrack func(remain int, comb []int, start int)

	backtrack = func(remain int, comb []int, start int) {
		if remain == 0 {
			var temp = make([]int, len(comb))
			// var temp []int
			copy(temp, comb)
			result = append(result, temp)
			return
		} else if remain < 0 {
			return
		}

		for i := start; i < len(candidates); i++ {
			comb = append(comb, candidates[i])
			backtrack(remain-candidates[i], comb, i) // Pass 'i' instead of 'start'
			comb = comb[:len(comb)-1]
		}
	}

	backtrack(target, []int{}, 0)
	return result
}

// func combinationSum(candidates []int, target int) [][]int {
// 	var output [][]int
// 	var mapDta = make(map[string]int)
// 	for i := 0; i < len(candidates); i++ {
// 		mapDta[fmt.Sprintf("%d", candidates[i])] = 0
// 	}

// 	for i := 0; i < len(candidates); i++ {
// 		reminder := target % candidates[i]
// 		fmt.Println("reminder", reminder, candidates[i])
// 		if inc, exist := mapDta[fmt.Sprintf("%d", candidates[i])]; reminder == 0 && exist && inc == 0 {
// 			if candidates[i] != target {
// 				var out []int
// 				for j := 0; j < target/candidates[i]; j++ {
// 					out = append(out, candidates[i])
// 				}
// 				output = append(output, out)
// 			} else {
// 				output = append(output, []int{candidates[i]})
// 			}
// 			mapDta[fmt.Sprintf("%d", candidates[i])]++

// 			continue
// 		}
// 		fmt.Println("mapDta", mapDta[fmt.Sprintf("%d", target-(candidates[i]+reminder))], target, (candidates[i] + reminder))
// 		if _, exist := mapDta[fmt.Sprintf("%d", target-(candidates[i]+reminder))]; exist {
// 			fmt.Println("d", target-(candidates[i]+reminder), (candidates[i]+reminder)+candidates[i])

// 			a, b := (candidates[i] + reminder), candidates[i]
// 			c := target - a
// 			if a > b {
// 				a, b = b, a
// 			}
// 			if a > c {
// 				a, c = c, a
// 			}
// 			if b > c {
// 				b, c = c, b
// 			}
// 			fmt.Println("debug", mapDta[fmt.Sprintf("%d_%d_%d", a, b, c)], a, b, c)
// 			switch {
// 			case mapDta[fmt.Sprintf("%d", reminder)] == 0 && mapDta[fmt.Sprintf("%d", c)] == 0 && mapDta[fmt.Sprintf("%d_%d_%d", a, b, c)] == 0:
// 				key := fmt.Sprintf("%d_%d_%d", a, b, c)
// 				mapDta[key] = int(1)
// 				output = append(output, []int{a, b, c})

// 			case (candidates[i]+reminder)+candidates[i] == target && mapDta[fmt.Sprintf("%d_%d", a, b)] == 0:
// 				key := fmt.Sprintf("%d_%d", a, b)
// 				if d, exist := mapDta[key]; exist {
// 					fmt.Println(exist, d)
// 					continue
// 					// fallthrough
// 				}
// 				mapDta[key] = int(1)
// 				output = append(output, []int{a, b})
// 			default:
// 				fmt.Println("target", target-(candidates[i]+reminder))

// 				rem := target - (candidates[i] + reminder + candidates[i])
// 				fmt.Println(rem, mapDta[fmt.Sprintf("%d_%d", a, b)])
// 				if rem <= 0 {
// 					continue
// 				}

// 				c := rem
// 				if a > b {
// 					a, b = b, a
// 				}
// 				if a > c {
// 					a, c = c, a
// 				}
// 				if b > c {
// 					b, c = c, b
// 				}
// 				key := fmt.Sprintf("%d_%d_%d", a, b, c)
// 				if _, exist := mapDta[key]; exist {
// 					continue
// 				}
// 				mapDta[key] = 0
// 				output = append(output, []int{a, b, c})
// 			}
// 		} else {

// 		}
// 	}
// 	return output
// }

func main() {
	in := []int{2, 3, 5}
	target := 8
	want := [][]int{{2, 2, 2, 2}, {2, 3, 3}, {3, 5}}
	result := combinationSum(in, target)
	fmt.Println(result, want)
}
