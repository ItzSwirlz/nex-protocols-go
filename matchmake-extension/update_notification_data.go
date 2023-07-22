// Package matchmake_extension implements the Matchmake Extension NEX protocol
package matchmake_extension

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// UpdateNotificationData sets the UpdateNotificationData handler function
func (protocol *MatchmakeExtensionProtocol) UpdateNotificationData(handler func(err error, client *nex.Client, callID uint32, uiType uint32, uiParam1 uint32, uiParam2 uint32, strParam string)) {
	protocol.updateNotificationDataHandler = handler
}

func (protocol *MatchmakeExtensionProtocol) handleUpdateNotificationData(packet nex.PacketInterface) {
	if protocol.updateNotificationDataHandler == nil {
		globals.Logger.Warning("MatchmakeExtension::UpdateNotificationData not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	uiType, err := parametersStream.ReadUInt32LE()
	if err != nil {
		go protocol.updateNotificationDataHandler(fmt.Errorf("Failed to read uiType from parameters. %s", err.Error()), client, callID, 0, 0, 0, "")
		return
	}

	uiParam1, err := parametersStream.ReadUInt32LE()
	if err != nil {
		go protocol.updateNotificationDataHandler(fmt.Errorf("Failed to read uiParam1 from parameters. %s", err.Error()), client, callID, 0, 0, 0, "")
		return
	}

	uiParam2, err := parametersStream.ReadUInt32LE()
	if err != nil {
		go protocol.updateNotificationDataHandler(fmt.Errorf("Failed to read uiParam2 from parameters. %s", err.Error()), client, callID, 0, 0, 0, "")
		return
	}

	strParam, err := parametersStream.ReadString()
	if err != nil {
		go protocol.updateNotificationDataHandler(fmt.Errorf("Failed to read strParam from parameters. %s", err.Error()), client, callID, 0, 0, 0, "")
		return
	}

	go protocol.updateNotificationDataHandler(nil, client, callID, uiType, uiParam1, uiParam2, strParam)
}
