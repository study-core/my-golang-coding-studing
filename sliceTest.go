package main

func main() {
	s := make([]myslice, 0)
	t := &myInt{a: 4}
	s = append(s, t)

	for i, v := range s {
		if i == 3 {

		}
	}
}

type myInt struct {
	a  int
}
type myslice []*myInt

func (cs myslice) deactivate(index int) caseList {
	last := len(cs) - 1
	cs[index], cs[last] = cs[last], cs[index]
	return cs[:last]
}