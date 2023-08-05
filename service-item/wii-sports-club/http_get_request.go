// Package protocol implements the ServiceItemWiiSportsClub protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
	service_item_wii_sports_club_types "github.com/PretendoNetwork/nex-protocols-go/service-item/wii-sports-club/types"
)

// HTTPGetRequest sets the HTTPGetRequest handler function
func (protocol *Protocol) HTTPGetRequest(handler func(err error, client *nex.Client, callID uint32, url *service_item_wii_sports_club_types.ServiceItemHTTPGetParam) uint32) {
	protocol.httpGetRequestHandler = handler
}

func (protocol *Protocol) handleHTTPGetRequest(packet nex.PacketInterface) {
	if protocol.httpGetRequestHandler == nil {
		globals.Logger.Warning("ServiceItemWiiSportsClub::HTTPGetRequest not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	url, err := parametersStream.ReadStructure(service_item_wii_sports_club_types.NewServiceItemHTTPGetParam())
	if err != nil {
		go protocol.httpGetRequestHandler(fmt.Errorf("Failed to read url from parameters. %s", err.Error()), client, callID, nil)
		return
	}

	go protocol.httpGetRequestHandler(nil, client, callID, url.(*service_item_wii_sports_club_types.ServiceItemHTTPGetParam))
}
