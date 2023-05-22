package matchmake_extension

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
	match_making "github.com/PretendoNetwork/nex-protocols-go/match-making"
)

// JoinMatchmakeSessionWithParam sets the JoinMatchmakeSessionWithParam handler function
func (protocol *MatchmakeExtensionProtocol) JoinMatchmakeSessionWithParam(handler func(err error, client *nex.Client, callID uint32, joinMatchmakeSessionParam *match_making.JoinMatchmakeSessionParam)) {
	protocol.JoinMatchmakeSessionWithParamHandler = handler
}

func (protocol *MatchmakeExtensionProtocol) HandleJoinMatchmakeSessionWithParam(packet nex.PacketInterface) {
	if protocol.JoinMatchmakeSessionWithParamHandler == nil {
		globals.Logger.Warning("MatchmakeExtension::JoinMatchmakeSessionWithParam not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	joinMatchmakeSessionParam, err := parametersStream.ReadStructure(match_making.NewJoinMatchmakeSessionParam())
	if err != nil {
		go protocol.JoinMatchmakeSessionWithParamHandler(err, client, callID, nil)
		return
	}

	go protocol.JoinMatchmakeSessionWithParamHandler(nil, client, callID, joinMatchmakeSessionParam.(*match_making.JoinMatchmakeSessionParam))
}