package entities

type CustomerStatus int

const (
	CustomerStatusActive CustomerStatus = iota + 1
	CustomerStatusInactive
	CustomerStatusSuspended
	CustomerStatusClosed
	CustomerStatusPending	
)
