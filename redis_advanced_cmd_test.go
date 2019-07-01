package godis

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestRedis_ConfigGet(t *testing.T) {
	redis := NewRedis(option)
	defer redis.Close()
	reply, err := redis.ConfigGet("timeout")
	assert.Nil(t, err, "err is nil")
	assert.Equal(t, []string{"timeout", "0"}, reply)
}

func TestRedis_ConfigSet(t *testing.T) {
	redis := NewRedis(option)
	defer redis.Close()
	reply, err := redis.ConfigSet("timeout", "30")
	assert.Nil(t, err)
	assert.Equal(t, "OK", reply)
	reply1, err := redis.ConfigGet("timeout")
	assert.Nil(t, err)
	assert.Equal(t, []string{"timeout", "30"}, reply1)
	reply, err = redis.ConfigSet("timeout", "0")
	assert.Nil(t, err)
	assert.Equal(t, "OK", reply)
	reply1, err = redis.ConfigGet("timeout")
	assert.Nil(t, err)
	assert.Equal(t, []string{"timeout", "0"}, reply1)
}

func TestRedis_SlowlogGet(t *testing.T) {
	redis := NewRedis(option)
	defer redis.Close()
	_, err := redis.SlowLogGet()
	assert.Nil(t, err)

	arr, err := redis.SlowLogGet(1)
	assert.Nil(t, err)
	assert.NotEmpty(t, arr)
}

func TestRedis_SlowlogLen(t *testing.T) {
	redis := NewRedis(option)
	defer redis.Close()
	l, err := redis.SlowLogLen()
	assert.Nil(t, err)
	assert.True(t, l >= 0)
}

func TestRedis_SlowlogReset(t *testing.T) {
	redis := NewRedis(option)
	defer redis.Close()
	str, err := redis.SlowLogReset()
	assert.Nil(t, err)
	assert.Equal(t, "OK", str)
}

func TestRedis_ObjectEncoding(t *testing.T) {
	flushAll()
	redis := NewRedis(option)
	defer redis.Close()
	redis.Set("godis", "good")
	encode, err := redis.ObjectEncoding("godis")
	assert.Nil(t, err)
	assert.Equal(t, "embstr", encode)
	redis.Set("godis", "12")
	encode, err = redis.ObjectEncoding("godis")
	assert.Nil(t, err)
	assert.Equal(t, "int", encode)
}

func TestRedis_ObjectIdletime(t *testing.T) {
	flushAll()
	redis := NewRedis(option)
	defer redis.Close()
	redis.Set("godis", "good")
	time.Sleep(1000 * time.Millisecond)
	idle, err := redis.ObjectIdleTime("godis")
	assert.Nil(t, err)
	assert.True(t, idle > 0)
}

func TestRedis_ObjectRefcount(t *testing.T) {
	flushAll()
	redis := NewRedis(option)
	defer redis.Close()
	redis.Set("godis", "good")
	count, err := redis.ObjectRefCount("godis")
	assert.Nil(t, err)
	assert.Equal(t, int64(1), count)
}
