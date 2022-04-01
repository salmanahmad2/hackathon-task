package server

import "net/url"

func GetURLValues(token string) (url.Values, error) {
	queryData, err := url.Parse(token)
	if err != nil {
		return nil, err
	}
	values, err := url.ParseQuery(queryData.RawQuery)
	if err != nil {
		return nil, err
	}
	return values, nil

}
