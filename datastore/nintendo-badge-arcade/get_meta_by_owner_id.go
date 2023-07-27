// Package protocol implements the Nintendo Badge Arcade DataStore protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	datastore_nintendo_badge_arcade_types "github.com/PretendoNetwork/nex-protocols-go/datastore/nintendo-badge-arcade/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetMetaByOwnerID sets the GetMetaByOwnerID function
func (protocol *Protocol) GetMetaByOwnerID(handler func(err error, client *nex.Client, callID uint32, param *datastore_nintendo_badge_arcade_types.DataStoreGetMetaByOwnerIDParam)) {
	protocol.getMetaByOwnerIDHandler = handler
}

func (protocol *Protocol) handleGetMetaByOwnerID(packet nex.PacketInterface) {
	if protocol.getMetaByOwnerIDHandler == nil {
		globals.Logger.Warning("DataStoreBadgeArcade::GetMetaByOwnerID not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	param, err := parametersStream.ReadStructure(datastore_nintendo_badge_arcade_types.NewDataStoreGetMetaByOwnerIDParam())
	if err != nil {
		go protocol.getMetaByOwnerIDHandler(fmt.Errorf("Failed to read param from parameters. %s", err.Error()), client, callID, nil)
		return
	}

	go protocol.getMetaByOwnerIDHandler(nil, client, callID, param.(*datastore_nintendo_badge_arcade_types.DataStoreGetMetaByOwnerIDParam))
}
