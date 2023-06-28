package message_delivery_types

import (
	"fmt"

	"github.com/PretendoNetwork/nex-go"
)

type MessageRecipient struct {
	nex.Structure
	m_uiRecipientType uint32
	m_principalID     uint32
	m_gatheringID     uint32
}

// ExtractFromStream extracts a MessageRecipient structure from a stream
func (messageRecipient *MessageRecipient) ExtractFromStream(stream *nex.StreamIn) error {
	var err error

	messageRecipient.m_uiRecipientType, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract MessageRecipient.m_uiRecipientType from stream. %s", err.Error())
	}

	messageRecipient.m_principalID, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract MessageRecipient.m_principalID from stream. %s", err.Error())
	}

	messageRecipient.m_gatheringID, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract MessageRecipient.m_gatheringID from stream. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of MessageRecipient
func (messageRecipient *MessageRecipient) Copy() nex.StructureInterface {
	copied := NewMessageRecipient()

	copied.m_uiRecipientType = messageRecipient.m_uiRecipientType
	copied.m_principalID = messageRecipient.m_principalID
	copied.m_gatheringID = messageRecipient.m_gatheringID

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (messageRecipient *MessageRecipient) Equals(structure nex.StructureInterface) bool {
	other := structure.(*MessageRecipient)

	if messageRecipient.m_uiRecipientType != other.m_uiRecipientType {
		return false
	}

	if messageRecipient.m_principalID != other.m_principalID {
		return false
	}

	if messageRecipient.m_gatheringID != other.m_gatheringID {
		return false
	}

	return true
}

// NewMessageRecipient returns a new MessageRecipient
func NewMessageRecipient() *MessageRecipient {
	return &MessageRecipient{}
}
