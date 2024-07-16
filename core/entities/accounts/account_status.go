package entities

type AccountStatus int

const (
	AccountStatusActive AccountStatus = iota + 1
	AccountStatusInactive
    AccountStatusDormant
    AccountStatusClosed
    AccountStatusSuspended
    AccountStatusPending
    AccountStatusOverdrawn
    AccountStatusFrozen
    AccountStatusRestricted
    AccountStatusVerified
    AccountStatusUnverified
)