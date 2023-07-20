package studyinterface01

type ByteCounter int

func (c *ByteCounter) Writer(p []byte) (int, error) {
	*c += ByteCounter(len(p))
	return len(p), nil
}
