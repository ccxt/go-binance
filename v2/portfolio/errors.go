package portfolio

import "github.com/adshao/go-binance/v2/common"

// Error represents a portfolio error extending the common APIError
type Error struct {
	common.APIError
}

// NewError creates a new portfolio Error
func NewError(code int64, message string) *Error {
	return &Error{
		APIError: common.APIError{
			Code:    code,
			Message: message,
		},
	}
}

// NewErrorFromResponse creates a new portfolio Error from raw response
func NewErrorFromResponse(code int64, message string, response []byte) *Error {
	return &Error{
		APIError: common.APIError{
			Code:     code,
			Message:  message,
			Response: response,
		},
	}
}

// Error returns the error message
func (e Error) Error() string {
	return e.APIError.Error()
}

// IsPortfolioError check if e is a Portfolio error
func IsPortfolioError(e error) bool {
	_, ok := e.(*Error)
	return ok
}

// 10xx - General Server or Network issues
const (
	ErrUnknown              = -1000 // An unknown error occurred while processing the request
	ErrDisconnected         = -1001 // Internal error; unable to process your request
	ErrUnauthorized         = -1002 // You are not authorized to execute this request
	ErrTooManyRequests      = -1003 // Too many requests
	ErrDuplicateIP          = -1004 // This IP is already on the white list
	ErrNoSuchIP             = -1005 // No such IP has been white listed
	ErrUnexpectedResp       = -1006 // An unexpected response was received from the message bus
	ErrTimeout              = -1007 // Timeout waiting for response from backend server
	ErrErrorMsgReceived     = -1010 // Error message received
	ErrNonWhiteList         = -1011 // This IP cannot access this route
	ErrInvalidMessage       = -1013 // Invalid message
	ErrUnknownOrderCompose  = -1014 // Unsupported order combination
	ErrTooManyOrders        = -1015 // Too many new orders
	ErrServiceShuttingDown  = -1016 // This service is no longer available
	ErrUnsupportedOperation = -1020 // This operation is not supported
	ErrInvalidTimestamp     = -1021 // Timestamp for this request is outside of the recvWindow
	ErrInvalidSignature     = -1022 // Signature for this request is not valid
	ErrStartTimeGreaterEnd  = -1023 // Start time is greater than end time
)

// 11xx - Request issues
const (
	ErrIllegalChars                   = -1100 // Illegal characters found in parameter
	ErrTooManyParameters              = -1101 // Too many parameters sent for this endpoint
	ErrMandatoryParamEmptyOrMalformed = -1102 // Mandatory parameter was not sent, empty/null, or malformed
	ErrUnknownParam                   = -1103 // An unknown parameter was sent
	ErrUnreadParameters               = -1104 // Not all sent parameters were read
	ErrParamEmpty                     = -1105 // Parameter was empty
	ErrParamNotRequired               = -1106 // Parameter was sent when not required
	ErrBadAsset                       = -1108 // Invalid asset
	ErrBadAccount                     = -1109 // Invalid account
	ErrBadInstrumentType              = -1110 // Invalid symbolType
	ErrBadPrecision                   = -1111 // Precision is over the maximum defined
	ErrNoDepth                        = -1112 // No orders on book for symbol
	ErrWithdrawNotNegative            = -1113 // Withdrawal amount must be negative
	ErrTIFNotRequired                 = -1114 // TimeInForce parameter sent when not required
	ErrInvalidTIF                     = -1115 // Invalid timeInForce
	ErrInvalidOrderType               = -1116 // Invalid orderType
	ErrInvalidSide                    = -1117 // Invalid side
	ErrEmptyNewClOrdID                = -1118 // New client order ID was empty
	ErrEmptyOrgClOrdID                = -1119 // Original client order ID was empty
	ErrBadInterval                    = -1120 // Invalid interval
	ErrBadSymbol                      = -1121 // Invalid symbol
	ErrInvalidListenKey               = -1125 // This listenKey does not exist
	ErrMoreThanXXHours                = -1127 // Lookup interval is too big
	ErrOptionalParamsBadCombo         = -1128 // Combination of optional parameters invalid
	ErrInvalidParameter               = -1130 // Invalid data sent for a parameter
	ErrInvalidNewOrderRespType        = -1136 // Invalid newOrderRespType
)

