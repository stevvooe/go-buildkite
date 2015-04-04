package buildkite

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestAgentsService_List(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/organizations/my-great-org/agents", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `[{"id":"123"},{"id":"1234"}]`)
	})

	agents, _, err := client.Agents.List("my-great-org", nil)
	if err != nil {
		t.Errorf("Agents.List returned error: %v", err)
	}

	want := []Agent{{ID: "123"}, {ID: "1234"}}
	if !reflect.DeepEqual(agents, want) {
		t.Errorf("Agents.List returned %+v, want %+v", agents, want)
	}
}

func TestAgentsService_Get(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/organizations/my-great-org/agents/123", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `{"id":"123"}`)
	})

	agent, _, err := client.Agents.Get("my-great-org", "123")
	if err != nil {
		t.Errorf("Agents.List returned error: %v", err)
	}

	want := &Agent{ID: "123"}
	if !reflect.DeepEqual(agent, want) {
		t.Errorf("Agents.List returned %+v, want %+v", agent, want)
	}
}

func TestAgentsService_Create(t *testing.T) {
	setup()
	defer teardown()

	params := map[string]string{
		"name": "new_agent_bob",
	}

	mux.HandleFunc("/v1/organizations/my-great-org/agents", func(w http.ResponseWriter, r *http.Request) {
		v := make(map[string]string)
		json.NewDecoder(r.Body).Decode(&v)

		testMethod(t, r, "POST")

		if !reflect.DeepEqual(v, params) {
			t.Errorf("Request body = %+v, want %+v", v, params)
		}

		fmt.Fprint(w, `{"id":"123"}`)
	})

	agent, _, err := client.Agents.Create("my-great-org", "new_agent_bob")
	if err != nil {
		t.Errorf("Agents.Create returned error: %v", err)
	}

	want := &Agent{ID: "123"}
	if !reflect.DeepEqual(agent, want) {
		t.Errorf("Agents.Create returned %+v, want %+v", agent, want)
	}

}
