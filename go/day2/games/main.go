package games

func panicOnErr(err error) {
	if err != nil {
		panic(err)
	}
}
