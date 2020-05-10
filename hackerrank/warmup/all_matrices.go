package main

//Given a 3x3 matrix rotate it and print all the combinations

func permute(arr [][]int32) {
	all := make([]int32, 0, 9)
	for i, a := range arr {
		for j, el := range a {
			all = append(all, el)
		}
	}

}
