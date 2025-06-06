package pages

import (
	"time"

	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/components"
	. "maragu.dev/gomponents/html"
)

func Home() Node {
	return HTML5(HTML5Props{
		Title:    "go components",
		Language: "en",
		Head: []Node{
			Meta(Charset("utf-8")),
			Meta(Name("viewport"), Content("width=device-width, initial-scale=1")),
			Link(Rel("icon"), Href("https://aqclf.xyz/favicon.svg")),
			Script(Src("https://cdn.tailwindcss.com")),
		},
		Body: []Node{Class("bg-pink-200 text-pink-900 min-h-screen flex flex-col items-center justify-center p-4"),
			Div(Class("space-y-6 text-center max-w-md"),
				H1(Class("text-3xl font-bold"), Text("go components are fun?")),
				P(Class("text-sm text-gray-600"), Text("they are actually pretty cool")),

				Div(Class("mt-4 p-4 border border-pink-900"),
					P(Class("text-sm uppercase tracking-wide"), Text("server time")),
					P(Class("text-xl font-mono mt-1 time"), Text(time.Now().Format("15:04:05"))),
					Script(Raw(`
const timeEl = document.querySelector('.time')
let [h, m, s] = timeEl.textContent.split(':').map(Number)
let serverClock = new Date()
serverClock.setHours(h, m, s, 0)

setInterval(() => {
	serverClock.setSeconds(serverClock.getSeconds() + 1)
	const hh = String(serverClock.getHours()).padStart(2, '0')
	const mm = String(serverClock.getMinutes()).padStart(2, '0')
	const ss = String(serverClock.getSeconds()).padStart(2, '0')
	timeEl.textContent = hh + ':' + mm + ':' + ss
}, 1000)
				  `)),
				),
			),
		},
	})
}
