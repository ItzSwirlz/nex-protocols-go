// Package protocol implements the Super Mario Maker DataStore protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	datastore_super_mario_maker_types "github.com/PretendoNetwork/nex-protocols-go/datastore/super-mario-maker/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// UploadCourseRecord sets the UploadCourseRecord handler function
func (protocol *Protocol) UploadCourseRecord(handler func(err error, client *nex.Client, callID uint32, param *datastore_super_mario_maker_types.DataStoreUploadCourseRecordParam)) {
	protocol.uploadCourseRecordHandler = handler
}

func (protocol *Protocol) handleUploadCourseRecord(packet nex.PacketInterface) {
	if protocol.uploadCourseRecordHandler == nil {
		globals.Logger.Warning("DataStoreSMM::UploadCourseRecord not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	param, err := parametersStream.ReadStructure(datastore_super_mario_maker_types.NewDataStoreUploadCourseRecordParam())
	if err != nil {
		go protocol.uploadCourseRecordHandler(fmt.Errorf("Failed to read param from parameters. %s", err.Error()), client, callID, nil)
		return
	}

	go protocol.uploadCourseRecordHandler(nil, client, callID, param.(*datastore_super_mario_maker_types.DataStoreUploadCourseRecordParam))
}