// 20xx - Processing Issues
const (
	ErrNewOrderRejected                = -2010 // NEW_ORDER_REJECTED
	ErrCancelRejected                  = -2011 // CANCEL_REJECTED
	ErrNoSuchOrder                     = -2013 // Order does not exist
	ErrBadAPIKeyFmt                    = -2014 // API-key format invalid
	ErrRejectedMBXKey                  = -2015 // Invalid API-key, IP, or permissions
	ErrNoTradingWindow                 = -2016 // No trading window could be found
	ErrBalanceNotSufficient            = -2018 // Balance is insufficient
	ErrMarginNotSufficient             = -2019 // Margin is insufficient
	ErrUnableToFill                    = -2020 // Unable to fill
	ErrOrderWouldImmediatelyTrigger    = -2021 // Order would immediately trigger
	ErrReduceOnlyReject                = -2022 // ReduceOnly Order is rejected
	ErrUserInLiquidation               = -2023 // User in liquidation mode now
	ErrPositionNotSufficient           = -2024 // Position is not sufficient
	ErrMaxOpenOrderExceeded            = -2025 // Max open order exceeded
	ErrReduceOnlyOrderTypeNotSupported = -2026 // Reduce only order type not supported
	ErrMaxLeverageRatio                = -2027 // Max leverage ratio reached
	ErrMinLeverageRatio                = -2028 // Min leverage ratio reached
)

