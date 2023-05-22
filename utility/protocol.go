package utility

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

const (
	// ProtocolID is the protocol ID for the Utility protocol
	ProtocolID = 0x6e

	// MethodAcquireNexUniqueID is the method ID for the method AcquireNexUniqueID
	MethodAcquireNexUniqueID = 0x1

	// MethodAcquireNexUniqueIDWithPassword is the method ID for the method AcquireNexUniqueIDWithPassword
	MethodAcquireNexUniqueIDWithPassword = 0x2

	// MethodAssociateNexUniqueIDWithMyPrincipalID is the method ID for the method AssociateNexUniqueIDWithMyPrincipalID
	MethodAssociateNexUniqueIDWithMyPrincipalID = 0x3

	// MethodAssociateNexUniqueIDsWithMyPrincipalID is the method ID for the method AssociateNexUniqueIDsWithMyPrincipalID
	MethodAssociateNexUniqueIDsWithMyPrincipalID = 0x4

	// MethodGetAssociatedNexUniqueIDWithMyPrincipalID is the method ID for the method GetAssociatedNexUniqueIDWithMyPrincipalID
	MethodGetAssociatedNexUniqueIDWithMyPrincipalID = 0x5

	// MethodGetAssociatedNexUniqueIDsWithMyPrincipalID is the method ID for the method GetAssociatedNexUniqueIDsWithMyPrincipalID
	MethodGetAssociatedNexUniqueIDsWithMyPrincipalID = 0x6

	// MethodGetIntegerSettings is the method ID for the method GetIntegerSettings
	MethodGetIntegerSettings = 0x7

	// MethodGetStringSettings is the method ID for the method GetStringSettings
	MethodGetStringSettings = 0x8
)

// UtilityProtocol handles the Utility nex protocol
type UtilityProtocol struct {
	Server                                            *nex.Server
	AcquireNexUniqueIDHandler                         func(err error, client *nex.Client, callID uint32)
	AcquireNexUniqueIDWithPasswordHandler             func(err error, client *nex.Client, callID uint32)
	AssociateNexUniqueIDWithMyPrincipalIDHandler      func(err error, client *nex.Client, callID uint32, uniqueIDInfo *UniqueIDInfo)
	AssociateNexUniqueIDsWithMyPrincipalIDHandler     func(err error, client *nex.Client, callID uint32, uniqueIDInfo []*UniqueIDInfo)
	GetAssociatedNexUniqueIDWithMyPrincipalIDHandler  func(err error, client *nex.Client, callID uint32)
	GetAssociatedNexUniqueIDsWithMyPrincipalIDHandler func(err error, client *nex.Client, callID uint32)
	GetIntegerSettingsHandler                         func(err error, client *nex.Client, callID uint32, integerSettingIndex uint32)
	GetStringSettingsHandler                          func(err error, client *nex.Client, callID uint32, stringSettingIndex uint32)
}

// Setup initializes the protocol
func (protocol *UtilityProtocol) Setup() {
	protocol.Server.On("Data", func(packet nex.PacketInterface) {
		request := packet.RMCRequest()

		if request.ProtocolID() == ProtocolID {
			switch request.MethodID() {
			case MethodAcquireNexUniqueID:
				go protocol.HandleAcquireNexUniqueID(packet)
			case MethodAcquireNexUniqueIDWithPassword:
				go protocol.HandleAcquireNexUniqueIDWithPassword(packet)
			case MethodAssociateNexUniqueIDWithMyPrincipalID:
				go protocol.HandleAssociateNexUniqueIDWithMyPrincipalID(packet)
			case MethodAssociateNexUniqueIDsWithMyPrincipalID:
				go protocol.HandleAssociateNexUniqueIDsWithMyPrincipalID(packet)
			case MethodGetAssociatedNexUniqueIDWithMyPrincipalID:
				go protocol.HandleGetAssociatedNexUniqueIDWithMyPrincipalID(packet)
			case MethodGetAssociatedNexUniqueIDsWithMyPrincipalID:
				go protocol.HandleGetAssociatedNexUniqueIDsWithMyPrincipalID(packet)
			case MethodGetIntegerSettings:
				go protocol.HandleGetIntegerSettings(packet)
			case MethodGetStringSettings:
				go protocol.HandleGetStringSettings(packet)
			default:
				go globals.RespondNotImplemented(packet, ProtocolID)
				fmt.Printf("Unsupported Utility method ID: %#v\n", request.MethodID())
			}
		}
	})
}

// NewUtilityProtocol returns a new UtilityProtocol
func NewUtilityProtocol(server *nex.Server) *UtilityProtocol {
	protocol := &UtilityProtocol{Server: server}

	protocol.Setup()

	return protocol
}