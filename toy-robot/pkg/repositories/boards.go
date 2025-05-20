// An in memory repository

package repositories

import (
	"errors"

	"github.com/nicholasf/go-exercises/toy-robot/pkg/entities"
)

type BoardRepository interface {
	Get(id int) (entities.Board, error)
}

type Boards struct {
	boards map[int]entities.Board
}

func NewBoardRepository() BoardRepository {
	b := &Boards{
		boards: make(map[int]entities.Board),
	}

	return b
}

func (b Boards) Get(id int) (entities.Board, error) {
	bd, ok := b.boards[id]

	if !ok {
		return entities.Board{}, errors.New("no such board")
	}

	return bd, nil
}
