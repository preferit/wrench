package wrench

import (
	. "github.com/gregoryv/web"
	"github.com/gregoryv/web/toc"
)

func NewHelpView() *HelpView {
	return &HelpView{}
}

type HelpView struct{}

func (me *HelpView) Render() *Page {
	navigation := Nav()
	content := Div(
		H1("Help"),
		navigation,

		Section(
			H2("Timesheet file format"),
			P("Timesheets are plain text and are specific to year and month"),
			Pre(Class("timesheet"),
				`2015 June
---------
23  1 Mon 8
    2 Tue 8
    3 Wed 8 (3 meeting)
    4 Thu 8
    5 Fri 6 Ended work 2 hours early, felt sick.
    6 Sat
    7 Sun
24  8 Mon 8
    9 Tue 8
   10 Wed 8
   11 Thu 8 (7 conference) (1 travel)
   12 Fri 8
   13 Sat
   14 Sun
25 15 Mon 8
   16 Tue 8
   17 Wed 8:30
   18 Thu 8
   19 Fri 8
   20 Sat
   21 Sun
26 22 Mon 8
   23 Tue 8
   24 Wed 8
   25 Thu 8
   26 Fri 8
   27 Sat
   28 Sun
27 29 Mon 8
   30 Tue 8`,
			),
		),
	)

	toc.GenerateIDs(content, "h2", "h3")
	navigation.With(toc.ParseTOC(content, "h2"))

	return NewPage(
		"help.html",
		Html(
			Head(
				Style(theme()),
			),
			Body(content, footer()),
		),
	)
}
