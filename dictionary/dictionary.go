package main

func main() {}

type Dictionary map[string]string

const (
	ErrorWordNotFound  = ErrDictionary("word not found")
	ErrorWordExists    = ErrDictionary("word already exists")
	ErrorWordNotExists = ErrDictionary("word not exists")
)

type ErrDictionary string

func (e ErrDictionary) Error() string {
	return string(e)
}

func (d Dictionary) Search(word string) (string, error) {
	if d[word] == "" {
		return "", ErrorWordNotFound
	}
	return d[word], nil
}

func (d Dictionary) Add(word string, definition string) error {
	_, err := d.Search(word)
	switch err {
	case ErrorWordNotFound:
		d[word] = definition
	case nil:
		return ErrorWordExists
	default:
		return err
	}

	return nil
}

func (d Dictionary) Update(word, newDefinition string) error {
	_, err := d.Search(word)
	switch err {
	case ErrorWordNotFound:
		return ErrorWordNotExists
	case nil:
		d[word] = newDefinition
	default:
		return err
	}

	return nil
}

func (d Dictionary) Delete(word string) {
	delete(d, word)
}
