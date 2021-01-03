package Menu

import (
	"Game/Window"
	"Game/Components/States"
	"github.com/faiface/pixel/pixelgl"
)

func ListenForActions(winConf Window.WindowConfig, currState *States.States){
	if (winConf.Win.MousePosition().X >= 379 && winConf.Win.MousePosition().X <= 590) && (winConf.Win.MousePosition().Y >= 320 && winConf.Win.MousePosition().Y <= 415) && winConf.Win.Pressed(pixelgl.MouseButtonLeft){
		currState.SetCreateLobbyMenu()
	}
	if (winConf.Win.MousePosition().X >= 379 && winConf.Win.MousePosition().X <= 590) && (winConf.Win.MousePosition().Y >= 144 && winConf.Win.MousePosition().Y <= 242) && winConf.Win.Pressed(pixelgl.MouseButtonLeft){
		currState.SetJoinLobbyMenu()
	}
}