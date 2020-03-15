// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "time"

// GeneralLedgerEntry undocumented
type GeneralLedgerEntry struct {
	// Entity is the base model of GeneralLedgerEntry
	Entity
	// PostingDate undocumented
	PostingDate *Date `json:"postingDate,omitempty"`
	// DocumentNumber undocumented
	DocumentNumber *string `json:"documentNumber,omitempty"`
	// DocumentType undocumented
	DocumentType *string `json:"documentType,omitempty"`
	// AccountID undocumented
	AccountID *UUID `json:"accountId,omitempty"`
	// AccountNumber undocumented
	AccountNumber *string `json:"accountNumber,omitempty"`
	// Description undocumented
	Description *string `json:"description,omitempty"`
	// DebitAmount undocumented
	DebitAmount *int `json:"debitAmount,omitempty"`
	// CreditAmount undocumented
	CreditAmount *int `json:"creditAmount,omitempty"`
	// LastModifiedDateTime undocumented
	LastModifiedDateTime *time.Time `json:"lastModifiedDateTime,omitempty"`
	// Account undocumented
	Account *Account `json:"account,omitempty"`
}
