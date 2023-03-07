package main

type workCmd struct {
	Add    addWorkCmd
	Get    getWorkCmd
	Update updateWorkCmd
	Delete deleteWorkCmd
	List   listWorkCmd
}

type addWorkCmd struct{}
type getWorkCmd struct{}
type updateWorkCmd struct{}
type deleteWorkCmd struct{}
type listWorkCmd struct{}

func (a *addWorkCmd) Execute([]string) error {
	return nil
}
func (g *getWorkCmd) Execute([]string) error {
	return nil
}
func (u *updateWorkCmd) Execute([]string) error {
	return nil
}
func (d *deleteWorkCmd) Execute([]string) error {
	return nil
}
func (l *listWorkCmd) Execute([]string) error {
	return nil
}
