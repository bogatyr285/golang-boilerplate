//go:generate mockgen -source=./doer.go -destination=../../../mocks/smth_fetcher_mock.go -package=mocks

package doer

import (
	"context"
	"fmt"
	"strings"

	"github.com/bogatyr285/golang-boilerplate/internal/models"
	"go.opencensus.io/trace"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc/codes"
)

const (
	threadMax = 10
)

type SomethingFetcher interface {
	GetSomething(ctx context.Context, query string) ([]*models.Something, error)
}

// pretend like we fetch smth from db, process it somehow and return back response
type Doer struct {
	somethingFetcher SomethingFetcher
	threadMax        int
}

func NewDoer(somethingFetcher SomethingFetcher) *Doer {
	return &Doer{somethingFetcher: somethingFetcher, threadMax: threadMax}
}

func (d *Doer) Do(ctx context.Context, input string) (string, error) {
	// could be decorator
	ctx, span := trace.StartSpan(ctx, "services.doer.do")
	defer span.End()

	records, err := d.somethingFetcher.GetSomething(ctx, input)
	if err != nil {
		span.AddAttributes(trace.StringAttribute("input", input))
		span.SetStatus(trace.Status{Code: int32(codes.Internal), Message: err.Error()})
		return "", fmt.Errorf("db err: %w", err)
	}
	// lets process it with threads because we can
	processedRecords := make([]string, len(records))
	errGr /*grCtx*/, _ := errgroup.WithContext(ctx)
	errGr.SetLimit(d.threadMax)
	for i := range records {
		i := i
		errGr.Go(func() error {
			// this could be anything which requires separate processing i.e API calls
			s := records[i].FieldString
			processedRecords[i] = s
			return nil
		})
	}
	if err := errGr.Wait(); err != nil {
		span.AddAttributes(trace.StringAttribute("input", input))
		span.SetStatus(trace.Status{Code: int32(codes.Internal), Message: err.Error()})
		return "", fmt.Errorf("processing err: %w", err)
	}
	return strings.Join(processedRecords, " "), nil
}
