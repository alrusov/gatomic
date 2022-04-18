package gatomic

import (
	"testing"
	"time"

	"github.com/alrusov/misc"
)

//----------------------------------------------------------------------------------------------------------------------------//

func Test1(t *testing.T) {
	step := 0

	{
		src := misc.NowUTC()
		g := New(src)

		step++
		v := Get(g)
		if !v.Equal(src) {
			t.Errorf(`[%d] %v != %v`, step, misc.Time2JSONutc(v), misc.Time2JSONutc(src))
		}

		step++
		src2 := src.Add(3 * time.Hour)
		old := Set(g, src2)
		if !old.Equal(src) {
			t.Errorf(`[%d] %v != %v`, step, misc.Time2JSONutc(old), misc.Time2JSONutc(src))
		}

		step++
		v = Get(g)
		if !v.Equal(src2) {
			t.Errorf(`[%d] %v != %v`, step, misc.Time2JSONutc(v), misc.Time2JSONutc(src2))
		}
	}

	{
		src := "blah-blah-blah"
		g := New(src)

		step++
		v := Get(g)
		if v != src {
			t.Errorf(`[%d] %v != %v`, step, v, src)
		}

		step++
		src2 := src + "-qwertyuiop"
		old := Set(g, src2)
		if old != src {
			t.Errorf(`[%d] %v != %v`, step, old, src)
		}

		step++
		v = Get(g)
		if v != src2 {
			t.Errorf(`[%d] %v != %v`, step, v, src2)
		}
	}

	{
		src := 3.62
		g := New(src)

		step++
		v := Get(g)
		if v != src {
			t.Errorf(`[%d] %v != %v`, step, v, src)
		}

		step++
		src2 := src + 4.12
		old := Set(g, src2)
		if old != src {
			t.Errorf(`[%d] %v != %v`, step, old, src)
		}

		step++
		v = Get(g)
		if v != src2 {
			t.Errorf(`[%d] %v != %v`, step, v, src2)
		}
	}

	{
		src := uint64(8239)
		g := New(src)

		step++
		v := Get(g)
		if v != src {
			t.Errorf(`[%d] %v != %v`, step, v, src)
		}

		step++
		src2 := src * 94
		old := Set(g, src2)
		if old != src {
			t.Errorf(`[%d] %v != %v`, step, old, src)
		}

		step++
		v = Get(g)
		if v != src2 {
			t.Errorf(`[%d] %v != %v`, step, v, src2)
		}
	}
}

//----------------------------------------------------------------------------------------------------------------------------//
