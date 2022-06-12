# Building a Basic Web Application in Go

### Highlights of (9<sup>th</sup>) commit

- Passing dynamic data from handlers to templates or page

- To solve <span style="color:red">import cycle</span> issue:

  - make models package

- <u><em>models</em></u> pkg will store models like template data, database models.

- always pass data types/structs using <u><b>pointers</b></u> (in most cases)

### Next Commit (10<sup>th</sup>)

Make a separate file for routes, routes.go

<u>mux</u> is a <u>http.Handler</u>

Read About mux and http server in Go [here](https://www.digitalocean.com/community/tutorials/how-to-make-an-http-server-in-go)

Some third party routers that can also be used:

- https://github.com/bmizerany/pat
- https://github.com/go-chi/chi (chi router has some built-in middleware too)

### Commit 11<sup>th</sup>

- Chi & Middleware

Using chi third party router for routing.

Middleware allows us to process requests as they come into your Web Aplication and perform some action on it.

How to use middleware? check chi template github (simple enough)

Remaining last: build your <u>middleware</u> (CSRF Token, NoSurf), <u>sessions</u>

### Commit

https://stackoverflow.com/questions/35888817/load-image-and-css-in-golang
