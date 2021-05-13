package notifier

import "io"

type Notifier interface {
	Notify(message io.Reader) error
}
