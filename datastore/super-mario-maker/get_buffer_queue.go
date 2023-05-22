package datastore_super_mario_maker

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetBufferQueue sets the GetBufferQueue handler function
func (protocol *DataStoreSuperMarioMakerProtocol) GetBufferQueue(handler func(err error, client *nex.Client, callID uint32, bufferQueueParam *BufferQueueParam)) {
	protocol.GetBufferQueueHandler = handler
}

func (protocol *DataStoreSuperMarioMakerProtocol) HandleGetBufferQueue(packet nex.PacketInterface) {
	if protocol.GetBufferQueueHandler == nil {
		globals.Logger.Warning("DataStoreSMM::GetBufferQueue not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	bufferQueueParam, err := parametersStream.ReadStructure(NewBufferQueueParam())

	if err != nil {
		go protocol.GetBufferQueueHandler(err, client, callID, nil)
		return
	}

	go protocol.GetBufferQueueHandler(nil, client, callID, bufferQueueParam.(*BufferQueueParam))
}