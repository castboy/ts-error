package args

import "ts-error/err_code"

type ArgsType err_code.ErrCodeSub

const (
	Trade ArgsType = iota + 1
	// etc
	CreateAccount
	UpdateAccount
	UpdateSource
	AddSymbol
	UpdateSymbol
	DeleteSymbol
	AddHoliday
	UpdateHoliday
	DeleteHoliday
	AddGroup
	UpdateGroup
	DeleteGroup
	UpdateSecurity
	DSTSwitch
)

var ArgsTypeMsg = map[ArgsType]string{
	Trade:          "Trade",
	CreateAccount:  "CreateAccount",
	UpdateAccount:  "UpdateAccount",
	UpdateSource:   "UpdateSource",
	AddSymbol:      "AddSymbol",
	UpdateSymbol:   "UpdateSymbol",
	DeleteSymbol:   "DeleteSymbol",
	AddHoliday:     "AddHoliday",
	UpdateHoliday:  "UpdateHoliday",
	DeleteHoliday:  "DeleteHoliday",
	AddGroup:       "AddGroup",
	UpdateGroup:    "UpdateGroup",
	DeleteGroup:    "DeleteGroup",
	UpdateSecurity: "UpdateSecurity",
	DSTSwitch:      "DSTSwitch",
}
