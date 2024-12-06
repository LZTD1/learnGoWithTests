package iterate

func Repeat(s string, c int) string {
	var r string
	for i := 0; i < c; i++ {
		r += s
	}
	return r
}
