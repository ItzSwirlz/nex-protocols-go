// Package match_making implements the Match Making NEX protocol
package match_making

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// CancelInvitation sets the CancelInvitation handler function
func (protocol *MatchMakingProtocol) CancelInvitation(handler func(err error, client *nex.Client, callID uint32, idGathering uint32, lstPrincipals []uint32, strMessage string)) {
	protocol.cancelInvitationHandler = handler
}

func (protocol *MatchMakingProtocol) handleCancelInvitation(packet nex.PacketInterface) {
	if protocol.cancelInvitationHandler == nil {
		globals.Logger.Warning("MatchMaking::CancelInvitation not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	idGathering, err := parametersStream.ReadUInt32LE()
	if err != nil {
		go protocol.cancelInvitationHandler(fmt.Errorf("Failed to read idGathering from parameters. %s", err.Error()), client, callID, 0, nil, "")
	}

	lstPrincipals, err := parametersStream.ReadListUInt32LE()
	if err != nil {
		go protocol.cancelInvitationHandler(fmt.Errorf("Failed to read lstPrincipals from parameters. %s", err.Error()), client, callID, 0, nil, "")
	}

	strMessage, err := parametersStream.ReadString()
	if err != nil {
		go protocol.cancelInvitationHandler(fmt.Errorf("Failed to read strMessage from parameters. %s", err.Error()), client, callID, 0, nil, "")
	}

	go protocol.cancelInvitationHandler(nil, client, callID, idGathering, lstPrincipals, strMessage)
}
