package errorhandlers

func IfErrorStartPanic(err error) {
	if err != nil {
		panic(err)
	}
}
