package game

import "snake/game/component"

type MonitorService struct {
	Monitor component.MonitorFunType
}

func NewMonitorService() *MonitorService {
	return &MonitorService{Monitor: component.InitMonitor()}
}

func (s *MonitorService) start() component.MonitorFunType {
	return s.Monitor
}
