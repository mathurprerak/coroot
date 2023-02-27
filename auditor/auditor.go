package auditor

import (
	"github.com/coroot/coroot/model"
	"sort"
)

type appAuditor struct {
	w       *model.World
	app     *model.Application
	reports []*model.AuditReport
}

func Audit(w *model.World) {
	for _, app := range w.Applications {
		a := &appAuditor{
			w:   w,
			app: app,
		}
		a.slo()
		a.instances()
		a.cpu()
		a.memory()
		a.storage()
		a.network()
		a.postgres()
		a.redis()
		a.jvm()
		a.logs()
		a.deployments()

		for _, r := range a.reports {
			widgets := enrichWidgets(r.Widgets, app.Events)
			sort.SliceStable(widgets, func(i, j int) bool {
				return widgets[i].Table != nil
			})
			r.Widgets = widgets

			for _, ch := range r.Checks {
				ch.Calc()
				if ch.Status > r.Status {
					r.Status = ch.Status
				}
			}
			switch r.Name {
			case model.AuditReportPostgres, model.AuditReportRedis, model.AuditReportInstances, model.AuditReportSLO:
				if app.Status < r.Status {
					app.Status = r.Status
				}
			}
			app.Reports = append(app.Reports, r)
		}
	}
}

func (a *appAuditor) addReport(name model.AuditReportName) *model.AuditReport {
	r := model.NewAuditReport(a.app, a.w.Ctx, a.w.CheckConfigs, name)
	a.reports = append(a.reports, r)
	return r
}

func enrichWidgets(widgets []*model.Widget, events []*model.ApplicationEvent) []*model.Widget {
	var res []*model.Widget
	for _, w := range widgets {
		if w.Chart != nil {
			if len(w.Chart.Series) == 0 {
				continue
			}
			w.Chart.AddEventsAnnotations(events)
		}
		if w.ChartGroup != nil {
			var charts []*model.Chart
			for _, ch := range w.ChartGroup.Charts {
				if len(ch.Series) == 0 {
					continue
				}
				charts = append(charts, ch)
				ch.AddEventsAnnotations(events)
			}
			if len(charts) == 0 {
				continue
			}
			w.ChartGroup.Charts = charts
			w.ChartGroup.AutoFeatureChart()
		}
		if w.LogPatterns != nil {
			for _, p := range w.LogPatterns.Patterns {
				if p.Instances != nil {
					p.Instances.AddEventsAnnotations(events)
				}
			}
		}
		res = append(res, w)
	}
	return res
}
