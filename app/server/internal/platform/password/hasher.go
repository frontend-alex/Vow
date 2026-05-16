package password

type Hasher struct{}

func NewHasher() *Hasher {
	return &Hasher{}
}

func (h *Hasher) Hash(plain string) (string, error) {
	return "", nil
}

func (h *Hasher) Compare(hash string, plain string) bool {
	return false
}
