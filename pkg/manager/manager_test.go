package manager_test

import (
	"testing"

	"github.com/tonnytg/models-persona/internal/domain"
	"github.com/tonnytg/models-persona/pkg/manager"

)

func TestManagerAdd(t *testing.T) {
  model := domain.Model{
    ID: 1,
    Name: "test name",
    Size: 100,
    Version: "1.0.0",
  }

  _, err := manager.Add(model)
  if err != nil {
    t.Error("error to add model")
  }

}
