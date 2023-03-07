package main

type userCmd struct {
	Add    addUserCmd
	Get    getUserCmd
	Update updateUserCmd
	Delete deleteUserCmd
	List   listUserCmd
}

type addUserCmd struct{}
type getUserCmd struct{}
type updateUserCmd struct{}
type deleteUserCmd struct{}
type listUserCmd struct{}

func (a *addUserCmd) Execute([]string) error {
	return nil
}
func (g *getUserCmd) Execute([]string) error {
	return nil
}
func (u *updateUserCmd) Execute([]string) error {
	return nil
}
func (d *deleteUserCmd) Execute([]string) error {
	return nil
}
func (l *listUserCmd) Execute([]string) error {
	return nil
}
