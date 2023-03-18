package handler

import (
	"fmt"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `
<!DOCTYPE html>
<html lang="en">
	<head>
		<meta charset="UTF-8" />
		<meta http-equiv="X-UA-Compatible" content="IE=edge" />
		<meta name="viewport" content="width=device-width, initial-scale=1.0" />
		<title>Dan Uhl Profile</title>
	</head>
	<body>
		<style>
			body {
				background: #333;
				color: #f0f0f0;
			}
			main {
				display: flex;
				justify-content: center;
			}
			main li a {
				color: cadetblue;
			}
			main li a:visited {
				color: blanchedalmond;
			}
		</style>

		<main>
			<article>
				<h1>Dan Uhl (duhl) Profile Page</h1>

				<p>Here's my stuff:</p>
				<ul>
					<li><a href="https://blog.reyan.me">Blog</a></li>
					<li>
						<a
							href="https://podcasts.apple.com/gb/podcast/luculent-lollygag/id1608573971"
							>Podcast</a
						>
					</li>
					<li><a href="https://github.com/danieluhl">Github</a></li>
					<li>
						<a href="https://discordapp.com/users/typing%20turtle#0978"
							>Discord</a
						>
					</li>
					<li>
						<a href="https://www.youtube.com/@typingturtle5155/featured"
							>Youtube</a
						>
					</li>
					<li><a href="https://www.twitch.tv/typing_turtle">Twitch</a></li>
					<li><a href="danielruhl@gmail.com">Email</a></li>
					<li><a href="https://twitter.com/danieluhl">Twitter</a></li>
					<li><a href="https://www.linkedin.com/in/danieluhl/">LinkedIn</a></li>
				</ul>
			</article>
		</main>
	</body>
</html>
	`)
}

// package handler

// import (
// 	"html/template"
// 	"log"
// 	"net/http"
// )

// func viewHandler(w http.ResponseWriter, r *http.Request) {
// 	renderTemplate(w)
// }

// var templates = template.Must(template.ParseFiles("profile.html"))

// func renderTemplate(w http.ResponseWriter) {
// 	err := templates.ExecuteTemplate(w, "profile.html", "Hello World!")
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 	}
// }

// func Handler() {
// 	http.HandleFunc("/", viewHandler)
// 	log.Fatal(http.ListenAndServe(":8080", nil))
// }
