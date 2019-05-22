package main

import "fmt"

func main() {
	mapa := make(candidateStorage)
	mapb := make(candidateStorage)

	ids := make([]string, 0)


	mapb["b"] = "B"
	ids = append(ids, "b")
	fmt.Println(mapa, mapb)
	copyCandidateMapByIds(mapa, mapb, ids)
	fmt.Println(mapa, mapb)

}


type candidateStorage map[string]string

func copyCandidateMapByIds(target, source candidateStorage, ids []string) {
	for _, id := range ids {
		if v, ok := source[id]; ok {
			target[id] = v
		}
	}
}