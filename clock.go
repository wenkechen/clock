package clock

import (
	"fmt"
	"sync"
	"time"

	"github.com/beevik/ntp"
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

func NetTime() time.Time {
	for i := 1; i <= 6; i++ {
		serverAddr := fmt.Sprintf("ntp%d.aliyun.com", i)
		currentTime, err := ntp.Time(serverAddr)
		if err == nil {
			return currentTime
		}
	}

	return time.Time{}
}
