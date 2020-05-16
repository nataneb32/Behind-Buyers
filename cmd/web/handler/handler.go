package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"../../../pkg/timespend"

	"../../../pkg/access"
	"../../../pkg/click"

	"../../../pkg/report/observer"
	"../../../pkg/storage"
)

type Handler struct {
	reportObserver *observer.ReportObserver
	access         *access.Listener
	timespend      *timespend.Listener
	click          *click.Listener
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

type AccessChartRequest struct {
	Page  string `json:"page"`
	From  int    `json:"from"`
	To    int    `json:"to"`
	Steps int    `json:"steps"`
}

func (h *Handler) AccessApiChart(w http.ResponseWriter, r *http.Request) {

	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	var req AccessChartRequest
	err = json.Unmarshal(b, &req)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	fmt.Println(req)

	chart := h.accessPlotter.PlotChartAccessOfPage(req.Page, int(req.From), int(req.To), int(req.Steps))
	fmt.Println(chart)
	output, _ := json.Marshal(chart)
	w.Write(output)
}

func NewHandler() *Handler {
	currentPath, err := os.Getwd()

	if err != nil {
		log.Fatal(err)
	}

	s := storage.NewFileStorage(filepath.Join(currentPath, "storage"))

	h := &Handler{
		reportObserver: observer.CreateReportObserver(),
		access:         access.CreateListener(s),
		timespend:      timespend.CreateListener(s),
		click:          click.CreateListener(s),
		accessPlotter:  access.NewPlotter(s),
		s:              s,
	}

	h.reportObserver.Subscribe(h.access)
	h.reportObserver.Subscribe(h.click)
	h.reportObserver.Subscribe(h.timespend)

	return h
}

func (h *Handler) Index(w http.ResponseWriter, r *http.Request) {
	chart := h.accessPlotter.PlotChartAccessOfPage("test", int(time.Now().UnixNano()/1e6)-60*1000*60, int(time.Now().UnixNano()/1e6), 1000*60)
	fmt.Println(chart)
	response, _ := json.Marshal(chart)
	w.Write(response)
}

func (h *Handler) IndexT(w http.ResponseWriter, r *http.Request) {
	chart := h.accessPlotter.PlotChartAccessOfPage("test", int(time.Now().UnixNano()/1e6)-60*1000*10, int(time.Now().UnixNano()/1e6), 1000*60)
	fmt.Println(chart)
	t, err := template.New("index.html").ParseFiles("./templates/index.gohtml")
	if err != nil {
		fmt.Println(err, t)
		return
	}
	_ = t.ExecuteTemplate(w, "chart", chart)

}

func (h *Handler) LoginTemplate(w http.ResponseWriter, r *http.Request) {
	chart := h.accessPlotter.PlotChartAccessOfPage("test", 0, int(time.Now().UnixNano()/1e6), 1000*60)
	t, err := template.New("login.html").ParseFiles("./templates/login.gohtml")
	if err != nil {
		fmt.Println(err, t)
		return
	}
	_ = t.ExecuteTemplate(w, "chart", chart)

}
