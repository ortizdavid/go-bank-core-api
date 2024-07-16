package entities

type TransactionStatus int

const (
    TransactionStatusPending TransactionStatus = iota + 1
    TransactionStatusCompleted
    TransactionStatusApproved
    TransactionStatusFailed
    TransactionStatusPendingApproval
    TransactionStatusOnHold
    TransactionStatusExpired
    TransactionStatusCancelled
    TransactionStatusRefunded
)