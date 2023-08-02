// Package protocol implements the Ranking2 protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetCategorySetting sets the GetCategorySetting handler function
func (protocol *Protocol) GetCategorySetting(handler func(err error, client *nex.Client, callID uint32, category uint32)) {
	protocol.getCategorySettingHandler = handler
}

func (protocol *Protocol) handleGetCategorySetting(packet nex.PacketInterface) {
	if protocol.getCategorySettingHandler == nil {
		globals.Logger.Warning("Ranking2::GetCategorySetting not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}
	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	category, err := parametersStream.ReadUInt32LE()
	if err != nil {
		go protocol.getCategorySettingHandler(fmt.Errorf("Failed to read category from parameters. %s", err.Error()), client, callID, 0)
		return
	}

	go protocol.getCategorySettingHandler(nil, client, callID, category)
}