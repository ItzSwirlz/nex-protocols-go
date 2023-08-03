// Package protocol implements the DataStore protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetRatings sets the GetRatings handler function
func (protocol *Protocol) GetRatings(handler func(err error, client *nex.Client, callID uint32, dataIDs []uint64, accessPassword uint64)) {
	protocol.getRatingsHandler = handler
}

func (protocol *Protocol) handleGetRatings(packet nex.PacketInterface) {
	if protocol.getRatingsHandler == nil {
		globals.Logger.Warning("DataStore::GetRatings not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	dataIDs, err := parametersStream.ReadListUInt64LE()
	if err != nil {
		go protocol.getRatingsHandler(fmt.Errorf("Failed to read dataIDs from parameters. %s", err.Error()), client, callID, nil, 0)
		return
	}

	accessPassword, err := parametersStream.ReadUInt64LE()
	if err != nil {
		go protocol.getRatingsHandler(fmt.Errorf("Failed to read accessPassword from parameters. %s", err.Error()), client, callID, nil, 0)
		return
	}

	go protocol.getRatingsHandler(nil, client, callID, dataIDs, accessPassword)
}