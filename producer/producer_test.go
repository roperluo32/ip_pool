package producer

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestBasicProducer(t *testing.T) {
	saver := &MockSaver{}
	pd := NewProducer(saver)

	pd.RegisterProxyGetter(&MockGetter{}) // MockGetter 0.8s 产生1个代理
	go pd.Run()
	time.Sleep(time.Millisecond * 3500) //睡眠3.5s，会产生4个代理ip
	pd.Stop()

	num, err := saver.GetTotalNum()
	assert.Nil(t, err)
	assert.Equal(t, 4, num)
}

func TestProducerZero(t *testing.T) {
	saver := &MockSaver{}
	pd := NewProducer(saver)

	pd.RegisterProxyGetter(&MockGetter{}) // MockGetter 0.8s 产生1个代理
	go pd.Run()
	time.Sleep(time.Millisecond * 500) //睡眠0.5s，会产生0个代理ip
	pd.Stop()

	num, err := saver.GetTotalNum()
	assert.Nil(t, err)
	assert.Equal(t, 0, num)
}
