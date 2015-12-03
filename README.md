# binloader
A HTTP server which builds Go binaries for a variety of platforms.

## Usage

binloader is used in our browser.

To download a Go command `github.com/codegangsta/gin`, you would GET `http://<GOARCH>.<GOOS>.<BINLOADER_HOST>/github.com/codegangsta/gin`.

If you are hosting binloader locally or do not want to use hostnames, this query works instead ``http://<BINLOADER_HOST>/github.com/codegangsta/gin/?goarch=<GOARCH>&goos=<GOOS>`
