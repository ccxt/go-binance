package portfolio

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type umTransactionDownloadLinkServiceTestSuite struct {
	baseTestSuite
}

func TestUMTransactionDownloadLinkService(t *testing.T) {
	suite.Run(t, new(umTransactionDownloadLinkServiceTestSuite))
}

func (s *umTransactionDownloadLinkServiceTestSuite) TestGetUMTransactionDownloadLink() {
	data := []byte(`{
		"downloadId": "545923594199212032",
		"status": "completed",
		"url": "www.binance.com",
		"s3Link": null,
		"notified": true,
		"expirationTimestamp": 1645009771000,
		"isExpired": null
	}`)
	s.mockDo(data, nil)
	defer s.assertDo()

	downloadID := "545923594199212032"
	recvWindow := int64(5000)

	s.assertReq(func(r *request) {
		e := newSignedRequest()
		e.setParam("downloadId", downloadID)
		e.setParam("recvWindow", recvWindow)
		s.assertRequestEqual(e, r)
	})

	res, err := s.client.NewGetUMTransactionDownloadLinkService().
		DownloadID(downloadID).
		RecvWindow(recvWindow).
		Do(newContext())

	s.r().NoError(err)
	s.r().Equal("545923594199212032", res.DownloadID)
	s.r().Equal("completed", res.Status)
	s.r().Equal("www.binance.com", res.URL)
	s.r().Equal(true, res.Notified)
	s.r().Equal(int64(1645009771000), res.ExpirationTimestamp)
}
