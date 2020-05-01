package observer

type Listener interface {
	Notify(field string, value interface{})
}

type ReportObserver struct {
	analytics []Listener
}

func (r *ReportObserver) Subscribe(a Listener) {
	r.analytics = append(r.analytics, a)
}

func (r *ReportObserver) Report(report map[string]interface{}) {
	for field, value := range report {
		for _, analytics := range r.analytics {
			analytics.Notify(field, value)
		}
	}
}

func CreateReportObserver() *ReportObserver {
	return &ReportObserver{
		analytics: []Listener{},
	}
}
