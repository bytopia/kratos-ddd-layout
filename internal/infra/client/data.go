package client

import (
	"github.com/bytopia/kratos-ddd-template/pkg/logging"
)

// Data .
type Data struct {
	// TODO wrapped database client
}

// NewData .
func NewData(log logging.LoggerAdapter) (*Data, func(), error) {
	cleanup := func() {
		log.Info("closing the data resources")
	}
	return &Data{}, cleanup, nil
}
