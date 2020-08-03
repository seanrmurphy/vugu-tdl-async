module github.com/vugu-examples/simple

go 1.14

replace github.com/seanrmurphy/vugu-tdl-async/backend/lambda/types => ../backend/lambda/types

require (
	github.com/go-openapi/runtime v0.19.20 // indirect
	github.com/go-openapi/strfmt v0.19.5
	github.com/google/uuid v1.1.1
	github.com/gorilla/websocket v1.4.2
	github.com/nirasan/go-oauth-pkce-code-verifier v0.0.0-20170819232839-0fbfe93532da
	github.com/seanrmurphy/vugu-tdl-async/backend/lambda/types v0.0.0-00010101000000-000000000000
	github.com/seanrmurphy/vugu-tdl-swagger v0.0.0-20200722145255-ebaa3fbe6d13
	github.com/vugu/vgrouter v0.0.0-20200725205318-eeb478c42e5d
	github.com/vugu/vjson v0.0.0-20200505061711-f9cbed27d3d9
	github.com/vugu/vugu v0.3.2
	nhooyr.io/websocket v1.8.6
)
