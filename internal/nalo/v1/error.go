package nalo

// 1701:Success, Message Submitted Successfully, In this case you will receive
// the response 1701|<CELL_NO>|<MESSAGE ID>, The message Id can
// then be used later to map the delivery reports to this message.
// 1702:Invalid URL Error, This means that one of the parameters was not
// provided or left blank
// 1703:Invalid value in username or password field
// 1704:Invalid value in "type" field
// 1705:Invalid Message
// 1706:Invalid Destination
// 1707:Invalid Source (Sender)
// 1708:Invalid value for "dlr" field
// 1709:User validation failed
// 1710:Internal Error
// 1025:Insufficient Credit User
// 1026:Insufficient Credit Reseller
// var (
// 	ErrInvalidURL error = iota
// 	ErrInvalidValueUsername
// 	ErrInvalidType
// 	ErrInvalidMessage
// 	ErrInvalidDestination
// 	ErrInvalidSource
// 	ErrInvalidDLR
// 	ErrInvalidUser
// 	ErrInternal
// 	ErrInsufficientCreditUser
// 	ErrInsufficientCreditReseller
// )
