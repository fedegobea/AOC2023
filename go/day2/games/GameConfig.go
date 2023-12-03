package games

type GameConfig struct {
	Red   int
	Green int
	Blue  int
}

func (g GameConfig) CheckValidRed(in int) bool {
	return g.Red >= in
}
func (g GameConfig) CheckValidGreen(in int) bool {
	return g.Green >= in
}
func (g GameConfig) CheckValidBlue(in int) bool {
	return g.Blue >= in
}

func (g GameConfig) IsPossibleRound(round Round) bool {
	return g.CheckValidRed(round.Red) && g.CheckValidGreen(round.Green) && g.CheckValidBlue(round.Blue)
}
