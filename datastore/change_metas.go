// Package protocol implements the DataStore protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	datastore_types "github.com/PretendoNetwork/nex-protocols-go/datastore/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// ChangeMetas sets the ChangeMetas handler function
func (protocol *Protocol) ChangeMetas(handler func(err error, client *nex.Client, callID uint32, dataIDs []uint64, params []*datastore_types.DataStoreChangeMetaParam, transactional bool)) {
	protocol.changeMetasHandler = handler
}

func (protocol *Protocol) handleChangeMetas(packet nex.PacketInterface) {
	if protocol.changeMetasHandler == nil {
		globals.Logger.Warning("DataStore::ChangeMetas not implemented")
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
		go protocol.changeMetasHandler(fmt.Errorf("Failed to read dataIDs from parameters. %s", err.Error()), client, callID, nil, nil, false)
		return
	}

	params, err := parametersStream.ReadListStructure(datastore_types.NewDataStoreChangeMetaParam())
	if err != nil {
		go protocol.changeMetasHandler(fmt.Errorf("Failed to read params from parameters. %s", err.Error()), client, callID, nil, nil, false)
		return
	}

	transactional, err := parametersStream.ReadBool()
	if err != nil {
		go protocol.changeMetasHandler(fmt.Errorf("Failed to read transactional from parameters. %s", err.Error()), client, callID, nil, nil, false)
		return
	}

	go protocol.changeMetasHandler(nil, client, callID, dataIDs, params.([]*datastore_types.DataStoreChangeMetaParam), transactional)
}