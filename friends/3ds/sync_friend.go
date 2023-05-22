package friends_3ds

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// SyncFriend sets the SyncFriend handler function
func (protocol *Friends3DSProtocol) SyncFriend(handler func(err error, client *nex.Client, callID uint32, lfc uint64, pids []uint32, lfcList []uint64)) {
	protocol.SyncFriendHandler = handler
}

func (protocol *Friends3DSProtocol) HandleSyncFriend(packet nex.PacketInterface) {
	if protocol.SyncFriendHandler == nil {
		globals.Logger.Warning("Friends3DS::SyncFriend not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	lfc := parametersStream.ReadUInt64LE()
	pids := parametersStream.ReadListUInt32LE()
	lfcList := parametersStream.ReadListUInt64LE()

	go protocol.SyncFriendHandler(nil, client, callID, lfc, pids, lfcList)
}