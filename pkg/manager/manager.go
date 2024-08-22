package manager

import (
	"errors"
"log"
	"github.com/tonnytg/models-persona/internal/domain"
)

func Add(m domain.Model) (*domain.Model, error) {
  log.Println("model add with success.")
  if m.Name == "" {
    return nil, errors.New("error to add model")
  }
  return &m, nil
}
