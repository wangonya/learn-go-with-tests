package iteration

func Repeat(s string, t int) string {
	repeated_s := ""
	for i := 1; i <= t; i++ {
		repeated_s += s
	}
	return repeated_s
}
