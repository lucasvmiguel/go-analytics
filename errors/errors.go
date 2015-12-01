package errors

import "errors"

var (
	ErrInvalidConnectionDatabase = errors.New("Invalid connection with database")
	ErrNewNotification           = errors.New("Invalid new notification")
	ErrSerializeNotification     = errors.New("Serialize notification")
	ErrFindCompany               = errors.New("Can't find company")
	ErrUpgradeWebsocket          = errors.New("Can't upgrade websocket")
	ErrAuthWebsocket             = errors.New("Auth websocket")
)
