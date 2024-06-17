package util

func Check(err error) {
	if err != nil {
		panic(err)
	}
}

func EatErr(i any, err error) (a any) {
	Check(err)
	return i
}