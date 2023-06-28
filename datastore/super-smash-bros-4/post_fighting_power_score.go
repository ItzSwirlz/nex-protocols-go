package datastore_super_smash_bros_4

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	datastore_super_smash_bros_4_types "github.com/PretendoNetwork/nex-protocols-go/datastore/super-smash-bros-4/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// PostFightingPowerScore sets the PostFightingPowerScore handler function
func (protocol *DataStoreSuperSmashBros4Protocol) PostFightingPowerScore(handler func(err error, client *nex.Client, callID uint32, params []*datastore_super_smash_bros_4_types.DataStorePostFightingPowerScoreParam)) {
	protocol.PostFightingPowerScoreHandler = handler
}

func (protocol *DataStoreSuperSmashBros4Protocol) HandlePostFightingPowerScore(packet nex.PacketInterface) {
	if protocol.PostFightingPowerScoreHandler == nil {
		globals.Logger.Warning("DataStoreSmash4::PostFightingPowerScore not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	params, err := parametersStream.ReadListStructure(datastore_super_smash_bros_4_types.NewDataStorePostFightingPowerScoreParam())
	if err != nil {
		go protocol.PostFightingPowerScoreHandler(fmt.Errorf("Failed to read params from parameters. %s", err.Error()), client, callID, nil)
		return
	}

	go protocol.PostFightingPowerScoreHandler(nil, client, callID, params.([]*datastore_super_smash_bros_4_types.DataStorePostFightingPowerScoreParam))
}
