package storage

// Interface for a simple persistent storage for a string value.
type ISimpleStorage interface {
	Put(value string) error
	Get() (string, error)
}
