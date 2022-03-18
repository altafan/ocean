package ports

import (
	"github.com/vulpemventures/ocean/internal/core/domain"
)

// BlockchainScanner is the abstraction for any kind of service representing an
// Elements node. It gives info about txs and utxos related to one or more HD
// accounts in a aync way (via channels), and lets broadcast transactions over
// the Liquid network.
type BlockchainScanner interface {
	// Start starts the service.
	Start()
	// Stop stops the service.
	Stop()

	// WatchForAccount instructs the scanner to start notifying about txs/utxos
	// related to the given list of addresses belonging to the given HD account.
	WatchForAccount(accountName string, addresses []domain.AddressInfo)
	// StopWatchForAccount instructs the scanner to stop notifying about
	// txs/utxos related to any address belonging to the given HD account.
	StopWatchForAccount(accountName string)

	// GetUtxoChannel returns the channel where notification about utxos realated
	// to the given HD account are sent.
	GetUtxoChannel(accountName string) chan []*domain.Utxo
	// GetTxChannel returns the channel where notification about txs realated to
	// the given HD account are sent.
	GetTxChannel(accountName string) chan *domain.Transaction

	// GetUtxos is a sync function to get info about the utxos represented by
	// given outpoints (UtxoKeys).
	GetUtxos(utxoKeys []domain.UtxoKey) ([]*domain.Utxo, error)
	// BroadcastTransaction sends the given raw tx (in hex string) over the
	// network in order to be included in a later block of the Liquid blockchain.
	BroadcastTransaction(txHex string) (string, error)
}