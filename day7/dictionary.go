package main

import "errors"

type Dictionary map[string]string

var (
	ErrNotFound         = errors.New("could not find the word you were looking for")
	ErrWordExist        = errors.New("cannot add a word that already exists")
	ErrWordDoesNotExist = errors.New("cannot update word because it does not exist")
)

func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]

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
		return ErrWordExist
	default:
		return err
	}

	return nil
}

func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExist
	case nil:
		d[word] = definition
	default:
		return err
	}

	return nil
}

func (d Dictionary) Delete(word string) {
	delete(d, word)

	// https: //dave.cheney.net/2014/12/01/five-suggestions-for-setting-up-a-go-project
}
