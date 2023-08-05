// Package protocol implements the Account Management protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// ChangePassword sets the ChangePassword handler function
func (protocol *Protocol) ChangePassword(handler func(err error, client *nex.Client, callID uint32, strNewKey string) uint32) {
	protocol.changePasswordHandler = handler
}

func (protocol *Protocol) handleChangePassword(packet nex.PacketInterface) {
	if protocol.changePasswordHandler == nil {
		globals.Logger.Warning("AccountManagement::ChangePassword not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	strNewKey, err := parametersStream.ReadString()
	if err != nil {
		go protocol.changePasswordHandler(fmt.Errorf("Failed to read strNewKey from parameters. %s", err.Error()), client, callID, "")
		return
	}

	go protocol.changePasswordHandler(nil, client, callID, strNewKey)
}
