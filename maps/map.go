package maps

type Dictionary map[string]string

type DictErr string

func (e DictErr) Error() string {
	return string(e)
}

var (
	ErrKeyNotFound      = DictErr("key not found")
	ErrKeyAlreadyExists = DictErr("key already exists")
	ErrKeyNotExists     = DictErr("key not exists")
)

func (d Dictionary) Search(key string) (string, error) {
	v, e := d[key]
	if !e {
		return "", ErrKeyNotFound
	}
	return v, nil
}
func (d Dictionary) Add(key, value string) error {
	_, err := d.Search(key)

	switch err {
	case ErrKeyNotFound:
		d[key] = value
	case nil:
		return ErrKeyAlreadyExists
	default:
		return err
	}
	return nil
}
func (d Dictionary) Update(key, value string) error {
	_, err := d.Search(key)
	switch err {
	case ErrKeyNotFound:
		return ErrKeyNotExists
	case nil:
		d[key] = value
	default:
		return err
	}

	return nil
}
func (d Dictionary) Delete(key string) error {
	_, err := d.Search(key)
	switch err {
	case ErrKeyNotFound:
		return ErrKeyNotExists
	case nil:
		delete(d, key)
	default:
		return err
	}
	return nil
}
