// Package protocol implements the NAT Traversal protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// ReportNATTraversalResultDetail sets the ReportNATTraversalResultDetail handler function
func (protocol *Protocol) ReportNATTraversalResultDetail(handler func(err error, client *nex.Client, callID uint32, cid uint32, result bool, detail int32, rtt uint32)) {
	protocol.reportNATTraversalResultDetailHandler = handler
}

func (protocol *Protocol) handleReportNATTraversalResultDetail(packet nex.PacketInterface) {
	if protocol.reportNATTraversalResultDetailHandler == nil {
		globals.Logger.Warning("NATTraversal::ReportNATTraversalResultDetail not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	// TODO - The NEX server should add a NATTraversalProtocolVersion method
	matchmakingVersion := protocol.Server.MatchMakingProtocolVersion()

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	cid, err := parametersStream.ReadUInt32LE()
	if err != nil {
		go protocol.reportNATTraversalResultDetailHandler(fmt.Errorf("Failed to read cid from parameters. %s", err.Error()), client, callID, 0, false, 0, 0)
		return
	}

	result, err := parametersStream.ReadBool()
	if err != nil {
		go protocol.reportNATTraversalResultDetailHandler(fmt.Errorf("Failed to read result from parameters. %s", err.Error()), client, callID, 0, false, 0, 0)
		return
	}

	detail, err := parametersStream.ReadInt32LE()
	if err != nil {
		go protocol.reportNATTraversalResultDetailHandler(fmt.Errorf("Failed to read detail from parameters. %s", err.Error()), client, callID, 0, false, 0, 0)
		return
	}

	var rtt uint32 = 0

	// TODO - Is this the right version?
	if matchmakingVersion.Major >= 3 && matchmakingVersion.Minor >= 0 {
		rtt, err = parametersStream.ReadUInt32LE()
		if err != nil {
			go protocol.reportNATTraversalResultDetailHandler(fmt.Errorf("Failed to read rtt from parameters. %s", err.Error()), client, callID, 0, false, 0, 0)
			return
		}
	}

	go protocol.reportNATTraversalResultDetailHandler(nil, client, callID, cid, result, detail, rtt)
}
