package table

// node represents the node in the BST tree.
type node struct {
	key Key
	value Value
	n int // the number of node that includes subtree of left and right and this root.
	left *node
	right *node
}

// bst implements Table interface by binary search tree.
type bst struct {
	root *node
}

// NewBST create a binary search implementation of Table interface.
func NewBST() Table {
	return &bst{
		root: nil,
	}
}

// Put key-value pair into the table, remove key from table if value is nil.
func (impl *bst) Put(k Key, v Value) {
	impl.root = impl.put(impl.root, k, v)
}

func (impl *bst) put(n *node, k Key, v Value) *node {
		if n == nil {
			return &node{
				key: k,
				value: v,
				n: 1,
			}
		}
		if k < n.key {
			n.left = impl.put(n.left, k, v)
		} else if k > n.key {
			n.right = impl.put(n.right, k, v)
		} else {
			n.value = v
		}
		n.n = impl.size(n.left) + impl.size(n.right) + 1
		return n
}

// Get the value paired with key, nil if the key is absent.
func (impl *bst) Get(k Key) Value {
	return impl.get(impl.root, k)
}

func (impl *bst) get(n *node, k Key) Value {
	if n == nil {
		return nil
	}
	switch {
	case k == n.key:
		return n.value
	case k < n.key:
		return impl.get(n.left, k)
	case k > n.key:
		return impl.get(n.right, k)
	}
	return nil
}

// Delete key and it's value from the table.
func (impl *bst) Delete (k Key) {

}

// Contains check if there is a pair with the key in the table.
func (impl *bst) Contains(k Key) bool {

}

// IsEmpty check if the table is empty.
func (impl *bst) IsEmpty() bool {
	return impl.Size() == 0
}

// Size return the count of pairs in the table.
func (impl *bst) Size() int {
	return impl.size(impl.root)
}

func (impl *bst) size(n *node) int {
	if n == nil {
		return 0
	}
	return n.n
}

// Min return the smallest key.
func (impl *bst) Min() Key {
	return impl.min(impl.root).key
}

func (impl *bst) min(n *node) *node {
	if n.left == nil {
		return n
	}
	return impl.min(n.left)
}

// Max return the largest key.
func (impl *bst) Max() Key {
	return impl.max(impl.root).key
}

func (impl *bst) max(n *node) *node {
	if n.right == nil {
		return n
	}
	return impl.max(n.right)
}

// Floor return the largest key that less or equal to the key.
func (impl *bst) Floor(k Key) Key {
	return impl.floor(impl.root, k).key
}

func (impl *bst) floor(n *node, k Key) *node {
	if n == nil {
		return nil
	}
	if k == n.key {
		return n
	}
	if k < n.key {
		return impl.floor(n.left, k)
	}
	t := impl.floor(n.right, k)
	if t == nil {
		return t
	}
	return n
}
// Ceiling return the smallest key that greater or equal to the key.
func (impl *bst) Ceiling(k Key) Key {

}

// Rank return the number of keys that is less than the key.
// i == rank(select(i)) for all i between 0 and size()-1
func (impl *bst) Rank(k Key) int {
	return impl.rank(impl.root, k)
}

func (impl *bst) rank(n *node, k Key) int {
	if n == nil {
		return 0
	}
	if k < n.key {
		return impl.rank(n.left, k)
	}
	if k > n.key {
		return 1 + impl.size(n.left) + impl.rank(n.right, k)
	}
	return impl.size(n.left)
}

// Select return the key that is of rank k.
// All keys in the table satisfy key == select(rank(key))
func (impl *bst) Select(k int) Key {
	return impl.internalSelect(impl.root, k).key
}

func (impl *bst) internalSelect(n *node, k int) *node {
	if n == nil {
		return nil
	}
	t := impl.size(n.left)
	if t > k {
		return impl.internalSelect(n.left, k)
	} else if k < k {
		return impl.internalSelect(n.right, k - t - 1)
	}
	return n
}

// DeleteMin delete the smallest key.
func (impl *bst) DeleteMin() {
	impl.Delete(impl.Min())
}

// DeleteMax delete the largest key.
func (impl *bst) DeleteMax() {
	impl.Delete(impl.Max())
}

// SizeBetween return the number of keys in [low, high]
func (impl *bst) SizeBetween(low, high Key) int {
	if high < low {
		return 0
	} else if impl.Contains(high) {
		return impl.Rank(high) - impl.Rank(low) + 1
	} else {
		return impl.Rank(high) - impl.Rank(low)
	}
}


