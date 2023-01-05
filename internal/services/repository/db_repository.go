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

type DatabaseRepository interface {
	GetSomething(ctx context.Context, query string) ([]*models.Something, error)
	InsertSomething(ctx context.Context, records *models.Something) error
}

// whatever db you like
type MongoDB struct {
}

func NewMongoDB( /*config*/ ) *MongoDB {
	return &MongoDB{}
}

func (d *MongoDB) GetSomething(ctx context.Context, query string) ([]*models.Something, error) {
	// pretend its from db
	return []*models.Something{{FieldString: query, FieldInt: 1337}}, nil
}
func (d *MongoDB) InsertSomething(ctx context.Context, records *models.Something) error {
	//return d.db.BatchInsert(ctx,records)
	return nil
}
