// Package persistent_store implements the NAT Traversal NEX protocol
package persistent_store

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// FindByGroup sets the FindByGroup handler function
func (protocol *PersistentStoreProtocol) FindByGroup(handler func(err error, client *nex.Client, callID uint32, uiGroup uint32)) {
	protocol.findByGroupHandler = handler
}

func (protocol *PersistentStoreProtocol) handleFindByGroup(packet nex.PacketInterface) {
	if protocol.findByGroupHandler == nil {
		globals.Logger.Warning("PersistentStore::FindByGroup not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	uiGroup, err := parametersStream.ReadUInt32LE()
	if err != nil {
		go protocol.findByGroupHandler(fmt.Errorf("Failed to read uiGroup from parameters. %s", err.Error()), client, callID, 0)
		return
	}

	go protocol.findByGroupHandler(nil, client, callID, uiGroup)
}