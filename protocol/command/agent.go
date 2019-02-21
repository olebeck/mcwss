package command

import (
	"fmt"
	"github.com/sandertv/mcwss/mctype"
)

// AgentPositionRequest produces the command used to get the position of a player's agent.
func AgentPositionRequest() string {
	return fmt.Sprintf("agent getposition")
}

// AgentMoveRequest produces the command used to move the agent of a player in a direction.
func AgentMoveRequest(direction mctype.Direction) string {
	return fmt.Sprintf("agent move %v", direction)
}

// AgentTurnRequest produces the command used to turn an agent. The direction must be either left or right.
func AgentTurnRequest(direction mctype.Direction) string {
	return fmt.Sprintf("agent turn %v", direction)
}

// AgentAttackRequest produces the command used to make an agent attack.
func AgentAttackRequest(direction mctype.Direction) string {
	return fmt.Sprintf("agent attack %v", direction)
}

// AgentPosition is sent by the server to get the position of the agent of a player.
type AgentPosition struct {
	// YRotation is the rotation on the Y axis of the agent. (yaw) This is always a full number.
	YRotation float64 `json:"y-rot"`
	// Position is the position of the agent. These are always full numbers.
	Position mctype.Position `json:"position"`
	// StatusCode is the status code of the command. 0 on success.
	StatusCode int `json:"statusCode"`
	// StatusMessage indicates if the command was successful with a message.
	StatusMessage string `json:"statusMessage"`
}

// AgentInstruction is a shared structure for agent commands that instruct the agent to do an action.
type AgentInstruction struct {
	// StatusCode is the status code of the command. 0 on success.
	StatusCode int `json:"statusCode"`
	// StatusMessage indicates if the command was successful with a message.
	StatusMessage string `json:"statusMessage"`
}

// AgentMove is sent by the server to move the agent of a player.
type AgentMove AgentInstruction

// AgentTurn is sent by the server to turn the agent of a player. The fields are the same as those of the
// AgentMove response.
type AgentTurn AgentInstruction

// AgentAttack is sent by the server to make the agent of a player attack in a direction.
type AgentAttack AgentInstruction
