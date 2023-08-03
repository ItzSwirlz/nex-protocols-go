// Package protocol implements the DataStore protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	datastore_types "github.com/PretendoNetwork/nex-protocols-go/datastore/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// ResetRatings sets the ResetRatings handler function
func (protocol *Protocol) ResetRatings(handler func(err error, client *nex.Client, callID uint32, target *datastore_types.DataStoreRatingTarget, transactional bool)) {
	protocol.resetRatingsHandler = handler
}

func (protocol *Protocol) handleResetRatings(packet nex.PacketInterface) {
	if protocol.resetRatingsHandler == nil {
		globals.Logger.Warning("DataStore::ResetRatings not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	target, err := parametersStream.ReadStructure(datastore_types.NewDataStoreRatingTarget())
	if err != nil {
		go protocol.resetRatingsHandler(fmt.Errorf("Failed to read target from parameters. %s", err.Error()), client, callID, nil, false)
		return
	}

	transactional, err := parametersStream.ReadBool()
	if err != nil {
		go protocol.resetRatingsHandler(fmt.Errorf("Failed to read transactional from parameters. %s", err.Error()), client, callID, nil, false)
		return
	}

	go protocol.resetRatingsHandler(nil, client, callID, target.(*datastore_types.DataStoreRatingTarget), transactional)
}