// Package protocol implements the DataStorePokemonBank protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	datastore_pokemon_bank_types "github.com/PretendoNetwork/nex-protocols-go/datastore/pokemon-bank/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// UploadPokemon sets the UploadPokemon handler function
func (protocol *Protocol) UploadPokemon(handler func(err error, client *nex.Client, callID uint32, param *datastore_pokemon_bank_types.GlobalTradeStationUploadPokemonParam) uint32) {
	protocol.uploadPokemonHandler = handler
}

func (protocol *Protocol) handleUploadPokemon(packet nex.PacketInterface) {
	if protocol.uploadPokemonHandler == nil {
		globals.Logger.Warning("DataStorePokemonBank::UploadPokemon not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	param, err := parametersStream.ReadStructure(datastore_pokemon_bank_types.NewGlobalTradeStationUploadPokemonParam())
	if err != nil {
		go protocol.uploadPokemonHandler(fmt.Errorf("Failed to read param from parameters. %s", err.Error()), client, callID, nil)
		return
	}

	go protocol.uploadPokemonHandler(nil, client, callID, param.(*datastore_pokemon_bank_types.GlobalTradeStationUploadPokemonParam))
}
