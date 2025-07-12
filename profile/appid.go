package profile

import (
	"context"

	"github.com/go-pantheon/fabrica-util/errors"
	"github.com/redis/go-redis/v9"
)

var (
	appID int64
)

func InitAppID(cli redis.UniversalClient) error {
	i, err := cli.Incr(context.Background(), "appid").Result()
	if err != nil {
		return errors.Wrap(err, "load appid failed")
	}

	appID = i

	return nil
}

func AppID() int64 {
	return appID
}
