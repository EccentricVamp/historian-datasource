package plugin

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/grafana/grafana-plugin-sdk-go/backend"
	"github.com/grafana/grafana-plugin-sdk-go/backend/instancemgmt"
	"github.com/grafana/grafana-plugin-sdk-go/backend/log"
	"github.com/grafana/grafana-plugin-sdk-go/data"
	"golang.org/x/oauth2"
)

type Datasource struct {
	log log.Logger

	client   *http.Client
	URL      *url.URL
	config   *oauth2.Config
	username string
	password string
}

func NewDatasource(settings backend.DataSourceInstanceSettings) (instancemgmt.Instance, error) {
	log := log.New()
	url, _ := url.Parse(settings.URL)
	password := settings.DecryptedSecureJSONData["password"]

	return &Datasource{
		log: log,

		URL: url,
		config: &oauth2.Config{
			ClientID:     "historian_public_rest_api",
			ClientSecret: "publicapisecret",
			Endpoint: oauth2.Endpoint{
				TokenURL: settings.URL + "/uaa/oauth/token",
			},
		},
		username: settings.User,
		password: password,
	}, nil
}

// CheckHealth handles health checks sent from Grafana to the plugin.
// The main use case for these health checks is the test button on the
// datasource configuration page which allows users to verify that
// a datasource is working as expected.
func (d *Datasource) CheckHealth(ctx context.Context, req *backend.CheckHealthRequest) (*backend.CheckHealthResult, error) {
	d.log.Info("CheckHealth called", "request", req)

	var status = backend.HealthStatusOk
	var message = "Data source is working"

	// TODO: implement this

	return &backend.CheckHealthResult{
		Status:  status,
		Message: message,
	}, nil
}

// QueryData handles multiple queries and returns multiple responses.
// req contains the queries []DataQuery (where each query contains RefID as a unique identifer).
// The QueryDataResponse contains a map of RefID to the response for each query, and each response
// contains Frames ([]*Frame).
func (d *Datasource) QueryData(ctx context.Context, req *backend.QueryDataRequest) (*backend.QueryDataResponse, error) {
	d.log.Info("QueryData called", "request", req)

	response := backend.NewQueryDataResponse()
	var raw RawDataQuery

	for _, q := range req.Queries {
		res := backend.DataResponse{}

		res.Error = json.Unmarshal(q.JSON, &raw)
		raw.Start = q.TimeRange.From
		raw.End = q.TimeRange.To
		if res.Error != nil {
			response.Responses[q.RefID] = res
			continue
		}

		data, err := d.listRawData(ctx, raw)
		if err != nil {
			res.Error = err
			response.Responses[q.RefID] = res
			continue
		}

		res.Frames, err = appendFrames(res.Frames, data, raw.MinQuality)
		if err != nil {
			res.Error = err
			response.Responses[q.RefID] = res
			continue
		}

		response.Responses[q.RefID] = res
	}

	return response, nil
}

func (d *Datasource) listRawData(ctx context.Context, query RawDataQuery) (*Response, error) {
	client, err := d.getClient(ctx)
	if err != nil {
		return nil, err
	}

	start := query.Start.Format(time.RFC3339)
	end := query.End.Format(time.RFC3339)
	u := fmt.Sprintf("%s/historian-rest-api/v1/datapoints/raw?tagNames=%v&start=%v&end=%v&direction=%v&count=%v", d.URL, query.Tags, start, end, query.Direction, query.Count)
	resp, err := client.Get(u)
	if err != nil {
		d.client = nil
		return nil, err
	}

	var response = new(Response)
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func appendFrames(frames data.Frames, r *Response, minQ int) (data.Frames, error) {
	var error error
	frame := data.NewFrame(
		"samples",
		data.NewField("time", nil, []time.Time{}),
		data.NewField("value", nil, []float64{}),
	)
	for _, d := range r.Data {
		for _, s := range d.Samples {
			if s.Quality < minQ {
				continue
			}
			value, err := strconv.ParseFloat(s.Value, 64)
			if err != nil {
				error = err
				continue
			}
			frame.AppendRow(
				s.TimeStamp,
				value,
			)
		}
	}

	return append(frames, frame), error
}

func (d *Datasource) getClient(ctx context.Context) (*http.Client, error) {
	if d.client == nil {
		token, err := d.config.PasswordCredentialsToken(ctx, d.username, d.password)
		if err != nil {
			return nil, err
		}
		d.client = d.config.Client(ctx, token)
	}

	return d.client, nil
}
