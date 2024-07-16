package entities

type AccountType int

const (
	AccountTypeSavings AccountType = iota + 1
	AccountTypeChecking 
	AccountTypeBusiness
	AccountTypeStudent
	AccountTypeJoint
)