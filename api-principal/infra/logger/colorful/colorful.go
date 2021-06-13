package colorful

import (
	"github.com/thyagofr/api-crud/infra/logger"
	"log"
	"os"
)

// Colorful holds 3 types of logger that prints in 3 different colors
type Colorful struct {
	// info prints info messages
	info *log.Logger

	// warning prints warning messages
	warning *log.Logger

	// err prints error messages
	err *log.Logger
}

// NewColorful returns a new colorful logger
func NewColorful() logger.Logger {
	return &Colorful{
		info:    log.New(os.Stdout, string(White)+"[INFO]\t", log.Ldate|log.Ltime),
		warning: log.New(os.Stdout, string(Yellow)+"[WARNING]\t", log.Ldate|log.Ltime),
		err:     log.New(os.Stdout, string(Red)+"[ERROR]\t", log.Ldate|log.Ltime),
	}
}

// Info prints relevant information
func (c *Colorful) Info(format string, v ...interface{}) {
	c.info.Printf(format, v...)
}

// Warning raises a warning
func (c *Colorful) Warning(format string, v ...interface{}) {
	c.warning.Printf(format, v...)
}

// Error alerts about an error
func (c *Colorful) Error(format string, v ...interface{}) {
	c.err.Printf(format, v...)
}

// Fatal alerts about a fatal error and exits the application
func (c *Colorful) Fatal(format string, v ...interface{}) {
	c.err.Printf(format, v...)
	os.Exit(1)
}
