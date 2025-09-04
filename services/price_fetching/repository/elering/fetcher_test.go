package elering

import "testing"

func TestCall(t *testing.T) {

	fetcher := NewEleringFetcher("https://dashboard.elering.ee")
	fetcher.FetchNextDay()
}
