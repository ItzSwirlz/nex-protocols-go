// Package protocol implements the DataStoreSuperMarioMaker protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	datastore_super_mario_maker_types "github.com/PretendoNetwork/nex-protocols-go/datastore/super-mario-maker/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// ReportCourse sets the ReportCourse handler function
func (protocol *Protocol) ReportCourse(handler func(err error, client *nex.Client, callID uint32, param *datastore_super_mario_maker_types.DataStoreReportCourseParam) uint32) {
	protocol.reportCourseHandler = handler
}

func (protocol *Protocol) handleReportCourse(packet nex.PacketInterface) {
	if protocol.reportCourseHandler == nil {
		globals.Logger.Warning("DataStoreSuperMarioMaker::ReportCourse not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	param, err := parametersStream.ReadStructure(datastore_super_mario_maker_types.NewDataStoreReportCourseParam())
	if err != nil {
		go protocol.reportCourseHandler(fmt.Errorf("Failed to read param from parameters. %s", err.Error()), client, callID, nil)
		return
	}

	go protocol.reportCourseHandler(nil, client, callID, param.(*datastore_super_mario_maker_types.DataStoreReportCourseParam))
}
