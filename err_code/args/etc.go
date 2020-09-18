package args

import "ts-error/err_code"

type AccountArg err_code.ErrCodeSubSub

const (
	AccountNameErr AccountArg = iota + 1
	AccountPhoneErr
	AccountGroupIDErr
	AccountLoginErr
	AccountPasswordErr
)

var AccountArgsMsg = map[AccountArg]string{
	AccountNameErr: "name",
	AccountPhoneErr: "phone",
	AccountGroupIDErr: "group",
	AccountLoginErr: "login",
	AccountPasswordErr: "password",
}

type SymbolArg err_code.ErrCodeSubSub

const (
	SymbolName SymbolArg = iota + 1
	SymbolSource
	SymbolSecurityID
	SymbolPercentage
	SymbolMarginInitial
)

var SymbolArgsMsg = map[SymbolArg]string{
	SymbolName: "name",
	SymbolSource: "source",
	SymbolSecurityID: "securityID",
	SymbolPercentage: "percentage",
	SymbolMarginInitial: "marginInitial",
}

type SourceArg err_code.ErrCodeSubSub

const (
	SourceName SourceArg = iota + 1
	SourceType
	SourceSwapType
	SourceMarginMode
	SourceCurrency
	SourceMarginCurrency
	SourceProfitCurrency
	SourceContractSize
	SourceDigits
	SourceStopsLevel
	SourceSwap3Day
	SourceSwapLong
	SourceSwapShort
	SourceMarketOwnerType
)

var SourceArgsMsg = map[SourceArg]string{
	SourceName: "name",
	SourceType: "type",
	SourceSwapType: "swapType",
	SourceMarginMode: "marginMode",
	SourceCurrency: "currency",
	SourceMarginCurrency: "marginCurrency",
	SourceProfitCurrency: "profitCurrency",
	SourceContractSize: "contractSize",
	SourceDigits: "digits",
	SourceStopsLevel: "stopsLevel",
	SourceSwap3Day: "swap3Day",
	SourceSwapLong: "swapLong",
	SourceSwapShort: "swapShort",
	SourceMarketOwnerType: "marketOwnerType",
}

type SessionArg err_code.ErrCodeSubSub

const (
	SessionID  SessionArg = iota + 1
	SessionSourceID
	SessionType
	SessionInfo
)

var SessionArgsMsg = map[SessionArg]string{
	SessionID: "id",
	SessionSourceID: "sourceID",
	SessionType: "type",
	SessionInfo: "info",
}

type HolidayArg err_code.ErrCodeSubSub

const (
	HolidayEnable HolidayArg = iota + 1
	HolidayDate
	HolidayFrom
	HolidayTo
	HolidayCatagory
	HolidaySymbol
	HolidayDescription
)

var HolidayArgsMsg = map[HolidayArg]string{
	HolidayEnable: "enable",
	HolidayDate: "date",
	HolidayFrom: "from",
	HolidayTo: "to",
	HolidayCatagory: "category",
	HolidaySymbol: "symbol",
	HolidayDescription: "description",
}

type GroupArg err_code.ErrCodeSubSub

const (
	GroupId GroupArg = iota + 1
	GroupName
	GroupTradeType
	GroupMarginStopOut
	GroupMarginCall
	GroupMarginMode
	GroupDepositCurrency
	GroupIsChargeSwap
	GroupComment
)

var GroupArgsMsg = map[GroupArg]string{
	GroupId: "id",
	GroupName: "name",
	GroupTradeType: "tradeType",
	GroupMarginStopOut: "marginStopOut",
	GroupMarginCall: "marginCall",
	GroupMarginMode: "marginMode",
	GroupDepositCurrency: "depositCurrency",
	GroupIsChargeSwap: "isChargeSwap",
	GroupComment: "comment",
}

type SecurityArg err_code.ErrCodeSubSub

const (
	SecurityGroupID SecurityArg = iota + 1
	SecuritySecurityID
	SecurityStatus
	SecuritySpreadDiff
	SecurityLotMin
	SecurityLotMax
	SecurityLotStep
	SecurityCommission
)

var SecurityArgsMsg = map[SecurityArg]string{
	SecurityGroupID:    "groupID",
	SecuritySecurityID: "securityID",
	SecurityStatus:     "status",
	SecuritySpreadDiff: "spreadDiff",
	SecurityLotMin:     "lotMin",
	SecurityLotMax:     "lotMax",
	SecurityLotStep:    "lotStep",
	SecurityCommission: "commission",
}

type DSTArg err_code.ErrCodeSubSub

const (
	DSTMarketOwnerType DSTArg = iota + 1
	DSTType
)

var DSTArgsMsg = map[DSTArg]string{
	DSTMarketOwnerType: "marketOwnerType",
	DSTType: "dstType",
}




