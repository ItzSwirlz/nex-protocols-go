// Package protocol implements the DataStorePokemonBank protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetTransactionParam sets the GetTransactionParam handler function
func (protocol *Protocol) GetTransactionParam(handler func(err error, client *nex.Client, callID uint32, slotID uint16)) {
	protocol.getTransactionParamHandler = handler
}

func (protocol *Protocol) handleGetTransactionParam(packet nex.PacketInterface) {
	if protocol.getTransactionParamHandler == nil {
		globals.Logger.Warning("DataStorePokemonBank::GetTransactionParam not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	slotID, err := parametersStream.ReadUInt16LE()
	if err != nil {
		go protocol.getTransactionParamHandler(fmt.Errorf("Failed to read slotID from parameters. %s", err.Error()), client, callID, 0)
		return
	}

	go protocol.getTransactionParamHandler(nil, client, callID, slotID)
}