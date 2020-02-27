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

func TestProducerTen(t *testing.T) {
	saver := &MockSaver{}
	pd := NewProducer(saver)

	pd.RegisterProxyGetter(&MockGetter{}) // MockGetter 0.8s 产生1个代理
	go pd.Run()
	time.Sleep(time.Millisecond * 8500) //睡眠8.5s，会产生10个代理ip
	pd.Stop()

	num, err := saver.GetTotalNum()
	assert.Nil(t, err)
	assert.Equal(t, 10, num)
}
