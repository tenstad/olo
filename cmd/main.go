package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	eapi "github.com/tenstad/olo/pkg/eventor/api"
	"k8s.io/utils/ptr"
)

func main() {
	c, err := eapi.NewClientWithResponses("https://eventor.orientering.no/api", func(c *eapi.Client) error {
		c.RequestEditors = []eapi.RequestEditorFn{
			func(_ context.Context, req *http.Request) error {
				req.Header.Add("ApiKey", os.Getenv("TOKEN"))
				return nil
			},
		}
		return nil
	})
	if err != nil {
		panic(err)
	}

	events, err := c.GetEventsWithResponse(context.TODO(), &eapi.GetEventsParams{
		FromDate:        ptr.To("2025-08-01 00:00:00"),
		OrganisationIds: ptr.To("150"),
	})
	if err != nil {
		panic(err)
	}

	fmt.Println(string(events.Body))

	b, err := json.Marshal(events.XML200)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(b))
}
