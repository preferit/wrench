package wrench

import (
	"time"

	. "github.com/gregoryv/web"
)

func NewReportsView(acc string) *ReportsView {
	return &ReportsView{
		Account: acc,
		Reports: []Report{},
	}
}

type ReportsView struct {
	Account string
	Reports []Report
}

type Report struct {
	Text string
}

func (me *ReportsView) Render() *Page {
	content := Div(
		H1("Reports"),
		"Logged in as: ", B(me.Account),
	)

	for _, report := range me.Reports {
		content.With(Pre(report.Text))
	}

	return NewPage("",
		Html(
			Head(Style(theme())),
			Body(content, footer()),
		),
	)
}

func NewIndexView() *IndexView {
	return &IndexView{}
}

type IndexView struct{}

func (me *IndexView) Render() *Page {
	content := Div(
		H1("Wrench"),
		A(Href("reports/"), "Reports"),
	)
	return NewPage(
		"index.html",
		Html(
			Head(
				Style(theme()),
			),
			Body(content, footer()),
		),
	)
}

func theme() *CSS {
	css := NewCSS()
	css.Style("html, body",
		"margin: 0 0",
		"padding: 0 0",
		"background-color: #e2e2e2",
	)
	css.Style("div",
		"background-color: white",
		"padding: 1em 1em 2em 1em",
		"min-height: 300",
	)
	css.Style("section",
		"margin-bottom: 5em",
	)
	css.Style("pre",
		"margin-left: 1em",
	)
	css.Style("footer",
		"border-top: 1px solid #727272",
		"padding: 0.6em 0.6em",
	)
	css.Style(".timesheet",
		"border: 1px #e2e2e2 dotted",
		"padding: 1em 1em",
		"background-color: #ffffe6",
	)
	return css
}

// When the service started so we know the uptime
var start = time.Now()

func footer() *Element {
	return Footer(
		"Uptime: ",
		time.Since(start).Round(time.Second).String(),
	)
}
