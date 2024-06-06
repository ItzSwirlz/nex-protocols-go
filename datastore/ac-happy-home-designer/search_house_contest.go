// Package protocol implements the Animal Crossing: Happy Home Designer protocol
package protocol

import (
	"fmt"

	"github.com/PretendoNetwork/nex-go/v2"
	datastore_ac_happy_home_designer_types "github.com/PretendoNetwork/nex-protocols-go/v2/datastore/ac-happy-home-designer/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
)

func (protocol *Protocol) handleSearchHouseContest(packet nex.PacketInterface) {
	if protocol.SearchHouseContest == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "DataStoreACHappyHomeDesigner::SearchHouseContest not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	param := datastore_ac_happy_home_designer_types.NewDataStoreSearchHouseParam()

	err := param.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.SearchHouseContest(fmt.Errorf("Failed to read param from parameters. %s", err.Error()), packet, callID, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.SearchHouseContest(nil, packet, callID, param)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
