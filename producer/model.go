package producer

// IPType ip类型
type IPType int
const (
	// TypGaoNi 高匿
	TypGaoNi IPType = iota  
) 

// IPItem 一个ip
type IPItem struct {
	ip string
	port int
	typ IPType
}