package model

type App interface {
	start() (bool , error)
}