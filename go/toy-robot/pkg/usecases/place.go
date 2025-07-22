// The Place command

package usecases

import (
	"github.com/nicholasf/go-exercises/toy-robot/pkg/entities"
)

type Command interface {
	Issue(b entities.Board) (boardId int, err error)
}

type Command2 struct {
	Type string
}

type Place struct {
	Command
	X      int
	Y      int
	Facing int
}
