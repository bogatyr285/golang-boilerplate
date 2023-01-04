package repository

import (
	"context"
	"fmt"

	"github.com/bogatyr285/golang-boilerplate/internal/models"
)

var (
	ErrNotFound      = fmt.Errorf("not found")
	ErrInternalDBErr = fmt.Errorf("internal db error")
)

// whatever db you like
type DatabaseRepository struct {
}

func NewDatabaseRepository( /*config*/ ) *DatabaseRepository {
	return &DatabaseRepository{}
}

func (d *DatabaseRepository) GetSomething(ctx context.Context, query string) ([]*models.Something, error) {
	// pretend its from db
	return []*models.Something{{FieldString: query, FieldInt: 1337}}, nil
}
func (d *DatabaseRepository) InsertSomething(ctx context.Context, records *models.Something) error {
	//return d.db.BatchInsert(ctx,records)
	return nil
}
