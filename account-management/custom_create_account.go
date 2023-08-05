// Package protocol implements the Account Management protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// CustomCreateAccount sets the CustomCreateAccount handler function
func (protocol *Protocol) CustomCreateAccount(handler func(err error, client *nex.Client, callID uint32, strPrincipalName string, strKey string, uiGroups uint32, strEmail string, oAuthData *nex.DataHolder) uint32) {
	protocol.customCreateAccountHandler = handler
}

func (protocol *Protocol) handleCustomCreateAccount(packet nex.PacketInterface) {
	if protocol.customCreateAccountHandler == nil {
		globals.Logger.Warning("AccountManagement::CustomCreateAccount not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	strPrincipalName, err := parametersStream.ReadString()
	if err != nil {
		go protocol.customCreateAccountHandler(fmt.Errorf("Failed to read strPrincipalName from parameters. %s", err.Error()), client, callID, "", "", 0, "", nil)
		return
	}

	strKey, err := parametersStream.ReadString()
	if err != nil {
		go protocol.customCreateAccountHandler(fmt.Errorf("Failed to read strKey from parameters. %s", err.Error()), client, callID, "", "", 0, "", nil)
		return
	}

	uiGroups, err := parametersStream.ReadUInt32LE()
	if err != nil {
		go protocol.customCreateAccountHandler(fmt.Errorf("Failed to read uiGroups from parameters. %s", err.Error()), client, callID, "", "", 0, "", nil)
		return
	}

	strEmail, err := parametersStream.ReadString()
	if err != nil {
		go protocol.customCreateAccountHandler(fmt.Errorf("Failed to read strEmail from parameters. %s", err.Error()), client, callID, "", "", 0, "", nil)
		return
	}

	oAuthData, err := parametersStream.ReadDataHolder()
	if err != nil {
		go protocol.customCreateAccountHandler(fmt.Errorf("Failed to read oAuthData from parameters. %s", err.Error()), client, callID, "", "", 0, "", nil)
		return
	}

	go protocol.customCreateAccountHandler(nil, client, callID, strPrincipalName, strKey, uiGroups, strEmail, oAuthData)
}
