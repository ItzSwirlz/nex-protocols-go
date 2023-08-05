// Package protocol implements the ServiceItemTeamKirbyClashDeluxe protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// HTTPGetResponse sets the HTTPGetResponse handler function
func (protocol *Protocol) HTTPGetResponse(handler func(err error, client *nex.Client, callID uint32, requestID uint32) uint32) {
	protocol.httpGetResponseHandler = handler
}

func (protocol *Protocol) handleHTTPGetResponse(packet nex.PacketInterface) {
	if protocol.httpGetResponseHandler == nil {
		globals.Logger.Warning("ServiceItemTeamKirbyClashDeluxe::HTTPGetResponse not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	requestID, err := parametersStream.ReadUInt32LE()
	if err != nil {
		go protocol.httpGetResponseHandler(fmt.Errorf("Failed to read requestID from parameters. %s", err.Error()), client, callID, 0)
		return
	}

	go protocol.httpGetResponseHandler(nil, client, callID, requestID)
}
