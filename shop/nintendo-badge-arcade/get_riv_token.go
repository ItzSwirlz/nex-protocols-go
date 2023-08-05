// Package protocol implements the Nintendo Badge Arcade Shop protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetRivToken sets the GetRivToken function
func (protocol *Protocol) GetRivToken(handler func(err error, client *nex.Client, callID uint32, itemCode string, referenceID []byte) uint32) {
	protocol.getRivTokenHandler = handler
}

func (protocol *Protocol) handleGetRivToken(packet nex.PacketInterface) {
	if protocol.getRivTokenHandler == nil {
		globals.Logger.Warning("ShopNintendoBadgeArcade::GetRivToken not implemented")
		go globals.RespondErrorCustom(packet, CustomProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	itemCode, err := parametersStream.ReadString()
	if err != nil {
		go protocol.getRivTokenHandler(fmt.Errorf("Failed to read itemCode from parameters. %s", err.Error()), client, callID, "", nil)
		return
	}

	referenceID, err := parametersStream.ReadQBuffer()
	if err != nil {
		go protocol.getRivTokenHandler(fmt.Errorf("Failed to read referenceID from parameters. %s", err.Error()), client, callID, "", nil)
		return
	}

	go protocol.getRivTokenHandler(nil, client, callID, itemCode, referenceID)
}
