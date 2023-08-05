// Package protocol implements the DataStore protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	datastore_types "github.com/PretendoNetwork/nex-protocols-go/datastore/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetMeta sets the GetMeta handler function
func (protocol *Protocol) GetMeta(handler func(err error, client *nex.Client, callID uint32, param *datastore_types.DataStoreGetMetaParam) uint32) {
	protocol.getMetaHandler = handler
}

func (protocol *Protocol) handleGetMeta(packet nex.PacketInterface) {
	if protocol.getMetaHandler == nil {
		globals.Logger.Warning("DataStore::GetMeta not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	param, err := parametersStream.ReadStructure(datastore_types.NewDataStoreGetMetaParam())
	if err != nil {
		go protocol.getMetaHandler(fmt.Errorf("Failed to read param from parameters. %s", err.Error()), client, callID, nil)
		return
	}

	go protocol.getMetaHandler(nil, client, callID, param.(*datastore_types.DataStoreGetMetaParam))
}
