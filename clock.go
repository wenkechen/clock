package clock

import (
	"sync"
	"time"
)

var now = time.Now()
var once sync.Once

var expireTime time.Time

func Sync(t time.Time, expireDate string) error {
	now = t
	var err error
	expireTime, err = time.Parse(time.DateOnly, expireDate)
	if err != nil {
		return err
	}

	once.Do(func() {
		go func() {
			ticker := time.NewTicker(time.Second)

			for {
				select {
				case <-ticker.C:
					now = now.Add(time.Second)
				}
			}
		}()
	})

	return nil
}

func Now() time.Time {
	return now
}

func IsExpired() bool {
	return expireTime.Before(now)
}
