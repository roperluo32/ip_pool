package producer

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"time"
)


func TestBasicProducer(t *testing.T) {
	saver := &MockSaver{}
	pd := NewProducer(saver)

	pd.RegisterProxyGetter(&MockGetter{})	// MockGetter 0.8s 产生1个代理
	go pd.Run()
	time.Sleep(time.Millisecond * 3500)    //睡眠3.5s，会产生4个代理ip
	pd.Stop()

	assert.Equal(t, 4, saver.GetTotalNum())
}

func TestProducerTen(t *testing.T) {
	saver := &MockSaver{}
	pd := NewProducer(saver)

	pd.RegisterProxyGetter(&MockGetter{})	// MockGetter 0.8s 产生1个代理
	go pd.Run()
	time.Sleep(time.Millisecond * 8500)    //睡眠8.5s，会产生10个代理ip
	pd.Stop()

	assert.Equal(t, 10, saver.GetTotalNum())
}