package test

import (
	
	"poke/internal"
	"testing"
	"time"
)




func TestAdd(t *testing.T){
	t.Parallel()


	


}


func TestReapLoop(t *testing.T) {
	const baseTime = 5 * time.Millisecond
	const waitTime = baseTime + 5*time.Millisecond
	cache := internal.NewCache(baseTime)
	cache.Add("url", []byte("testdata"))

	_, ok := cache.Get("url")
	if !ok {
		t.Errorf("expected to find key")
		return
	}

	time.Sleep(waitTime)

	_, ok = cache.Get("url")
	if ok {
		t.Errorf("expected to not find key")
		return
	}
}