// 40xx - Filters and Other Issues
const (
	ErrInvalidOrderStatus                 = -4000 // Invalid order status
	ErrPriceLessThanZero                  = -4001 // Price less than zero
	ErrPriceGreaterThanMaxPrice           = -4002 // Price greater than max price
	ErrQtyLessThanZero                    = -4003 // Quantity less than zero
	ErrQtyLessThanMinQty                  = -4004 // Quantity less than min quantity
	ErrQtyGreaterThanMaxQty               = -4005 // Quantity greater than max quantity
	ErrStopPriceLessThanZero              = -4006 // Stop price less than zero
	ErrStopPriceGreaterThanMaxPrice       = -4007 // Stop price greater than max price
	ErrTickSizeLessThanZero               = -4008 // Tick size less than zero
	ErrMaxPriceLessThanMinPrice           = -4009 // Max price less than min price
	ErrMaxQtyLessThanMinQty               = -4010 // Max quantity less than min quantity
	ErrStepSizeLessThanZero               = -4011 // Step size less than zero
	ErrMaxNumOrdersLessThanZero           = -4012 // Max number of orders less than zero
	ErrPriceLessThanMinPrice              = -4013 // Price less than min price
	ErrPriceNotIncreasedByTickSize        = -4014 // Price not increased by tick size
	ErrInvalidClOrdIDLen                  = -4015 // Invalid client order ID length
	ErrPriceHigherThanMultiplierUp        = -4016 // Price higher than multiplier up
	ErrMultiplierUpLessThanZero           = -4017 // Multiplier up less than zero
	ErrMultiplierDownLessThanZero         = -4018 // Multiplier down less than zero
	ErrCompositeScaleOverflow             = -4019 // Composite scale overflow
	ErrTargetStrategyInvalid              = -4020 // Target strategy invalid
	ErrInvalidDepthLimit                  = -4021 // Invalid depth limit
	ErrWrongMarketStatus                  = -4022 // Wrong market status
	ErrQtyNotIncreasedByStepSize          = -4023 // Quantity not increased by step size
	ErrPriceLowerThanMultiplierDown       = -4024 // Price lower than multiplier down
	ErrMultiplierDecimalLessThanZero      = -4025 // Multiplier decimal less than zero
	ErrCommissionInvalid                  = -4026 // Commission invalid
	ErrInvalidAccountType                 = -4027 // Invalid account type
	ErrInvalidLeverage                    = -4028 // Invalid leverage
	ErrInvalidTickSizePrecision           = -4029 // Invalid tick size precision
	ErrInvalidStepSizePrecision           = -4030 // Invalid step size precision
	ErrInvalidWorkingType                 = -4031 // Invalid working type
	ErrExceedMaxCancelOrderSize           = -4032 // Exceed max cancel order size
	ErrInsuranceAccountNotFound           = -4033 // Insurance account not found
	ErrInvalidBalanceType                 = -4044 // Invalid balance type
	ErrMaxStopOrderExceeded               = -4045 // Max stop order exceeded
	ErrNoNeedToChangeMarginType           = -4046 // No need to change margin type
	ErrThereExistsOpenOrders              = -4047 // There exists open orders
	ErrThereExistsQuantity                = -4048 // There exists quantity
	ErrAddIsolatedMarginReject            = -4049 // Add isolated margin reject
	ErrCrossBalanceInsufficient           = -4050 // Cross balance insufficient
	ErrIsolatedBalanceInsufficient        = -4051 // Isolated balance insufficient
	ErrNoNeedToChangeAutoAddMargin        = -4052 // No need to change auto add margin
	ErrAutoAddCrossedMarginReject         = -4053 // Auto add crossed margin reject
	ErrAddIsolatedMarginNoPositionReject  = -4054 // Add isolated margin no position reject
	ErrAmountMustBePositive               = -4055 // Amount must be positive
	ErrInvalidAPIKeyType                  = -4056 // Invalid API key type
	ErrInvalidRSAPublicKey                = -4057 // Invalid RSA public key
	ErrMaxPriceTooLarge                   = -4058 // Max price too large
	ErrNoNeedToChangePositionSide         = -4059 // No need to change position side
	ErrInvalidPositionSide                = -4060 // Invalid position side
	ErrPositionSideNotMatch               = -4061 // Position side not match
	ErrReduceOnlyConflict                 = -4062 // Reduce only conflict
	ErrInvalidOptionsRequestType          = -4063 // Invalid options request type
	ErrInvalidOptionsTimeFrame            = -4064 // Invalid options time frame
	ErrInvalidOptionsAmount               = -4065 // Invalid options amount
	ErrInvalidOptionsEventType            = -4066 // Invalid options event type
	ErrPositionSideChangeExistsOpenOrders = -4067 // Position side change exists open orders
	ErrPositionSideChangeExistsQuantity   = -4068 // Position side change exists quantity
	ErrInvalidOptionsPremiumFee           = -4069 // Invalid options premium fee
	ErrInvalidClOptionsIDLen              = -4070 // Invalid cl options ID length
	ErrInvalidOptionsDirection            = -4071 // Invalid options direction
	ErrOptionsPremiumNotUpdate            = -4072 // Options premium not update
	ErrOptionsPremiumInputLessThanZero    = -4073 // Options premium input less than zero
	ErrOptionsAmountBiggerThanUpper       = -4074 // Options amount bigger than upper
	ErrOptionsPremiumOutputZero           = -4075 // Options premium output zero
	ErrOptionsPremiumTooDiff              = -4076 // Options premium too diff
	ErrOptionsPremiumReachLimit           = -4077 // Options premium reach limit
	ErrOptionsCommonError                 = -4078 // Options common error
	ErrInvalidOptionsID                   = -4079 // Invalid options ID
	ErrOptionsUserNotFound                = -4080 // Options user not found
	ErrOptionsNotFound                    = -4081 // Options not found
	ErrInvalidBatchPlaceOrderSize         = -4082 // Invalid batch place order size
	ErrPlaceBatchOrdersFail               = -4083 // Place batch orders fail
	ErrUpcomingMethod                     = -4084 // Upcoming method
	ErrInvalidNotionalLimitCoef           = -4085 // Invalid notional limit coefficient
	ErrInvalidPriceSpreadThreshold        = -4086 // Invalid price spread threshold
	ErrReduceOnlyOrderPermission          = -4087 // Reduce only order permission
	ErrNoPlaceOrderPermission             = -4088 // No place order permission
	ErrInvalidContractType                = -4104 // Invalid contract type
	ErrInvalidClientTranIDLen             = -4114 // Invalid client transaction ID length
	ErrDuplicatedClientTranID             = -4115 // Duplicated client transaction ID
	ErrReduceOnlyMarginCheckFailed        = -4118 // Reduce only margin check failed
	ErrMarketOrderReject                  = -4131 // Market order reject
	ErrInvalidActivationPrice             = -4135 // Invalid activation price
	ErrQuantityExistsWithClosePosition    = -4137 // Quantity exists with close position
	ErrReduceOnlyMustBeTrue               = -4138 // Reduce only must be true
	ErrOrderTypeCannotBeMKT               = -4139 // Order type cannot be MKT
	ErrInvalidOpeningPositionStatus       = -4140 // Invalid opening position status
	ErrSymbolAlreadyClosed                = -4141 // Symbol already closed
	ErrStrategyInvalidTriggerPrice        = -4142 // Strategy invalid trigger price
	ErrInvalidPair                        = -4144 // Invalid pair
	ErrIsolatedLeverageRejectWithPosition = -4161 // Isolated leverage reject with position
	ErrMinNotional                        = -4164 // Min notional
	ErrInvalidTimeInterval                = -4165 // Invalid time interval
	ErrPriceHigherThanStopMultiplierUp    = -4183 // Price higher than stop multiplier up
	ErrPriceLowerThanStopMultiplierDown   = -4184 // Price lower than stop multiplier down
)

// 50xx - Order Execution Issues
const (
	ErrFOKOrderReject      = -5021 // FOK order rejected
	ErrGTXOrderReject      = -5022 // GTX order rejected
	ErrMERecvWindowReject  = -5028 // ME recvWindow rejected
	ErrTooManyRequestQueue = -5041 // Too many requests in queue
)
