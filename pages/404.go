package pages

import (
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/components"
	. "maragu.dev/gomponents/html"
)

func NotFound() Node {
	return HTML5(HTML5Props{
		Title:    "404",
		Language: "en",
		Head: []Node{
			Meta(Charset("utf-8")),
			Meta(Name("viewport"), Content("width=device-width, initial-scale=1")),
			Link(Rel("icon"), Href("https://aqclf.xyz/favicon.svg")),
			Script(Src("https://cdn.tailwindcss.com")),
		},
		Body: []Node{Class("bg-pink-200 text-pink-900 min-h-screen flex flex-col items-center justify-center p-4"),
			Div(Class("space-y-6 text-center max-w-md"),
				H1(Class("text-3xl font-bold"), Text("404")),
				P(Class("text-sm text-gray-600"), Text("page not found")),
			),
		},
	})
}
