package timestamp

import (
	"fmt"
	"testing"
)

func TestTimestamp(t *testing.T) {
	n := Now(nil)
	ts := New(&n)
	d, e := ts.MarshalJSON()
	fmt.Printf("ts:%s,error:%v", string(d), e)
}
