package model

// IPType ip类型
type IPType int

const (
	// TypGaoNi 高匿
	TypGaoNi IPType = iota
)

// IPItem 一个ip
type IPItem struct {
	IP   string
	Port int
	Typ  IPType
}
