package leetcode

func twoSum(input []int, target int) []int {
	if len(input) <= 1 {
		return nil
	}

	m := make(map[int]int, 0)
	for i := 0; i < len(input); i++ {
		another := target - input[i]
		if val, ok := m[another]; ok {
			return []int{i, val}
		}
		m[input[i]] = i
	}

	return nil
}
