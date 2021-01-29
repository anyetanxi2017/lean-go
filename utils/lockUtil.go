package utils

import (
	"gin-vue-admin/global"
	"github.com/go-redis/redis"
	uuid "github.com/satori/go.uuid"
	"go.uber.org/zap"
	"time"
)

const LockKey = "lock_"
const LockErr = "操作太快"

// 分布式锁实现
type RLock struct {
	cli        *redis.Client
	key        string
	value      string
	expiration time.Duration
}

// 加锁
func (rl *RLock) Lock() (lockSuccess bool) {
	// 未存在锁则进行加锁
	isSet, err := rl.cli.SetNX(rl.key, rl.value, rl.expiration).Result()
	if err != nil {
		global.GVA_LOG.Error("加锁错误!"+err.Error(), zap.Any("err", err))
		return false
	}
	if !isSet {
		return false
	}
	return true
}

// 解锁
func (rl *RLock) Unlock() {
	// 如果key为空不进行解锁
	if rl.key == "" {
		return
	}
	// 检验该锁是否为该线程加的锁，不是则返回
	val := rl.cli.Get(rl.key).Val()
	if val == "" {
		return
	}
	if val != rl.value {
		return
	}
	// 是该线程加的锁则进行解除
	if err := rl.cli.Del(rl.key).Err(); err != nil {
		global.GVA_LOG.Error("锁解除错误!"+err.Error(), zap.Any("err", err))
	}
}

type LockUtil struct {
}

func (l LockUtil) NewRLockDefault(k string) *RLock {
	return l.NewRLock(LockKey+k, 0)
}
func (l LockUtil) NewRLock(key string, expiration time.Duration) *RLock {
	if expiration == 0 {
		expiration = 10 * time.Second
	}
	return &RLock{
		cli:        global.GVA_REDIS,
		key:        key,
		value:      uuid.NewV1().String(),
		expiration: expiration,
	}

}
