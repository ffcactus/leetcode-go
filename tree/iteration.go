package tree

// Iterator defines the interface to iterate a tree.
type Iterator interface {
	SetCallback(cb Callback)
}

// Callback is the interface when you need receive tree node on iteration.
type Callback interface {
	Process(node *TreeNode)
}

// NewDFSIterator creates a DFS tree iterator.
func DFS(root *TreeNode, preCB Callback) {
	if root == nil {
		return
	}
	if preCB != nil {
		preCB.Process(root)
	}
	if root.Left != nil {
		DFS(root.Left, preCB)
	}
	if root.Right != nil {
		DFS(root.Right, preCB)
	}
}

func dfs(root *TreeNode, cb Callback) {

}
