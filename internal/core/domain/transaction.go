package domain

// Transaction is the data structure representing an Elements tx with extra
// info like whether it is conifirmed/unconfirmed and the name of the accounts
// owning one or more of its inputs.
type Transaction struct {
	TxID        string
	TxHex       string
	BlockHash   string
	BlockHeight uint32
	Accounts    map[string]struct{}
}

// IsConfirmed returns whther the tx is included in the blockchain.
func (t *Transaction) IsConfirmed() bool {
	return t.BlockHash != ""
}

// Confirm marks the tx as confirmed.
func (t *Transaction) Confirm(blockHash string, blockHeight uint32) {
	if t.IsConfirmed() {
		return
	}

	t.BlockHash = blockHash
	t.BlockHeight = blockHeight
}

// AddAccount adds the given account to the map of those involved in the tx.
func (t *Transaction) AddAccount(accountName string) {
	if t.Accounts == nil {
		t.Accounts = make(map[string]struct{})
	}
	t.Accounts[accountName] = struct{}{}
}

// GetAccounts returns the account map as a slice of account names.
func (t *Transaction) GetAccounts() []string {
	accounts := make([]string, 0, len(t.Accounts))
	for account := range t.Accounts {
		accounts = append(accounts, account)
	}
	return accounts
}

// EqualAccounts compares the account map of the current and the given txs,
// returning whether they contain the same account names.
func (t *Transaction) EqualAccounts(tx *Transaction) bool {
	for account := range tx.Accounts {
		if _, ok := t.Accounts[account]; ok {
			return false
		}
	}
	return true
}
