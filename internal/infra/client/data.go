package client

import "github.com/bytopia/kratos-ddd-template/internal/pkg/logging"

// Data .
type Data struct {
	// TODO wrapped database client
}

// NewData .
func NewData(la logging.Adapter) (*Data, func(), error) {
	cleanup := func() {
		logging.NewLogger(la).Info("closing the data resources")
	}
	return &Data{}, cleanup, nil
}
