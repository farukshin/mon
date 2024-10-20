package main

type notification struct {
	Uid    string     `json:"uid"`
	Type   string     `json:"type"`
	Name   string     `json:"name"`
	Params []keyValue `json:"params"`
}

type keyValue struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func (ntf *notification) send(mes string) error {
	return nil
}
