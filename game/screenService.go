package game

import "snake/game/component"

type ScreenService struct {
	Screen component.ScreenFunType
}

func NewScreenService() *ScreenService {
	return &ScreenService{Screen: component.InitScreen()}
}

func (s *ScreenService) start() component.ScreenFunType {
	return s.Screen
}
