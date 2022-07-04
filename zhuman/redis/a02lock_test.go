package main

import (
	"log"
	"math"
	"strings"
	"testing"
	"time"

	"github.com/go-redis/redis/v7"
	uuid "github.com/satori/go.uuid"
)

func TestLock(t *testing.T) {
	login := "TestName1"
	llogin := strings.ToLower(login)
	lock := AcquireLockWithTimeout("user:"+llogin, 10, 10)
	defer ReleaseLock("user:"+llogin, lock)
}

func AcquireLockWithTimeout(lockname string, acquireTimeout, lockTimeout float64) string {
	identifier := uuid.NewV4().String()
	lockname = "lock:" + lockname
	finalLockTimeout := math.Ceil(lockTimeout)
	log.Println(identifier)

	end := time.Now().UnixNano() + int64(acquireTimeout*1e9)
	for time.Now().UnixNano() < end {
		if Conn.SetNX(lockname, identifier, 0).Val() {
			Conn.Expire(lockname, time.Duration(finalLockTimeout)*time.Second)
			return identifier
		} else if Conn.TTL(lockname).Val() < 0 {
			Conn.Expire(lockname, time.Duration(finalLockTimeout)*time.Second)
		}
		time.Sleep(10 * time.Millisecond)
	}
	return ""
}

func ReleaseLock(lockname, identifier string) bool {
	lockname = "lock:" + lockname
	var flag = true
	for flag {
		err := Conn.Watch(func(tx *redis.Tx) error {
			pipe := tx.TxPipeline()
			if tx.Get(lockname).Val() == identifier {
				pipe.Del(lockname)
				if _, err := pipe.Exec(); err != nil {
					return err
				}
				flag = true
				return nil
			}

			tx.Unwatch()
			flag = false
			return nil
		})

		if err != nil {
			log.Println("watch failed in ReleaseLock, err is: ", err)
			return false
		}
	}
	return flag
}
