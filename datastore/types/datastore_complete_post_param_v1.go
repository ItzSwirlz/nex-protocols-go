package datastore_types

import (
	"fmt"

	"github.com/PretendoNetwork/nex-go"
)

type DataStoreCompletePostParamV1 struct {
	nex.Structure
	DataID    uint32
	IsSuccess bool
}

// ExtractFromStream extracts a DataStoreCompletePostParamV1 structure from a stream
func (dataStoreCompletePostParamV1 *DataStoreCompletePostParamV1) ExtractFromStream(stream *nex.StreamIn) error {
	var err error

	dataStoreCompletePostParamV1.DataID, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreCompletePostParamV1.DataID. %s", err.Error())
	}

	dataStoreCompletePostParamV1.IsSuccess, err = stream.ReadBool()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreCompletePostParamV1.IsSuccess. %s", err.Error())
	}

	return nil
}

// Bytes encodes the DataStoreCompletePostParamV1 and returns a byte array
func (dataStoreCompletePostParamV1 *DataStoreCompletePostParamV1) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteUInt32LE(dataStoreCompletePostParamV1.DataID)
	stream.WriteBool(dataStoreCompletePostParamV1.IsSuccess)

	return stream.Bytes()
}

// Copy returns a new copied instance of DataStoreCompletePostParamV1
func (dataStoreCompletePostParamV1 *DataStoreCompletePostParamV1) Copy() nex.StructureInterface {
	copied := NewDataStoreCompletePostParamV1()

	copied.DataID = dataStoreCompletePostParamV1.DataID
	copied.IsSuccess = dataStoreCompletePostParamV1.IsSuccess

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (dataStoreCompletePostParamV1 *DataStoreCompletePostParamV1) Equals(structure nex.StructureInterface) bool {
	other := structure.(*DataStoreCompletePostParamV1)

	if dataStoreCompletePostParamV1.DataID != other.DataID {
		return false
	}

	if dataStoreCompletePostParamV1.IsSuccess != other.IsSuccess {
		return false
	}

	return true
}

// NewDataStoreCompletePostParamV1 returns a new DataStoreCompletePostParamV1
func NewDataStoreCompletePostParamV1() *DataStoreCompletePostParamV1 {
	return &DataStoreCompletePostParamV1{}
}
