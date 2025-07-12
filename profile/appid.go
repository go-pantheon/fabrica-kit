package profile

import (
	"context"
	"strconv"
	"time"

	"github.com/go-pantheon/fabrica-util/errors"
	"github.com/redis/go-redis/v9"
)

const (
	appIDKey = "appid"
	maxAppID = 0xFFFF // 65535
)

var (
	appID int64
)

//nolint:gocognit
func InitAppID(cli redis.UniversalClient) error {
	for range 3 {
		// incr appid first
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		id, err := cli.Incr(ctx, appIDKey).Result()
		if err != nil {
			return errors.Wrap(err, "incr appid failed")
		}

		if id <= maxAppID {
			appID = id
			return nil
		}

		txf := func(tx *redis.Tx) error {
			currentStr, err := tx.Get(ctx, appIDKey).Result()
			if err != nil && !errors.Is(err, redis.Nil) {
				return errors.Wrap(err, "get current appid in transaction failed")
			}

			var current int64

			if !errors.Is(err, redis.Nil) {
				if current, err = strconv.ParseInt(currentStr, 10, 64); err != nil {
					return errors.Wrap(err, "parse current appid failed")
				}
			}

			pipe := tx.TxPipeline()

			if current <= maxAppID {
				incrCmd := pipe.Incr(ctx, appIDKey)

				if _, err = pipe.Exec(ctx); err != nil {
					return errors.Wrap(err, "incr appid in transaction failed")
				}

				appID = incrCmd.Val()
			} else {
				_ = pipe.Set(ctx, appIDKey, 1, 0)

				if _, err = pipe.Exec(ctx); err != nil {
					return errors.Wrap(err, "reset appid in transaction failed")
				}

				appID = 1
			}

			return nil
		}

		if err = cli.Watch(ctx, txf, appIDKey); err == nil {
			return nil
		}

		if errors.Is(err, redis.TxFailedErr) {
			continue
		}

		return errors.Wrap(err, "watch appid transaction failed")
	}

	return errors.New("appid initialization failed: too many retries")
}

func AppID() int64 {
	return appID
}
