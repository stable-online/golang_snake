package component

func newScreenData() *screenData {
	return &screenData{snakes: new(snake), foodPoint: new(scope)}
}
