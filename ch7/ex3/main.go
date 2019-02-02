package main

import "fmt"

func main() {
	t := &tree{
		value: 0,
		left: &tree{
			value: 1,
			left: &tree{
				value: 3,
			},
		},
		right: &tree{
			value: 2,
			right: &tree{
				value: 4,
			},
		},
	}

	fmt.Println(t)
}

type tree struct {
	value       int
	left, right *tree
}

func (t *tree) String() string {
	return fmt.Sprintln(travel(t, []int{}))
}

func travel(t *tree, vals []int) []int {
	if t == nil {
		return vals
	}
	vals = append(vals, t.value)
	vals = travel(t.left, vals)
	vals = travel(t.right, vals)
	return vals
}
