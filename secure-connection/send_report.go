// Package protocol implements the Secure Connection protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// SendReport sets the SendReport handler function
func (protocol *Protocol) SendReport(handler func(err error, client *nex.Client, callID uint32, reportID uint32, reportData []byte) uint32) {
	protocol.sendReportHandler = handler
}

func (protocol *Protocol) handleSendReport(packet nex.PacketInterface) {
	if protocol.sendReportHandler == nil {
		globals.Logger.Warning("SecureConnection::SendReport not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	reportID, err := parametersStream.ReadUInt32LE()
	if err != nil {
		go protocol.sendReportHandler(fmt.Errorf("Failed to read reportID from parameters. %s", err.Error()), client, callID, 0, nil)
		return
	}

	reportData, err := parametersStream.ReadQBuffer()
	if err != nil {
		go protocol.sendReportHandler(fmt.Errorf("Failed to read reportData from parameters. %s", err.Error()), client, callID, 0, nil)
		return
	}

	go protocol.sendReportHandler(nil, client, callID, reportID, reportData)
}
