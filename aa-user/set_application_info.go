// Package protocol implements the AAUser protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	aauser_types "github.com/PretendoNetwork/nex-protocols-go/aa-user/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// SetApplicationInfo sets the SetApplicationInfo handler function
func (protocol *Protocol) SetApplicationInfo(handler func(err error, client *nex.Client, callID uint32, applicationInfo []*aauser_types.ApplicationInfo)) {
	protocol.setApplicationInfoHandler = handler
}

func (protocol *Protocol) handleSetApplicationInfo(packet nex.PacketInterface) {
	if protocol.setApplicationInfoHandler == nil {
		globals.Logger.Warning("AAUser::SetApplicationInfo not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	applicationInfo, err := parametersStream.ReadListStructure(aauser_types.NewApplicationInfo())
	if err != nil {
		go protocol.setApplicationInfoHandler(fmt.Errorf("Failed to read applicationInfo from parameters. %s", err.Error()), client, callID, nil)
		return
	}

	go protocol.setApplicationInfoHandler(nil, client, callID, applicationInfo.([]*aauser_types.ApplicationInfo))
}
