package component

import (
	"github.com/nsf/termbox-go"
)

// Game 游戏服务接口
//
// @Description:
type Game interface {
	Start()
}

//gameService 定义游戏实体
type gameService struct {
	//
	//  screenApp
	//  @Description: 屏幕
	//
	screenApp Screen

	//
	//  monitorApp
	//  @Description: 控制
	//
	monitorApp Monitor
}

//NewGameService 实例化游戏服务
func NewGameService() Game {
	return &gameService{screenApp: newScreenApp(), monitorApp: newMonitorApp()}
}

//Start 开始游戏
func (g *gameService) Start() {
	//初始化插件
	if initErr := termbox.Init(); initErr != nil {
		panic(initErr)
	}

	//函数退出时, 关闭包 rules
	defer termbox.Close()

	//游戏启动
	g.run(g.monitorApp.start())
}

//run 游戏启动
func (g *gameService) run(m *monitorApp) {
	//退出户标识
CloseGame:
	for {
		select {

		//键盘移动事件
		case operator := <-m.move:
			g.screenApp.setDirection(operator)

		//点击ECS事件
		case <-m.quit:
			//退出刷新
			break CloseGame

		//游戏角色状态
		case status := <-g.screenApp.getSnakeStatusChan():
			//设置人物状态
			g.screenApp.setGameOver(status)

		default:
			//如果蛇还没死
			if g.screenApp.getActivity() {
				//进行渲染界面
				if err := g.screenApp.start(); err != nil {
					panic(err.Error())
				}
			}

			g.screenApp.flush()
		}
	}
}
