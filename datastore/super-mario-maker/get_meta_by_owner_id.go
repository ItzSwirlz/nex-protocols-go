// Package protocol implements the DataStoreSuperMarioMaker protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	datastore_super_mario_maker_types "github.com/PretendoNetwork/nex-protocols-go/datastore/super-mario-maker/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetMetaByOwnerID sets the GetMetaByOwnerID handler function
func (protocol *Protocol) GetMetaByOwnerID(handler func(err error, client *nex.Client, callID uint32, param *datastore_super_mario_maker_types.DataStoreGetMetaByOwnerIDParam) uint32) {
	protocol.getMetaByOwnerIDHandler = handler
}

func (protocol *Protocol) handleGetMetaByOwnerID(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.getMetaByOwnerIDHandler == nil {
		globals.Logger.Warning("DataStoreSuperMarioMaker::GetMetaByOwnerID not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	param, err := parametersStream.ReadStructure(datastore_super_mario_maker_types.NewDataStoreGetMetaByOwnerIDParam())
	if err != nil {
		errorCode = protocol.getMetaByOwnerIDHandler(fmt.Errorf("Failed to read param from parameters. %s", err.Error()), client, callID, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.getMetaByOwnerIDHandler(nil, client, callID, param.(*datastore_super_mario_maker_types.DataStoreGetMetaByOwnerIDParam))
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}