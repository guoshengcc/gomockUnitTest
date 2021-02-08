package main

// Ifoo test interface
type Ifoo interface {
	GetData(id string) (string, error)
}
