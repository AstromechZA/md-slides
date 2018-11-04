package css

import (
	"html/template"
	"log"
)

func AddCommonStyleTemplate(root *template.Template) {
	if _, err := root.New("style.common").Parse(`
html {
	font-size: 18px;
}

body {
	background: #444444;
}

.slide-wrap {
	align-self: center;
	display: grid;
	box-sizing: border-box;
	background: #fffff8;
	padding: 1rem;
    border-radius: 0.3rem;
	box-shadow: 0px 0.2rem 0.6rem black;
	padding-left: 3rem;
    padding-right: 3rem;
	position: absolute;
	overflow: hidden;
	grid-auto-columns: 1fr;
	grid-auto-rows: 1fr;
}

.slide-wrap-halign-left {justify-items: start;}
.slide-wrap-halign-center {justify-items: center;}
.slide-wrap-halign-right {justify-items: end;}
.slide-wrap-valign-top {align-items: start;}
.slide-wrap-valign-center {align-items: center;}
.slide-wrap-valign-bottom {align-items: end;}
.slide-wrap-talign-left {text-align: left;}
.slide-wrap-talign-center {text-align: center;}
.slide-wrap-talign-right {text-align: right;}

.page-number {
	font-family: Palatino, "Palatino Linotype", "Palatino LT STD", "Book Antiqua", Georgia, serif;
	position: absolute;
	bottom: 0;
	right: 0;
	margin: 0.5rem;
	font-size: 0.9em;
	color: darkgrey;
}

.page-footer {
	font-family: Palatino, "Palatino Linotype", "Palatino LT STD", "Book Antiqua", Georgia, serif;
	position: absolute;
	bottom: 0;
	left: 0;
	margin: 0.5rem;
	font-size: 0.9em;
	color: darkgrey;
}

pre.chroma {
	padding: 1rem;
	border-radius: 0.5rem;
	border: 1px solid lightgrey;
	white-space: pre-wrap;
	background: white;
}
`); err != nil {
		log.Fatalf("failed to parse: %s", err)
	}
}
