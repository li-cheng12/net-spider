package fetcher

import (
	"testing"
)

func TestFetch(t *testing.T) {
	Fetch("http://album.zhenai.com/u/1166905636") // 有的url会403
}


