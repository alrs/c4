package c4

// An asset is data identified by it's C4 ID.
type Asset interface {
	ID() *ID
}
