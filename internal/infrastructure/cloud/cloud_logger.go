package cloud

import (
	"context"

	"cloud.google.com/go/logging"
	"github.com/pabloantipan/go-api-gateway-poc/config"
	"google.golang.org/api/option"
)

var logName = "go-api-gateway-poc"

type CloudLogger struct {
	client *logging.Client
	logger *logging.Logger
}

func NewCloudLogger(cfg *config.Config) (*CloudLogger, error) {
	ctx := context.Background()

	client, err := logging.NewClient(ctx, cfg.ProjectID,
		option.WithCredentialsFile(cfg.CloudLoggingCredentialsFile),
	)
	if err != nil {
		return nil, err
	}

	logger := client.Logger(logName)

	return &CloudLogger{
		client: client,
		logger: logger,
	}, nil
}

func (cl *CloudLogger) LogRequest(ctx context.Context, method, path string, status int) {
	cl.logger.Log(logging.Entry{
		Payload: map[string]interface{}{
			"method": method,
			"path":   path,
			"status": status,
			"type":   "request",
		},
		Severity: logging.Info,
	})
}

func (cl *CloudLogger) LogError(ctx context.Context, err error, method, path string) {
	cl.logger.Log(logging.Entry{
		Payload: map[string]interface{}{
			"error":  err.Error(),
			"method": method,
			"path":   path,
			"type":   "error",
		},
		Severity: logging.Error,
	})
}
