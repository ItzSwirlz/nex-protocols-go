// Package protocol implements the DataStoreSuperMarioMaker protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	datastore_super_mario_maker_types "github.com/PretendoNetwork/nex-protocols-go/datastore/super-mario-maker/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetCustomRanking sets the GetCustomRanking handler function
func (protocol *Protocol) GetCustomRanking(handler func(err error, packet nex.PacketInterface, callID uint32, param *datastore_super_mario_maker_types.DataStoreGetCustomRankingParam) uint32) {
	protocol.getCustomRankingHandler = handler
}

func (protocol *Protocol) handleGetCustomRanking(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.getCustomRankingHandler == nil {
		globals.Logger.Warning("DataStoreSuperMarioMaker::GetCustomRanking not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	param, err := parametersStream.ReadStructure(datastore_super_mario_maker_types.NewDataStoreGetCustomRankingParam())
	if err != nil {
		errorCode = protocol.getCustomRankingHandler(fmt.Errorf("Failed to read param from parameters. %s", err.Error()), packet, callID, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.getCustomRankingHandler(nil, packet, callID, param.(*datastore_super_mario_maker_types.DataStoreGetCustomRankingParam))
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
