package storage

import (
	"fmt"
	"strconv"

	badger "github.com/dgraph-io/badger/v4"
	"github.com/feedloop/syde/internal/model"
)

// Counter keys in BadgerDB have the form "c:<kind>" and store the
// HIGHEST ID ever issued for that kind as an ASCII decimal string.
// NextID increments atomically inside a txn.

func counterKey(kind model.EntityKind) []byte {
	return []byte("c:" + string(kind))
}

// readCounter returns the stored counter (0 if absent).
func (idx *Index) readCounter(txn *badger.Txn, kind model.EntityKind) (int, error) {
	item, err := txn.Get(counterKey(kind))
	if err == badger.ErrKeyNotFound {
		return 0, nil
	}
	if err != nil {
		return 0, err
	}
	var n int
	err = item.Value(func(v []byte) error {
		parsed, perr := strconv.Atoi(string(v))
		if perr != nil {
			return perr
		}
		n = parsed
		return nil
	})
	return n, err
}

// GetCounter returns the current (max-issued) counter for a kind.
func (idx *Index) GetCounter(kind model.EntityKind) (int, error) {
	var n int
	err := idx.db.View(func(txn *badger.Txn) error {
		v, err := idx.readCounter(txn, kind)
		if err != nil {
			return err
		}
		n = v
		return nil
	})
	return n, err
}

// IncrementCounter atomically bumps the kind's counter and returns the
// new value. This is the only way new entity IDs are allocated.
func (idx *Index) IncrementCounter(kind model.EntityKind) (int, error) {
	var next int
	err := idx.db.Update(func(txn *badger.Txn) error {
		current, err := idx.readCounter(txn, kind)
		if err != nil {
			return err
		}
		next = current + 1
		return txn.Set(counterKey(kind), []byte(strconv.Itoa(next)))
	})
	return next, err
}

// SetCounter forces the counter to the given value. Used by Reindex to
// restore counters from scanned entity IDs, and when importing data.
func (idx *Index) SetCounter(kind model.EntityKind, n int) error {
	return idx.db.Update(func(txn *badger.Txn) error {
		return txn.Set(counterKey(kind), []byte(strconv.Itoa(n)))
	})
}

// NextID allocates the next counter value for a given kind using the
// already-open index and returns an ID of the form "PFX-0001".
// Callers must pass the live *Index (not open a new one) because
// BadgerDB holds an exclusive directory lock.
func NextID(idx *Index, kind model.EntityKind) (string, error) {
	n, err := idx.IncrementCounter(kind)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s-%04d", kind.IDPrefix(), n), nil
}
