package main

import (
	"net/http"
)

// RenderApolloSandbox serves the Apollo Sandbox UI
func RenderApolloSandbox(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte(`<!DOCTYPE html>
<html>
<head>
  <meta charset="utf-8">
  <meta name="viewport" content="width=device-width, initial-scale=1">
  <title>Apollo Sandbox</title>
  <style>
    body { margin: 0; padding: 0; overflow: hidden; }
    #sandbox { width: 100%; height: 100vh; }
  </style>
</head>
<body>
  <div id="sandbox"></div>

  <script src="https://embeddable-sandbox.cdn.apollographql.com/_latest/embeddable-sandbox.umd.production.min.js"></script>
  <script>
    const sandbox = new window.EmbeddedSandbox({
      target: '#sandbox',
      // Ensure it uses relative URL to avoid CORS issues
      initialEndpoint: '/graphql',
      includeCookies: true,
      initialState: {
        // Enable introspection explicitly in the Explorer
        explorer: {
          // Enable schema introspection features
          introspection: { 
            enable: true
          }
        }
      }
    });
  </script>
</body>
</html>`))
}