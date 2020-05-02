package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"time"

	"../../../pkg/access"

	"../../../pkg/report/observer"
	"../../../pkg/storage"
)

type Handler struct {
	reportObserver *observer.ReportObserver
	access         *access.Listener
	accessPlotter  *access.Plotter
	s              storage.Storage
}

// It's the handlers for register a new report
func (h *Handler) RegisterReport(w http.ResponseWriter, r *http.Request) {
	var buffer bytes.Buffer
	buffer.ReadFrom(r.Body)

	jsonReport := make(map[string]interface{})

	json.Unmarshal(buffer.Bytes(), &jsonReport)

	h.reportObserver.Report(jsonReport)

	response, _ := json.Marshal(h.access.GetData())
	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

func NewHandler() *Handler {
	s := storage.NewInMemory()

	h := &Handler{
		reportObserver: observer.CreateReportObserver(),
		access:         access.CreateListener(s),
		accessPlotter:  access.NewPlotter(s),
		s:              s,
	}

	h.reportObserver.Subscribe(h.access)

	return h
}

func (h *Handler) Index(w http.ResponseWriter, r *http.Request) {
	chart := h.accessPlotter.PlotChartAccessOfPage("test", 0, int(time.Now().UnixNano()/1e6), 1000*60*60)
	response, _ := json.Marshal(chart)
	w.Write(response)
}

func (h *Handler) IndexT(w http.ResponseWriter, r *http.Request) {
	chart := h.accessPlotter.PlotChartAccessOfPage("test", 0, int(time.Now().UnixNano()/1e6), 1000*60)
	t, err := template.New("index.html").ParseFiles("./templates/index.gohtml")
	if err != nil {
		fmt.Println(err, t)
		return
	}
	_ = t.ExecuteTemplate(w, "chart", chart)

}
