package log

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/getsentry/sentry-go"
	"github.com/pkg/errors"
)

// nolint: gochecknoinits
func init() {
	err := sentry.Init(sentry.ClientOptions{
		Dsn:         os.Getenv("SENTRY_APIS_TARGET"), // Explicitly set token
		Debug:       false,                           // Enable printing of SDK debug messages.
		Environment: os.Getenv("ENVIRONMENT"),
	})
	if err != nil {
		log.Printf("sentry.Init: %s", err)
	}

	defer sentry.Flush(2 * time.Second)
}

// notifySentry formats and sends an error
func notifySentry(err error) {
	sentry.WithScope(func(scope *sentry.Scope) {
		scope.AddEventProcessor(func(event *sentry.Event, hint *sentry.EventHint) *sentry.Event {
			if len(event.Exception) == 0 {
				return event
			}

			// sentry seems to take the last exception in the list to use for the issue title and type
			event.Exception[len(event.Exception)-1].Type = fmt.Sprintf("%T", errors.Cause(hint.OriginalException))
			event.Exception[len(event.Exception)-1].Value = hint.OriginalException.Error()

			return event
		})
		sentry.CaptureException(err)
	})
}
