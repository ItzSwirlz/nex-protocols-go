// Package protocol implements the Ticket Granting protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// RequestTicket sets the RequestTicket handler function
func (protocol *Protocol) RequestTicket(handler func(err error, client *nex.Client, callID uint32, idSource uint32, idTarget uint32) uint32) {
	protocol.requestTicketHandler = handler
}

func (protocol *Protocol) handleRequestTicket(packet nex.PacketInterface) {
	if protocol.requestTicketHandler == nil {
		globals.Logger.Warning("TicketGranting::RequestTicket not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	idSource, err := parametersStream.ReadUInt32LE()
	if err != nil {
		go protocol.requestTicketHandler(fmt.Errorf("Failed to read idSource from parameters. %s", err.Error()), client, callID, 0, 0)
		return
	}

	idTarget, err := parametersStream.ReadUInt32LE()
	if err != nil {
		go protocol.requestTicketHandler(fmt.Errorf("Failed to read idTarget from parameters. %s", err.Error()), client, callID, 0, 0)
		return
	}

	go protocol.requestTicketHandler(nil, client, callID, idSource, idTarget)
}
