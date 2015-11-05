# 1.12

Modify the Lissajous server to read parameter values from the URL. For exam- ple, you might arrange it so that a URL like `http://localhost:8000/?cycles=20` sets the number of cycles to 20 instead of the default 5. 

See the `render` handler, which parses the URL query string for `cycles` and
`size` parameters. 

For example, a request for `http://localhost:8000/lissajous?cycles=5&size=300`
would generate a lissajous animation with five cycles and an image canvas of
600 pixels (-300 ... +300).
