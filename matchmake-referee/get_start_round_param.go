// Package protocol implements the Matchmake Referee protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetStartRoundParam sets the GetStartRoundParam handler function
func (protocol *Protocol) GetStartRoundParam(handler func(err error, client *nex.Client, callID uint32, roundID uint64)) {
	protocol.getStartRoundParamHandler = handler
}

func (protocol *Protocol) handleGetStartRoundParam(packet nex.PacketInterface) {
	if protocol.getStartRoundParamHandler == nil {
		globals.Logger.Warning("MatchmakeReferee::GetStartRoundParam not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	roundID, err := parametersStream.ReadUInt64LE()
	if err != nil {
		go protocol.getStartRoundParamHandler(fmt.Errorf("Failed to read roundID from parameters. %s", err.Error()), client, callID, 0)
		return
	}

	go protocol.getStartRoundParamHandler(nil, client, callID, roundID)
}