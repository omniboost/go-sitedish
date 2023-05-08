package sitedish_test

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"

	sitedish "github.com/omniboost/go-sitedish"
)

func TestOrdersArchivedGet(t *testing.T) {
	req := client.NewOrdersArchivedGet()
	s := time.Now().AddDate(0, 0, -7)
	e := time.Now().AddDate(0, 0, 1)
	req.QueryParams().Start = sitedish.DateTime{time.Date(s.Year(), s.Month(), s.Day(), 0, 0, 0, 0, s.Location())}
	req.QueryParams().End = sitedish.DateTime{time.Date(e.Year(), e.Month(), e.Day(), 0, 0, 0, 0, e.Location())}
	req.PathParams().Platform = "all"
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}
