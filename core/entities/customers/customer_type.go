package entities

type CustomerType int

const (
	CustomerTypeIndividual CustomerType = iota + 1
	CustomerTypeBusiness
	CustomerTypeVIP
	CustomerTypeNonProfit
	CustomerGovernment
)
