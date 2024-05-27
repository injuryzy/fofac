package fetch

import (
	"testing"
)

func TestName(t *testing.T) {
	search := FofaSearch{
		FofaQuery: FofaQuery{
			Page:   1,
			Size:   10,
			Full:   true,
			Key:    "6c72716f741f878b25547bed0bdb716f",
			Email:  "jensenhuang455@gmail.com",
			Query:  `asn="40065" && title="电影"`,
			Before: "2024-05-25",
			//After:  "2024-05-20",

			TimeInterval: 1,
		},
	}
	search.QueryResult()
	//search.QueryT()
}
