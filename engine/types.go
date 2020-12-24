package engine

// 每一个url 对应一个解析器
type Request struct {
	Url       string
	ParserFun func([]byte) ParserResult
}

type ParserResult struct {
	Requests []Request
	Items    []interface{}
}

func NilParser([]byte) ParserResult {
	return ParserResult{}
}
