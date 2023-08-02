// Package protocol implements the Ranking2 protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
	ranking2_types "github.com/PretendoNetwork/nex-protocols-go/ranking2/types"
)

const (
	// ProtocolID is the protocol ID for the Ranking2 protocol
	ProtocolID = 0x7A

	// MethodPutScore is the method ID for the method PutScore
	MethodPutScore = 0x1

	// MethodGetCommonData is the method ID for the method GetCommonData
	MethodGetCommonData = 0x2

	// MethodPutCommonData is the method ID for the method PutCommonData
	MethodPutCommonData = 0x3

	// MethodDeleteCommonData is the method ID for the method DeleteCommonData
	MethodDeleteCommonData = 0x4

	// MethodGetRanking is the method ID for the method GetRanking
	MethodGetRanking = 0x5

	// MethodGetRankingByPrincipalId is the method ID for the method GetRankingByPrincipalId
	MethodGetRankingByPrincipalID = 0x6

	// MethodGetCategorySetting is the method ID for the method GetCategorySetting
	MethodGetCategorySetting = 0x7

	// MethodGetRankingChart is the method ID for the method GetRankingChart
	MethodGetRankingChart = 0x8

	// MethodGetRankingCharts is the method ID for the method GetRankingCharts
	MethodGetRankingCharts = 0x9

	// MethodGetEstimateScoreRank is the method ID for the method GetEstimateScoreRank
	MethodGetEstimateScoreRank = 0xA
)

// Protocol stores all the RMC method handlers for the Ranking2 protocol and listens for requests
type Protocol struct {
	Server                         *nex.Server
	putScoreHandler                func(err error, client *nex.Client, callID uint32, scoreDataList []*ranking2_types.Ranking2ScoreData, nexUniqueID uint64)
	getCommonDataHandler           func(err error, client *nex.Client, callID uint32, optionFlags uint32, principalID uint32, nexUniqueID uint64)
	putCommonDataHandler           func(err error, client *nex.Client, callID uint32, commonData *ranking2_types.Ranking2CommonData, nexUniqueID uint64)
	deleteCommonDataHandler        func(err error, client *nex.Client, callID uint32, nexUniqueID uint64)
	getRankingHandler              func(err error, client *nex.Client, callID uint32, getParam *ranking2_types.Ranking2GetParam)
	getRankingByPrincipalIDHandler func(err error, client *nex.Client, callID uint32, getParam *ranking2_types.Ranking2GetParam, principalIDList []uint32)
	getCategorySettingHandler      func(err error, client *nex.Client, callID uint32, category uint32)
	getRankingChartHandler         func(err error, client *nex.Client, callID uint32, info *ranking2_types.Ranking2ChartInfoInput)
	getRankingChartsHandler        func(err error, client *nex.Client, callID uint32, infoArray []*ranking2_types.Ranking2ChartInfoInput)
	getEstimateScoreRankHandler    func(err error, client *nex.Client, callID uint32, input *ranking2_types.Ranking2EstimateScoreRankInput)
}

// Setup initializes the protocol
func (protocol *Protocol) Setup() {
	protocol.Server.On("Data", func(packet nex.PacketInterface) {
		request := packet.RMCRequest()

		if request.ProtocolID() == ProtocolID {
			switch request.MethodID() {
			case MethodPutScore:
				protocol.handlePutScore(packet)
			case MethodGetCommonData:
				protocol.handleGetCommonData(packet)
			case MethodPutCommonData:
				protocol.handlePutCommonData(packet)
			case MethodDeleteCommonData:
				protocol.handleDeleteCommonData(packet)
			case MethodGetRanking:
				protocol.handleGetRanking(packet)
			case MethodGetRankingByPrincipalID:
				protocol.handleGetRankingByPrincipalID(packet)
			case MethodGetCategorySetting:
				protocol.handleGetCategorySetting(packet)
			case MethodGetRankingChart:
				protocol.handleGetRankingChart(packet)
			case MethodGetRankingCharts:
				protocol.handleGetRankingCharts(packet)
			case MethodGetEstimateScoreRank:
				protocol.handleGetEstimateScoreRank(packet)
			default:
				go globals.RespondNotImplemented(packet, ProtocolID)
				fmt.Printf("Unsupported Ranking2 method ID: %#v\n", request.MethodID())
			}
		}
	})
}

// NewProtocol returns a new Ranking2 protocol
func NewProtocol(server *nex.Server) *Protocol {
	protocol := &Protocol{Server: server}

	protocol.Setup()

	return protocol
}
