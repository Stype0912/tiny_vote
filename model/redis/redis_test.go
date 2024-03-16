package redis

import "testing"

func TestRedisConnect(t *testing.T) {
	t.Parallel()
	Init()
	var err error
	err = rdb.Set("ticket", "123", 0).Err()
	if err != nil {
		t.Error(err)
	}
	err = rdb.Set("ticket", "456", 0).Err()
	if err != nil {
		t.Error(err)
	}
	t.Log(GetTicket())
}
