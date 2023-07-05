// Package ticket_granting implements the Ticket Granting NEX protocol
package ticket_granting

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetName sets the GetName handler function
func (protocol *TicketGrantingProtocol) GetName(handler func(err error, client *nex.Client, callID uint32, userPID uint32)) {
	protocol.GetNameHandler = handler
}

func (protocol *TicketGrantingProtocol) handleGetName(packet nex.PacketInterface) {
	if protocol.GetNameHandler == nil {
		globals.Logger.Warning("TicketGranting::GetName not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	id, err := parametersStream.ReadUInt32LE()
	if err != nil {
		go protocol.GetNameHandler(fmt.Errorf("Failed to read id from parameters. %s", err.Error()), client, callID, 0)
		return
	}

	go protocol.GetNameHandler(nil, client, callID, id)
}