package main

import (
	"fmt"
	"os"
	"strings"
)

type UserManager interface {
	AddUser(string) error
	GetUserList() ([]string, error)
}

type FileBasedUserManager struct {
	path string
}

func NewFileBasedUserManager(path string) (*FileBasedUserManager, error) {
	// create the db file if it does not already exist
	if _, err := os.Stat(path); os.IsNotExist(err) {
		_, err := os.Create(path)
		if err != nil {
			return nil, err
		}
	}

	return &FileBasedUserManager{
		path: path,
	}, nil
}

func (fb *FileBasedUserManager) GetUserList() ([]string, error) {
	data, err := os.ReadFile(fb.path)
	if err != nil {
		return []string{}, err
	}

	userList := []string{}

	for _, connectCode := range strings.Split(string(data), "\n") {
		if connectCodeIsValid(connectCode) {
			userList = append(userList, connectCode)
		}
	}

	return userList, nil
}

// AddUser adds a user to the file.
// This method will return an error if the user already exists
// TODO: should probably sanitize inputs ? lol ??
func (fb *FileBasedUserManager) AddUser(username string) error {
	// return an error if the user already exists
	userExists, err := fb.contains(username)
	if err != nil {
		return err
	}
	if userExists {
		return fmt.Errorf("db already contains user with id: %s", username)
	}

	f, err := os.OpenFile(fb.path, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	if _, err := f.WriteString(fmt.Sprintf("%s\n", username)); err != nil {
		return err
	}

	return nil
}

func (fb *FileBasedUserManager) contains(username string) (bool, error) {
	userList, err := fb.GetUserList()
	if err != nil {
		return false, err
	}

	for _, userId := range userList {
		if userId == username {
			return true, nil
		}
	}

	return false, nil
}
