package options

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type IndexServiceTestSuite struct {
	baseTestSuite
}

func TestIndexService(t *testing.T) {
	suite.Run(t, new(IndexServiceTestSuite))
}

func (s *IndexServiceTestSuite) TestIndex() {
	data := []byte(`{
		"indexPrice": "68193.79851064",
		"time": 1717036339000
	   }`)
	s.mockDo(data, nil)
	defer s.assertDo()

	index, err := s.client.NewIndexService().Do(newContext())

	targetIndex := &Index{
		IndexPrice: "68193.79851064",
		Time:       1717036339000,
	}

	s.r().Equal(err, nil, "err != nil")
	s.r().Equal(index.IndexPrice, targetIndex.IndexPrice, "IndexPrice")
	s.r().Equal(index.Time, targetIndex.Time, "Time")
}
