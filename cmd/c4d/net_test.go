package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
	"time"

	"github.com/gorilla/websocket"

	"github.com/cheekybits/is"
)

func TestHTTPResponse(t *testing.T) {
	is := is.New(t)
	handler := func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "something failed", http.StatusInternalServerError)
	}

	req, err := http.NewRequest("GET", "http://example.com/foo", nil)
	is.NoErr(err)

	w := httptest.NewRecorder()
	handler(w, req)
	is.Equal(w.Code, 500)
	is.Equal(w.Body.String(), "something failed\n")
}

func TestHTTPindex(t *testing.T) {
	is := is.New(t)
	req, err := http.NewRequest("GET", "http://example.com/foo", nil)
	is.NoErr(err)

	w := httptest.NewRecorder()
	Index(w, req)
	is.Equal(w.Code, 200)
	is.Equal(w.HeaderMap["Content-Type"][0], "application/json;charset=UTF-8")
	var v Assets
	err = json.Unmarshal(w.Body.Bytes(), &v)
	is.NoErr(err)
	expected := Assets{
		Asset{Name: "Some asset data"},
		Asset{Name: "Some other asset data"},
	}

	for i, t := range expected {
		is.Equal(v[i].Name, t.Name)
	}

}

func TestHTTPassetindex(t *testing.T) {
	is := is.New(t)
	req, err := http.NewRequest("GET", "http://example.com/foo", nil)
	is.NoErr(err)

	w := httptest.NewRecorder()
	AssetsIndex(w, req)
	is.Equal(w.Code, 200)
	var v Assets
	err = json.Unmarshal(w.Body.Bytes(), &v)
	is.NoErr(err)
	expected := Assets{
		Asset{Name: "Some asset data"},
		Asset{Name: "Some other asset data"},
	}

	for i, t := range expected {
		is.Equal(v[i].Name, t.Name)
	}
}

func TestHTTPassetshow(t *testing.T) {
	is := is.New(t)
	req, err := http.NewRequest("GET", "http://example.com/foo", nil)
	is.NoErr(err)

	w := httptest.NewRecorder()
	AssetShow(w, req)
	is.Equal(w.Code, 200)

	// fmt.Printf("%d - %s", w.Code, w.Body.String())
}

func TestHTTPassetcreate(t *testing.T) {
	is := is.New(t)

	// router := NewRouter()
	ts := httptest.NewServer(http.HandlerFunc(AssetCreate))
	defer ts.Close()

	jsonStr := []byte(`{"name":"New Asset"}`)
	req, err := http.NewRequest("POST", ts.URL, bytes.NewBuffer(jsonStr))
	req.Header.Set("X-Custom-Header", "myvalue")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	is.NoErr(err)
	defer resp.Body.Close()
	is.Equal(resp.StatusCode, 201)
	is.Equal(resp.Header["Content-Type"][0], "application/json; charset=UTF-8")
	data := make([]byte, 100)
	n, err := resp.Body.Read(data)
	is.Equal(err, io.EOF)
	is.Equal(n, 28)
	data = data[:n]
	var v Asset
	err = json.Unmarshal(data, &v)
	is.NoErr(err)
	exptd := Asset{
		Id:   3,
		Name: "Write presentation",
	}
	is.Equal(v.Id, exptd.Id)
}

func TestWebsocket(t *testing.T) {
	is := is.New(t)
	// router := NewRouter()
	ts := httptest.NewServer(http.HandlerFunc(AssetWebsocket))
	defer ts.Close()
	turl, err := url.Parse(ts.URL)
	is.NoErr(err)
	turl.Scheme = "ws"
	wsurl := turl
	t.Log(wsurl)
	c, _, err := websocket.DefaultDialer.Dial(wsurl.String(), nil)
	is.NoErr(err)
	defer c.Close()

	done := make(chan struct{})
	go func() {
		defer c.Close()
		defer close(done)
		for {
			_, message, err := c.ReadMessage()
			if err != nil {
				fmt.Println("read:", err)
				return
			}
			fmt.Printf("recv: %s", message)
		}
	}()

	ticker := time.NewTicker(time.Millisecond * 100)
	defer ticker.Stop()

	interrupt := time.After(time.Second * 3)
	for {
		select {
		case t := <-ticker.C:
			err := c.WriteMessage(websocket.TextMessage, []byte(t.String()))
			if err != nil {
				fmt.Println("write:", err)
				return
			}
		case <-interrupt:
			fmt.Println("interrupt")
			// To cleanly close a connection, a client should send a close
			// frame and wait for the server to close the connection.
			err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				fmt.Println("write close:", err)
				return
			}
			select {
			case <-done:
			case <-time.After(time.Second):
			}
			c.Close()
			return
		}
	}

	// err = c.WriteMessage(websocket.TextMessage, []byte("Step 1"))
	// is.NoErr(err)
	// err = c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
	// is.NoErr(err)

}
