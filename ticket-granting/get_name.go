// Package protocol implements the Ticket Granting protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetName sets the GetName handler function
func (protocol *Protocol) GetName(handler func(err error, client *nex.Client, callID uint32, userPID uint32) uint32) {
	protocol.getNameHandler = handler
}

func (protocol *Protocol) handleGetName(packet nex.PacketInterface) {
	if protocol.getNameHandler == nil {
		globals.Logger.Warning("TicketGranting::GetName not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	id, err := parametersStream.ReadUInt32LE()
	if err != nil {
		go protocol.getNameHandler(fmt.Errorf("Failed to read id from parameters. %s", err.Error()), client, callID, 0)
		return
	}

	go protocol.getNameHandler(nil, client, callID, id)
}
