// Package protocol implements the DataStorePokemonBank protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// RequestMigration sets the RequestMigration handler function
func (protocol *Protocol) RequestMigration(handler func(err error, client *nex.Client, callID uint32, oneTimePassword string, boxes []uint32)) {
	protocol.requestMigrationHandler = handler
}

func (protocol *Protocol) handleRequestMigration(packet nex.PacketInterface) {
	if protocol.requestMigrationHandler == nil {
		globals.Logger.Warning("DataStorePokemonBank::RequestMigration not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	oneTimePassword, err := parametersStream.ReadString()
	if err != nil {
		go protocol.requestMigrationHandler(fmt.Errorf("Failed to read oneTimePassword from parameters. %s", err.Error()), client, callID, "", nil)
		return
	}

	boxes, err := parametersStream.ReadListUInt32LE()
	if err != nil {
		go protocol.requestMigrationHandler(fmt.Errorf("Failed to read boxes from parameters. %s", err.Error()), client, callID, "", nil)
		return
	}

	go protocol.requestMigrationHandler(nil, client, callID, oneTimePassword, boxes)
}
