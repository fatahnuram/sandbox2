package sandbox2

import (
	"encoding/json"
	"io"
	"net/http"
	"path"
	"sandbox2/middleware"
	"strings"
)

type CustomError struct {
	Status string `json:"status"`
	Error  string `json:"error"`
}

type Message struct {
	Status string `json:"status"`
	Msg    string `json:"msg"`
}

func ShiftPath(p string) (string, string) {
	p = path.Clean("/" + p)
	i := strings.Index(p[1:], "/") + 1
	if i <= 0 {
		return p[1:], "/"
	}
	return p[1:i], p[i:]
}

func Respond(w http.ResponseWriter, r *http.Request, status int, data interface{}) error {
	// wrap error as json
	if e, ok := data.(error); ok {
		tmp := CustomError{}
		tmp.Status = "error"
		tmp.Error = e.Error()
		data = tmp
	}

	// wrap one-line msg as json
	if s, ok := data.(string); ok {
		tmp := Message{
			Status: "ok",
			Msg:    s,
		}
		data = tmp
	}

	resp, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(resp)

	middleware.LogRequest(r, status)
	return nil
}

func ParseBody(body io.ReadCloser, result interface{}) error {
	defer body.Close()
	decoder := json.NewDecoder(body)
	return decoder.Decode(result)
}
