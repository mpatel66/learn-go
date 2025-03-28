package dictionary

type Dictionary map[string]string

type DictionaryError string

const (
	ErrNotFound   = DictionaryError("could not find the word you were looking for")
	ErrWordExists = DictionaryError("this word already exists in the dictionary")
)

func (e DictionaryError) Error() string {
	return string(e)
}

func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word] // map can return two values, one of which is a boolean that indicates if it found the key

	if !ok {

		return "", ErrNotFound
	}
	return definition, nil
}

func (d Dictionary) Add(word, definition string) error {

	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		d[word] = definition
	case nil:
		return ErrWordExists
	default:
		return err
	}

	return nil

}

func (d Dictionary) Update(word, definition string) {
	d[word] = definition
}
