// Package collection shows a example that use go's standard library sort.Interface
package collection

// StringCollection represents the class that implement sort.Interface.
type StringCollection struct {
	raw []string
}

func (impl StringCollection) Len() int {
	return len(impl.raw)
}

func (impl StringCollection) Less(i, j int) bool {
	return impl.raw[i] < impl.raw[j]
}

func (impl StringCollection) Swap(i, j int) {
	impl.raw[i], impl.raw[j] = impl.raw[j], impl.raw[i]
}
