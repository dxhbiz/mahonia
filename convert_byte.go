package mahonia

// ConvertString converts a  string from UTF-8 to e's encoding.
func (e Encoder) ConvertBytes(s string) []byte {
	dest := make([]byte, len(s)+10)
	destPos := 0

	for _, rune := range s {
	retry:
		size, status := e(dest[destPos:], rune)

		if status == NO_ROOM {
			newDest := make([]byte, len(dest)*2)
			copy(newDest, dest)
			dest = newDest
			goto retry
		}

		if status == STATE_ONLY {
			destPos += size
			goto retry
		}

		destPos += size
	}

	return dest[:destPos]
}
