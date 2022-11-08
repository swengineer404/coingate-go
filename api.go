package coingate

import "net/http"

type apiClient struct {
	client *http.Client
	key    string
}

func newAPIClient(key string) {

}
