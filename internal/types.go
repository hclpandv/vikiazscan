package internal

type OrphanedResource struct {
	Name          string
	ResourceGroup string
	Type          string
	Location      string
	SKUName       string // <--- Make sure this exists
	DiskSize      int
	Tags          string
}
