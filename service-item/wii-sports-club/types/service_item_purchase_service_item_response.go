// Package types implements all the types used by the ServiceItem protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// ServiceItemPurchaseServiceItemResponse is a type within the ServiceItem protocol
type ServiceItemPurchaseServiceItemResponse struct {
	types.Structure
	*ServiceItemEShopResponse
	NullablePurchaceInfo *types.List[*ServiceItemPurchaceInfo]
}

// WriteTo writes the ServiceItemPurchaseServiceItemResponse to the given writable
func (sipsir *ServiceItemPurchaseServiceItemResponse) WriteTo(writable types.Writable) {
	sipsir.ServiceItemEShopResponse.WriteTo(writable)

	contentWritable := writable.CopyNew()

	sipsir.NullablePurchaceInfo.WriteTo(writable)

	content := contentWritable.Bytes()

	sipsir.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the ServiceItemPurchaseServiceItemResponse from the given readable
func (sipsir *ServiceItemPurchaseServiceItemResponse) ExtractFrom(readable types.Readable) error {
	var err error

	err = sipsir.ServiceItemEShopResponse.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemPurchaseServiceItemResponse.ServiceItemEShopResponse. %s", err.Error())
	}

	err = sipsir.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemPurchaseServiceItemResponse header. %s", err.Error())
	}

	err = sipsir.NullablePurchaceInfo.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemPurchaseServiceItemResponse.NullablePurchaceInfo. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of ServiceItemPurchaseServiceItemResponse
func (sipsir *ServiceItemPurchaseServiceItemResponse) Copy() types.RVType {
	copied := NewServiceItemPurchaseServiceItemResponse()

	copied.StructureVersion = sipsir.StructureVersion
	copied.ServiceItemEShopResponse = sipsir.ServiceItemEShopResponse.Copy().(*ServiceItemEShopResponse)
	copied.NullablePurchaceInfo = sipsir.NullablePurchaceInfo.Copy().(*types.List[*ServiceItemPurchaceInfo])

	return copied
}

// Equals checks if the given ServiceItemPurchaseServiceItemResponse contains the same data as the current ServiceItemPurchaseServiceItemResponse
func (sipsir *ServiceItemPurchaseServiceItemResponse) Equals(o types.RVType) bool {
	if _, ok := o.(*ServiceItemPurchaseServiceItemResponse); !ok {
		return false
	}

	other := o.(*ServiceItemPurchaseServiceItemResponse)

	if sipsir.StructureVersion != other.StructureVersion {
		return false
	}

	if !sipsir.ServiceItemEShopResponse.Equals(other.ServiceItemEShopResponse) {
		return false
	}

	return sipsir.NullablePurchaceInfo.Equals(other.NullablePurchaceInfo)
}

// String returns the string representation of the ServiceItemPurchaseServiceItemResponse
func (sipsir *ServiceItemPurchaseServiceItemResponse) String() string {
	return sipsir.FormatToString(0)
}

// FormatToString pretty-prints the ServiceItemPurchaseServiceItemResponse using the provided indentation level
func (sipsir *ServiceItemPurchaseServiceItemResponse) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("ServiceItemPurchaseServiceItemResponse{\n")
	b.WriteString(fmt.Sprintf("%sServiceItemEShopResponse (parent): %s,\n", indentationValues, sipsir.ServiceItemEShopResponse.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sNullablePurchaceInfo: %s,\n", indentationValues, sipsir.NullablePurchaceInfo))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewServiceItemPurchaseServiceItemResponse returns a new ServiceItemPurchaseServiceItemResponse
func NewServiceItemPurchaseServiceItemResponse() *ServiceItemPurchaseServiceItemResponse {
	sipsir := &ServiceItemPurchaseServiceItemResponse{
		ServiceItemEShopResponse: NewServiceItemEShopResponse(),
		NullablePurchaceInfo:     types.NewList[*ServiceItemPurchaceInfo](),
	}

	sipsir.NullablePurchaceInfo.Type = NewServiceItemPurchaceInfo()

	return sipsir
}
