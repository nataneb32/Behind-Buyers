package observer

type Listener interface {
	Notify(report map[string]interface{})
}

type ReportObserver struct {
	analytics []Listener
}

func (r *ReportObserver) Subscribe(a Listener) {
	r.analytics = append(r.analytics, a)
}

func (r *ReportObserver) Report(report map[string]interface{}) {
	for _, analytics := range r.analytics {
		analytics.Notify(report)
	}
}

func CreateReportObserver() *ReportObserver {
	return &ReportObserver{
		analytics: []Listener{},
	}
}
