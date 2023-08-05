// Package protocol implements the Matchmake Extension protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
	match_making_types "github.com/PretendoNetwork/nex-protocols-go/match-making/types"
)

// JoinMatchmakeSessionWithParam sets the JoinMatchmakeSessionWithParam handler function
func (protocol *Protocol) JoinMatchmakeSessionWithParam(handler func(err error, client *nex.Client, callID uint32, joinMatchmakeSessionParam *match_making_types.JoinMatchmakeSessionParam) uint32) {
	protocol.joinMatchmakeSessionWithParamHandler = handler
}

func (protocol *Protocol) handleJoinMatchmakeSessionWithParam(packet nex.PacketInterface) {
	if protocol.joinMatchmakeSessionWithParamHandler == nil {
		globals.Logger.Warning("MatchmakeExtension::JoinMatchmakeSessionWithParam not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	joinMatchmakeSessionParam, err := parametersStream.ReadStructure(match_making_types.NewJoinMatchmakeSessionParam())
	if err != nil {
		go protocol.joinMatchmakeSessionWithParamHandler(fmt.Errorf("Failed to read joinMatchmakeSessionParam from parameters. %s", err.Error()), client, callID, nil)
		return
	}

	go protocol.joinMatchmakeSessionWithParamHandler(nil, client, callID, joinMatchmakeSessionParam.(*match_making_types.JoinMatchmakeSessionParam))
}
