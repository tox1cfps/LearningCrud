package produtos

func EncontrarIndicePorID(id int) int {
	for i, p := range produtos {
		if p.ID == id {
			return i
		}
	}
	return -1
}
