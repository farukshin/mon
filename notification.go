package main

type notification struct {
	uid   string
	typ   string
	name  string
	param []keyValue
}

type keyValue struct {
	key   string
	value string
}

func (ntf *notification) send(mes string) error {
	return nil
}
