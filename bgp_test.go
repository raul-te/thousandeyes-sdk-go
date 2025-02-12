package thousandeyes

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClient_AddBGPAlertRule(t *testing.T) {
	test := BGP{TestName: String("test"), AlertRules: &[]AlertRule{}}
	expected := BGP{TestName: String("test"), AlertRules: &[]AlertRule{{RuleID: Int(1)}}}
	test.AddAlertRule(1)
	assert.Equal(t, expected, test)
}

func TestClient_GetBGP(t *testing.T) {
	out := `{"test":[{"createdDate":"2020-02-06 15:28:07","createdBy":"ThousandEyes SRE (test.example@thousandeyes.com)","enabled":1,"savedEvent":0,"testId":122621,"testName":"test123","type":"bgp","prefix": "1.2.3.0/20","interval":300,"alertsEnabled":1,"liveShare":0,"probeMode":"AUTO","agents":[{"agentId":48620,"agentName":"Seattle, WA (Trial) - IPv6","agentType":"Cloud","countryId":"US","ipAddresses":["135.84.184.153"],"location":"Seattle Area","network":"Astute Hosting Inc. (AS 54527)","prefix":"135.84.184.0/22"}],"sharedWithAccounts":[{"aid":176592,"name":"Cloudreach"}],"apiLinks":[{"rel":"self","href":"https://api.thousandeyes.com/v6/tests/1226221"},{"rel":"data","href":"https://api.thousandeyes.com/v6/web/dns-trace/1226221"},{"rel":"data","href":"https://api.thousandeyes.com/v6/net/metrics/1226221"},{"rel":"data","href":"https://api.thousandeyes.com/v6/net/path-vis/1226221"},{"rel":"data","href":"https://api.thousandeyes.com/v6/net/bgp-metrics/1226221"}]}]}`
	setup()
	var client = &Client{APIEndpoint: server.URL, AuthToken: "foo"}
	mux.HandleFunc("/tests/122621.json", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)
		_, _ = w.Write([]byte(out))
	})

	// Define expected values from the API (based on the JSON we print out above)
	expected := BGP{
		TestID:        Int64(122621),
		Enabled:       Bool(true),
		CreatedBy:     String("ThousandEyes SRE (test.example@thousandeyes.com)"),
		CreatedDate:   String("2020-02-06 15:28:07"),
		SavedEvent:    Bool(false),
		AlertsEnabled: Bool(true),
		TestName:      String("test123"),
		Type:          String("bgp"),
		LiveShare:     Bool(false),
		Prefix:        String("1.2.3.0/20"),
		SharedWithAccounts: &[]SharedWithAccount{
			{
				AID:              Int(176592),
				AccountGroupName: String("Cloudreach"),
			},
		},
		APILinks: &[]APILink{
			{
				Href: String("https://api.thousandeyes.com/v6/tests/1226221"),
				Rel:  String("self"),
			},
			{
				Href: String("https://api.thousandeyes.com/v6/web/dns-trace/1226221"),
				Rel:  String("data"),
			},
			{
				Href: String("https://api.thousandeyes.com/v6/net/metrics/1226221"),
				Rel:  String("data"),
			},
			{
				Href: String("https://api.thousandeyes.com/v6/net/path-vis/1226221"),
				Rel:  String("data"),
			},
			{
				Href: String("https://api.thousandeyes.com/v6/net/bgp-metrics/1226221"),
				Rel:  String("data"),
			},
		},
	}

	res, err := client.GetBGP(122621)
	teardown()
	assert.Nil(t, err)
	assert.Equal(t, &expected, res)
}

