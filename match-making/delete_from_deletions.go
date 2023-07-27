// Package protocol implements the Match Making protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// DeleteFromDeletions sets the DeleteFromDeletions handler function
func (protocol *Protocol) DeleteFromDeletions(handler func(err error, client *nex.Client, callID uint32, lstDeletions []uint32)) {
	protocol.deleteFromDeletionsHandler = handler
}

func (protocol *Protocol) handleDeleteFromDeletions(packet nex.PacketInterface) {
	if protocol.deleteFromDeletionsHandler == nil {
		globals.Logger.Warning("MatchMaking::DeleteFromDeletions not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	lstDeletions, err := parametersStream.ReadListUInt32LE()
	if err != nil {
		go protocol.deleteFromDeletionsHandler(fmt.Errorf("Failed to read lstDeletions from parameters. %s", err.Error()), client, callID, nil)
	}

	go protocol.deleteFromDeletionsHandler(nil, client, callID, lstDeletions)
}
