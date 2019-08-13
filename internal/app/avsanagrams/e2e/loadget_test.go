package e2e

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"avsanagrams/internal/app/avsanagrams/api"
	"avsanagrams/internal/pkg/storage"

	"github.com/stretchr/testify/require"
)

func handlers() http.Handler {
	r := http.NewServeMux()

	store := storage.New()
	apiHandlers := api.New(store)

	r.HandleFunc("/load", apiHandlers.Load)
	r.HandleFunc("/get", apiHandlers.Get)

	return r
}

func Test_LoadGet(t *testing.T) {
	srv := httptest.NewServer(handlers())
	defer srv.Close()

	reqBody := []byte(`["foobar", "aabb", "baba", "boofar", "test"]`)
	reqBodyBuff := bytes.NewBuffer(reqBody)

	res, err := http.Post(fmt.Sprintf("%s/load", srv.URL), "application/json", reqBodyBuff)
	defer res.Body.Close()

	require.NoError(t, err)
	require.Equal(t, http.StatusOK, res.StatusCode)

	tt := []struct {
		word string
		resp []byte
	}{
		{
			word: "foobar",
			resp: []byte(`["foobar","boofar"]`),
		}, {
			word: "raboof",
			resp: []byte(`["foobar","boofar"]`),
		}, {
			word: "abba",
			resp: []byte(`["aabb","baba"]`),
		}, {
			word: "test",
			resp: []byte(`["test"]`),
		}, {
			word: "qwerty",
			resp: []byte(`null`),
		},
	}

	for i, test := range tt {
		t.Run(fmt.Sprintf("step %d", i), func(t *testing.T) {
			res, err := http.Get(fmt.Sprintf("%s/get?word=%s", srv.URL, test.word))
			defer res.Body.Close()
			require.NoError(t, err)

			expected := test.resp

			actual, err := ioutil.ReadAll(res.Body)
			require.NoError(t, err)

			require.Equal(t, expected, actual)
		})
	}
}
