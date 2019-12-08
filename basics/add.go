package basics

func Add(ops ...int) int {
	accumulator := 0
	for _, op := range ops {
		accumulator += op
	}
	return accumulator
}
