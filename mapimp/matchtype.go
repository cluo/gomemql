package gomemql_map

type matchType int

// 数据匹配模式
const (
	matchType_Unknown = iota
	matchType_Equal
	matchType_NotEqual
	matchType_Great
	matchType_GreatEqual
	matchType_Less
	matchType_LessEqual
	matchType_MAX
)

var sign2MatchType = map[string]matchType{
	"==": matchType_Equal,
	"!=": matchType_NotEqual,
	">":  matchType_Great,
	">=": matchType_GreatEqual,
	"<":  matchType_Less,
	"<=": matchType_LessEqual,
}

var matchType2Sign = map[matchType]string{}

func getMatchTypeBySign(sign string) matchType {

	if v, ok := sign2MatchType[sign]; ok {
		return v
	}

	return matchType_Unknown
}

func getSignByMatchType(t matchType) string {
	if v, ok := matchType2Sign[t]; ok {
		return v
	}

	return "unknown"
}

func init() {

	for k, v := range sign2MatchType {
		matchType2Sign[v] = k
	}
}
