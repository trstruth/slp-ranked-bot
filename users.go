package main

type UserGetter interface {
	getUserList() []string
}

type FileBasedUserGetter struct {
	path string
}
