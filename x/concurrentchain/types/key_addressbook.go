package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// AddressbookKeyPrefix is the prefix to retrieve all Addressbook
	AddressbookKeyPrefix = "Addressbook/value/"
)

// AddressbookKey returns the store key to retrieve a Addressbook from the index fields
func AddressbookKey(
	index string,
) []byte {
	var key []byte

	indexBytes := []byte(index)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}
