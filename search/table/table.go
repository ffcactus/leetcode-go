package table

// Key represents key used in Table.
type Key int

// Value represents the value used in Table.
type Value interface{}

// Table defines the common interface that a symbol-table should provide.
type Table interface {
	// Put key-value pair into the table, remove key from table if value is nil.
	Put(k Key, v Value)

	// Get the value paired with key, nil if the key is absent.
	Get(k Key) Value

	// Delete key and it's value from the table.
	Delete (k Key)

	// Contains check if there is a pair with the key in the table.
	Contains(k Key) bool

	// IsEmpty check if the table is empty.
	IsEmpty() bool

	// Size return the count of pairs in the table.
	Size() int

	// Min return the smallest key.
	Min() Key

	// Max return the largest key.
	Max() Key

	// Floor return the largest key that less or equal to the key.
	Floor(k Key) Key

	// Ceiling return the smallest key that greater or equal to the key.
	Ceiling(k Key) Key

	// Rank return the number of keys that is less than the key.
	// i == rank(select(i)) for all i between 0 and size()-1
	Rank(k Key) int

	// Select return the key that is of rank k.
	// All keys in the table satisfy key == select(rank(key))
	Select(k int) Key

	// DeleteMin delete the smallest key.
	DeleteMin()

	// DeleteMax delete the largest key.
	DeleteMax()

	// SizeBetween return the number of keys in [low, high]
	SizeBetween(low, high Key) int
}
