package enum

type HashType struct {
	value string
}

var MD5 = HashType{"{MD5}"}

func (hashType *HashType) String() string {
	return hashType.value
}