func TestClient_GetBGPJsonError(t *testing.T) {
	out := `{"test":[{"createdDate":"2020-02-06 15:28:07",createdBy":"ThousandEyes SRE (test.example@thousandeyes.com)","enabled":1,"savedEvent":0,"testId":122621,"testName":"test123","type":"bgp","prefix": "1.2.3.0/20","interval":300,"alertsEnabled":1,"liveShare":0,"probeMode":"AUTO","agents":[{"agentId":48620,"agentName":"Seattle, WA (Trial) - IPv6","agentType":"Cloud","countryId":"US","ipAddresses":["135.84.184.153"],"location":"Seattle Area","network":"Astute Hosting Inc. (AS 54527)","prefix":"135.84.184.0/22"}],"sharedWithAccounts":[{"aid":176592,"name":"Cloudreach"}],"domain": "webex.com","dnsTransportProtocol":  "UDP"}]"apiLinks":[{"rel":"self","href":"https://api.thousandeyes.com/v6/tests/1226221"},{"rel":"data","href":"https://api.thousandeyes.com/v6/web/dns-trace/1226221"},{"rel":"data","href":"https://api.thousandeyes.com/v6/net/metrics/1226221"},{"rel":"data","href":"https://api.thousandeyes.com/v6/net/path-vis/1226221"},{"rel":"data","href":"https://api.thousandeyes.com/v6/net/bgp-metrics/1226221"}]}]}`
	setup()
	var client = &Client{APIEndpoint: server.URL, AuthToken: "foo"}
	mux.HandleFunc("/tests/122621.json", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)
		_, _ = w.Write([]byte(out))
	})
	_, err := client.GetBGP(122621)
	assert.Error(t, err)
	assert.EqualError(t, err, "Could not decode JSON response: invalid character 'c' looking for beginning of object key string")
}

func TestClient_CreateBGP(t *testing.T) {
	out := `{"test":[{"createdDate":"2020-02-06 15:28:07","createdBy":"ThousandEyes SRE (test.example@thousandeyes.com)","enabled":1,"savedEvent":0,"testId":122621,"testName":"test123","type":"bgp","prefix": "1.2.3.0/20","interval":300,"alertsEnabled":1,"liveShare":0,"probeMode":"AUTO","agents":[{"agentId":48620,"agentName":"Seattle, WA (Trial) - IPv6","agentType":"Cloud","countryId":"US","ipAddresses":["135.84.184.153"],"location":"Seattle Area","network":"Astute Hosting Inc. (AS 54527)","prefix":"135.84.184.0/22"}],"sharedWithAccounts":[{"aid":176592,"name":"Cloudreach"}],"apiLinks":[{"rel":"self","href":"https://api.thousandeyes.com/v6/tests/1226221"},{"rel":"data","href":"https://api.thousandeyes.com/v6/web/dns-trace/1226221"},{"rel":"data","href":"https://api.thousandeyes.com/v6/net/metrics/1226221"},{"rel":"data","href":"https://api.thousandeyes.com/v6/net/path-vis/1226221"},{"rel":"data","href":"https://api.thousandeyes.com/v6/net/bgp-metrics/1226221"}]}]}`
	setup()
	var client = &Client{APIEndpoint: server.URL, AuthToken: "foo"}
	mux.HandleFunc("/tests/bgp/new.json", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "POST", r.Method)
		w.WriteHeader(http.StatusCreated)
		_, _ = w.Write([]byte(out))
	})

	// Define expected values from the API (based on the JSON we print out above)
	expected := BGP{
		TestID:        Int64(122621),
		Enabled:       Bool(true),
		CreatedBy:     String("ThousandEyes SRE (test.example@thousandeyes.com)"),
		CreatedDate:   String("2020-02-06 15:28:07"),
		SavedEvent:    Bool(false),
		TestName:      String("test123"),
		Type:          String("bgp"),
		LiveShare:     Bool(false),
		Prefix:        String("1.2.3.0/20"),
		AlertsEnabled: Bool(true),
		SharedWithAccounts: &[]SharedWithAccount{
			{
				AID:              Int(176592),
				AccountGroupName: String("Cloudreach"),
			},
		},

		APILinks: &[]APILink{
			{
				Href: String("https://api.thousandeyes.com/v6/tests/1226221"),
				Rel:  String("self"),
			},
			{
				Href: String("https://api.thousandeyes.com/v6/web/dns-trace/1226221"),
				Rel:  String("data"),
			},
			{
				Href: String("https://api.thousandeyes.com/v6/net/metrics/1226221"),
				Rel:  String("data"),
			},
			{
				Href: String("https://api.thousandeyes.com/v6/net/path-vis/1226221"),
				Rel:  String("data"),
			},
			{
				Href: String("https://api.thousandeyes.com/v6/net/bgp-metrics/1226221"),
				Rel:  String("data"),
			},
		},
	}
	create := BGP{
		TestName: String("test1"),
		Prefix:   String("1.2.3.0/20"),
	}
	res, err := client.CreateBGP(create)
	teardown()
	assert.Nil(t, err)
	assert.Equal(t, &expected, res)
}

