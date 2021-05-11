package errorhandling

import (
	"errors"
	"fmt"
)

func ItemWithIdNotFound(name string, id int) error {
	return errors.New(fmt.Sprintf("%v with ID %v not found", name, id))
}
