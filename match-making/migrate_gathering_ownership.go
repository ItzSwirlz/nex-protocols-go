// Package protocol implements the Match Making protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// MigrateGatheringOwnership sets the MigrateGatheringOwnership handler function
func (protocol *Protocol) MigrateGatheringOwnership(handler func(err error, client *nex.Client, callID uint32, gid uint32, lstPotentialNewOwnersID []uint32, participantsOnly bool)) {
	protocol.migrateGatheringOwnershipHandler = handler
}

func (protocol *Protocol) handleMigrateGatheringOwnership(packet nex.PacketInterface) {
	if protocol.migrateGatheringOwnershipHandler == nil {
		globals.Logger.Warning("MatchMaking::MigrateGatheringOwnership not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	gid, err := parametersStream.ReadUInt32LE()
	if err != nil {
		go protocol.migrateGatheringOwnershipHandler(fmt.Errorf("Failed to read gid from parameters. %s", err.Error()), client, callID, 0, nil, false)
	}

	lstPotentialNewOwnersID, err := parametersStream.ReadListUInt32LE()
	if err != nil {
		go protocol.migrateGatheringOwnershipHandler(fmt.Errorf("Failed to read gid from parameters. %s", err.Error()), client, callID, 0, nil, false)
	}

	participantsOnly, err := parametersStream.ReadBool()
	if err != nil {
		go protocol.migrateGatheringOwnershipHandler(fmt.Errorf("Failed to read participantsOnly from parameters. %s", err.Error()), client, callID, 0, nil, false)
	}

	go protocol.migrateGatheringOwnershipHandler(nil, client, callID, gid, lstPotentialNewOwnersID, participantsOnly)
}