func TestClient_DeleteBGP(t *testing.T) {
	setup()

	mux.HandleFunc("/tests/bgp/1/delete.json", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNoContent)
		assert.Equal(t, "POST", r.Method)
	})

	var client = &Client{APIEndpoint: server.URL, AuthToken: "foo"}
	id := 1
	err := client.DeleteBGP(id)

	if err != nil {
		t.Fatal(err)
	}
}

func TestClient_UpdateBGP(t *testing.T) {
	setup()
	out := `{"test":[{"testId":1,"testName":"test123","type":"bgp","prefix": "1.2.3.0/20" }]}`
	mux.HandleFunc("/tests/bgp/1/update.json", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "POST", r.Method)
		_, _ = w.Write([]byte(out))
	})

	var client = &Client{APIEndpoint: server.URL, AuthToken: "foo"}
	id := 1
	dnsS := BGP{}
	res, err := client.UpdateBGP(id, dnsS)
	if err != nil {
		t.Fatal(err)
	}
	expected := BGP{
		TestID:   Int64(1),
		TestName: String("test123"),
		Type:     String("bgp"),
		Prefix:   String("1.2.3.0/20"),
	}
	assert.Equal(t, &expected, res)

}

func TestClient_GetBGPError(t *testing.T) {
	setup()
	var client = &Client{APIEndpoint: server.URL, AuthToken: "foo"}
	mux.HandleFunc("/tests/bgp/1.json", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)
		w.WriteHeader(http.StatusBadRequest)
	})

	_, err := client.GetBGP(1)
	teardown()
	assert.Error(t, err)
}

func TestClient_GetBGPStatusCode(t *testing.T) {
	setup()
	out := `{"test":[{"testId":1,"testName":"test123","type":"bgp"}]}`
	var client = &Client{APIEndpoint: server.URL, AuthToken: "foo"}
	mux.HandleFunc("/tests/1.json", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(out))
	})

	_, err := client.GetBGP(1)
	teardown()
	assert.ErrorContains(t, err, "Response did not contain formatted error: %!s(<nil>). HTTP response code: 400")
}

func TestClient_CreateBGPStatusCode(t *testing.T) {
	setup()
	var client = &Client{APIEndpoint: server.URL, AuthToken: "foo"}
	mux.HandleFunc("/tests/bgp/new.json", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "POST", r.Method)
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`{}`))
	})
	_, err := client.CreateBGP(BGP{})
	teardown()
	assert.ErrorContains(t, err, "Response did not contain formatted error: %!s(<nil>). HTTP response code: 400")
}

func TestClient_UpdateBGPStatusCode(t *testing.T) {
	setup()
	var client = &Client{APIEndpoint: server.URL, AuthToken: "foo"}
	mux.HandleFunc("/tests/bgp/1/update.json", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "POST", r.Method)
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`{}`))
	})
	_, err := client.UpdateBGP(1, BGP{})
	teardown()
	assert.ErrorContains(t, err, "Response did not contain formatted error: %!s(<nil>). HTTP response code: 400")
}

func TestClient_DeleteBGPStatusCode(t *testing.T) {
	setup()
	var client = &Client{APIEndpoint: server.URL, AuthToken: "foo"}
	mux.HandleFunc("/tests/bgp/1/delete.json", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "POST", r.Method)
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`{}`))
	})
	err := client.DeleteBGP(1)
	teardown()
	assert.ErrorContains(t, err, "Response did not contain formatted error: %!s(<nil>). HTTP response code: 400")
}
