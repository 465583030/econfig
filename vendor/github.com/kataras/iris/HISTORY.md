# History

**How to upgrade**: remove your `$GOPATH/src/github.com/kataras/iris` folder, open your command-line and execute this command: `go get -u github.com/kataras/iris`.

## 4.2.6 -> 4.2.7

- **ADDED**: You are now able to use a raw fasthttp handler as the router instead of the default Iris' one. Example [here](https://github.com/iris-contrib/examples/blob/master/custom_fasthttp_router/main.go). But remember that I'm always recommending to use the Iris' default which supports subdomains, group of routes(parties), auto path correction and many other built'n features. This exists for specific users who told me that they need a feature like that inside Iris, we have no performance cost at all so that's ok to exists.

## 4.2.5 -> 4.2.6

- **CHANGE**: Updater (See 4.2.4 and 4.2.3) runs in its own goroutine now, unless the `iris.Config.CheckForUpdatesSync` is true.
- **ADDED**: To align with fasthttp server's configuration, iris has these new Server Configuration's fields, which allows you to set a type of rate limit:
```go
// Maximum number of concurrent client connections allowed per IP.
//
// By default unlimited number of concurrent connections
// may be established to the server from a single IP address.
MaxConnsPerIP int

// Maximum number of requests served per connection.
//
// The server closes connection after the last request.
// 'Connection: close' header is added to the last response.
//
// By default unlimited number of requests may be served per connection.
MaxRequestsPerConn int

// Usage: iris.ListenTo{iris.OptionServerListeningAddr(":8080"), iris.OptionServerMaxConnsPerIP(300)}
//    or: iris.ListenTo(iris.ServerConfiguration{ListeningAddr: ":8080", MaxConnsPerIP: 300, MaxRequestsPerConn:100})
// for an optional second server with a different port you can always use:
//        iris.AddServer(iris.ServerConfiguration{ListeningAddr: ":9090", MaxConnsPerIP: 300, MaxRequestsPerConn:100})
```

## 4.2.4 -> 4.2.5

- **ADDED**: `iris.CheckForUpdates(force bool)` which can run the updater(look 4.2.4) at runtime too, updater is tested and worked at dev machine.

## 4.2.3 -> 4.2.4

- **NEW Experimental feature**: Updater with a `CheckForUpdates` [configuration](https://github.com/kataras/iris/blob/master/configuration.go) field, as requested [here](https://github.com/kataras/iris/issues/401)
```go
// CheckForUpdates will try to search for newer version of Iris based on the https://github.com/kataras/iris/releases
// If a newer version found then the app will ask the he dev/user if want to update the 'x' version
// if 'y' is pressed then the updater will try to install the latest version
// the updater, will notify the dev/user that the update is finished and should restart the App manually.
// Notes:
// 1. Experimental feature
// 2. If setted to true, the app will have a little startup delay
// 3. If you as developer edited the $GOPATH/src/github/kataras or any other Iris' Go dependencies at the past
//    then the update process will fail.
//
// Usage: iris.Set(iris.OptionCheckForUpdates(true)) or
//        iris.Config.CheckForUpdates = true or
//        app := iris.New(iris.OptionCheckForUpdates(true))
// Default is false
CheckForUpdates bool
```

## 4.2.2 -> 4.2.3

- [Add IsAjax() convenience method](https://github.com/kataras/iris/issues/423)

## 4.2.1 -> 4.2.2

- Fix [sessiondb issue 416](https://github.com/kataras/iris/issues/416)

## 4.2.0 -> 4.2.1

- **CHANGE**: No front-end changes if you used the default response engines before. Response Engines to Serializers, `iris.ResponseEngine` `serializer.Serializer`, comes from `kataras/go-serializer` which is installed automatically when you upgrade iris with `-u` flag.

    - the repo "github.com/iris-contrib/response" is a clone of "github.com/kataras/go-serializer", to keep compatibility state. examples and gitbook updated to work with the last.

    - `iris.UseResponse(iris.ResponseEngine, ...string)func (string)` was used to register custom response engines, now you use: `iris.UseSerializer(key string, s serializer.Serializer)`.

    - `iris.ResponseString` same defintion but differnet name now: `iris.SerializeToString`

[Serializer examples](https://github.com/iris-contrib/examples/tree/master/serialize_engines) and [Book section](https://kataras.gitbooks.io/iris/content/serialize-engines.html) updated.


## 4.1.7 -> 4.2.0

- **ADDED**: `iris.TemplateSourceString(src string, binding interface{}) string` this will parse the src raw contents to the template engine and return the string result & `context.RenderTemplateSource(status int, src string, binding interface{}, options ...map[string]interface{}) error` this will parse the src raw contents to the template engine and render the result to the client, as requseted [here](https://github.com/kataras/iris/issues/409).


This version has 'breaking' changes if you were, directly, passing custom configuration to a custom iris instance before.
As the TODO2 I had to think and implement a way to make configuration even easier and more simple to use.

With last changes in place, Iris is using new, cross-framework, and more stable packages made by me(so don't worry things are working and will as you expect) to render [templates](https://github.com/kataras/go-template), manage [sessions](https://github.com/kataras/go-sesions) and [websockets](https://github.com/kataras/go-websocket). So the `/kataras/iris/config` is no longer need to be there, we don't have core packages inside iris which need these configuration to other package-folder than the main anymore(in order to avoid the import-cycle), new file `/kataras/iris/configuration.go` is created for the configuration, which lives inside the main package, means that now:

- **if you want to pass directly configuration to a new custom iris instance, you don't have to import the github.com/kataras/iris/config package**

Naming changes:

- `config.Iris` -> `iris.Configuration`, which is the parent/main configuration. Added: `TimeFormat` and `Other` (pass any dynamic custom, other options there)
- `config.Sessions` -> `iris.SessionsConfiguration`
- `config.Websocket` -> `iris.WebscoketConfiguration`
- `config.Server` -> `iris.ServerConfiguration`
- `config.Tester` -> `iris.TesterConfiguration`

All these changes wasn't made only to remove the `./config` folder but to make easier for you to pass the exact configuration field/option you need to edit at the top of the default configuration, without need to pass the whole Configuration object. **Attention**: old way, pass `iris.Configuration` directly, is still valid object to pass to the  `iris.New`, so don't be afraid for breaking change, the only thing you will need to edit is the names of the configuration you saw on the previous paragraph.

**Configuration Declaration**:

instead of old, but still valid to pass to the `iris.New`:
- `iris.New(iris.Configuration{Charset: "UTF-8", Sessions: iris.SessionsConfiguration{Cookie: "cookienameid"}})`
now you can just write this:
- `iris.New(iris.OptionCharset("UTF-8"), iris.OptionSessionsCookie("cookienameid"))`

`.New` **by configuration**
```go
import "github.com/kataras/iris"
//...
myConfig := iris.Configuration{Charset: "UTF-8", IsDevelopment:true, Sessions: iris.SessionsConfiguration{Cookie:"mycookie"}, Websocket: iris.WebsocketConfiguration{Endpoint: "/my_endpoint"}}
iris.New(myConfig)
```

`.New` **by options**

```go
import "github.com/kataras/iris"
//...
iris.New(iris.OptionCharset("UTF-8"), iris.OptionIsDevelopment(true), iris.OptionSessionsCookie("mycookie"), iris.OptionWebsocketEndpoint("/my_endpoint"))

// if you want to set configuration after the .New use the .Set:
iris.Set(iris.OptionDisableBanner(true))
```

**List** of all available options:
```go
// OptionDisablePathCorrection corrects and redirects the requested path to the registed path
// for example, if /home/ path is requested but no handler for this Route found,
// then the Router checks if /home handler exists, if yes,
// (permant)redirects the client to the correct path /home
//
// Default is false
OptionDisablePathCorrection(val bool)

// OptionDisablePathEscape when is false then its escapes the path, the named parameters (if any).
OptionDisablePathEscape(val bool)

// OptionDisableBanner outputs the iris banner at startup
//
// Default is false
OptionDisableBanner(val bool)

// OptionLoggerOut is the destination for output
//
// Default is os.Stdout
OptionLoggerOut(val io.Writer)

// OptionLoggerPreffix is the logger's prefix to write at beginning of each line
//
// Default is [IRIS]
OptionLoggerPreffix(val string)

// OptionProfilePath a the route path, set it to enable http pprof tool
// Default is empty, if you set it to a $path, these routes will handled:
OptionProfilePath(val string)

// OptionDisableTemplateEngines set to true to disable loading the default template engine (html/template) and disallow the use of iris.UseEngine
// Default is false
OptionDisableTemplateEngines(val bool)

// OptionIsDevelopment iris will act like a developer, for example
// If true then re-builds the templates on each request
// Default is false
OptionIsDevelopment(val bool)

// OptionTimeFormat time format for any kind of datetime parsing
OptionTimeFormat(val string)

// OptionCharset character encoding for various rendering
// used for templates and the rest of the responses
// Default is "UTF-8"
OptionCharset(val string)

// OptionGzip enables gzip compression on your Render actions, this includes any type of render, templates and pure/raw content
// If you don't want to enable it globaly, you could just use the third parameter on context.Render("myfileOrResponse", structBinding{}, iris.RenderOptions{"gzip": true})
// Default is false
OptionGzip(val bool)

// OptionOther are the custom, dynamic options, can be empty
// this fill used only by you to set any app's options you want
// for each of an Iris instance
OptionOther(val ...options.Options) //map[string]interface{}, options is github.com/kataras/go-options

// OptionSessionsCookie string, the session's client cookie name, for example: "qsessionid"
OptionSessionsCookie(val string)

// OptionSessionsDecodeCookie set it to true to decode the cookie key with base64 URLEncoding
// Defaults to false
OptionSessionsDecodeCookie(val bool)

// OptionSessionsExpires the duration of which the cookie must expires (created_time.Add(Expires)).
// If you want to delete the cookie when the browser closes, set it to -1 but in this case, the server side's session duration is up to GcDuration
//
// Default infinitive/unlimited life duration(0)
OptionSessionsExpires(val time.Duration)

// OptionSessionsCookieLength the length of the sessionid's cookie's value, let it to 0 if you don't want to change it
// Defaults to 32
OptionSessionsCookieLength(val int)

// OptionSessionsGcDuration every how much duration(GcDuration) the memory should be clear for unused cookies (GcDuration)
// for example: time.Duration(2)*time.Hour. it will check every 2 hours if cookie hasn't be used for 2 hours,
// deletes it from backend memory until the user comes back, then the session continue to work as it was
//
// Default 2 hours
OptionSessionsGcDuration(val time.Duration)

// OptionSessionsDisableSubdomainPersistence set it to true in order dissallow your q subdomains to have access to the session cookie
// defaults to false
OptionSessionsDisableSubdomainPersistence(val bool)

// OptionWebsocketWriteTimeout time allowed to write a message to the connection.
// Default value is 15 * time.Second
OptionWebsocketWriteTimeout(val time.Duration)

// OptionWebsocketPongTimeout allowed to read the next pong message from the connection
// Default value is 60 * time.Second
OptionWebsocketPongTimeout(val time.Duration)

// OptionWebsocketPingPeriod send ping messages to the connection with this period. Must be less than PongTimeout
// Default value is (PongTimeout * 9) / 10
OptionWebsocketPingPeriod(val time.Duration)

// OptionWebsocketMaxMessageSize max message size allowed from connection
// Default value is 1024
OptionWebsocketMaxMessageSize(val int64)

// OptionWebsocketBinaryMessages set it to true in order to denotes binary data messages instead of utf-8 text
// see https://github.com/kataras/iris/issues/387#issuecomment-243006022 for more
// defaults to false
OptionWebsocketBinaryMessages(val bool)

// OptionWebsocketEndpoint is the path which the websocket server will listen for clients/connections
// Default value is empty string, if you don't set it the Websocket server is disabled.
OptionWebsocketEndpoint(val string)

// OptionWebsocketReadBufferSize is the buffer size for the underline reader
OptionWebsocketReadBufferSize(val int)

// OptionWebsocketWriteBufferSize is the buffer size for the underline writer
OptionWebsocketWriteBufferSize(val int)

// OptionTesterListeningAddr is the virtual server's listening addr (host)
// Default is "iris-go.com:1993"
OptionTesterListeningAddr(val string)

// OptionTesterExplicitURL If true then the url (should) be prepended manually, useful when want to test subdomains
// Default is false
OptionTesterExplicitURL(val bool)

// OptionTesterDebug if true then debug messages from the httpexpect will be shown when a test runs
// Default is false
OptionTesterDebug(val bool)


```

Now, some of you maybe use more than one server inside their iris instance/app, so you used the `iris.AddServer(config.Server{})`, which now becomes `iris.AddServer(iris.ServerConfiguration{})`, ServerConfiguration has also (optional) options to pass there and to `iris.ListenTo(OptionServerListeningAddr("mydomain.com"))`:


```go
// examples:
iris.AddServer(iris.OptionServerCertFile("file.cert"),iris.OptionServerKeyFile("file.key"))
iris.ListenTo(iris.OptionServerReadBufferSize(42000))

// or, old way but still valid:
iris.AddServer(iris.ServerConfiguration{ListeningAddr: "mydomain.com", CertFile: "file.cert", KeyFile: "file.key"})
iris.ListenTo(iris.ServerConfiguration{ReadBufferSize:42000, ListeningAddr: "mydomain.com"})
```

**List** of all Server's options:

```go
OptionServerListeningAddr(val string)

OptionServerCertFile(val string)

OptionServerKeyFile(val string)

// AutoTLS enable to get certifications from the Letsencrypt
// when this configuration field is true, the CertFile & KeyFile are empty, no need to provide a key.
//
// example: https://github.com/iris-contrib/examples/blob/master/letsencyrpt/main.go
OptionServerAutoTLS(val bool)

// Mode this is for unix only
OptionServerMode(val os.FileMode)
// OptionServerMaxRequestBodySize Maximum request body size.
//
// The server rejects requests with bodies exceeding this limit.
//
// By default request body size is 8MB.
OptionServerMaxRequestBodySize(val int)

// Per-connection buffer size for requests' reading.
// This also limits the maximum header size.
//
// Increase this buffer if your clients send multi-KB RequestURIs
// and/or multi-KB headers (for example, BIG cookies).
//
// Default buffer size is used if not set.
OptionServerReadBufferSize(val int)

// Per-connection buffer size for responses' writing.
//
// Default buffer size is used if not set.
OptionServerWriteBufferSize(val int)

// Maximum duration for reading the full request (including body).
//
// This also limits the maximum duration for idle keep-alive
// connections.
//
// By default request read timeout is unlimited.
OptionServerReadTimeout(val time.Duration)

// Maximum duration for writing the full response (including body).
//
// By default response write timeout is unlimited.
OptionServerWriteTimeout(val time.Duration)

// RedirectTo, defaults to empty, set it in order to override the station's handler and redirect all requests to this address which is of form(HOST:PORT or :PORT)
//
// NOTE: the http status is 'StatusMovedPermanently', means one-time-redirect(the browser remembers the new addr and goes to the new address without need to request something from this server
// which means that if you want to change this address you have to clear your browser's cache in order this to be able to change to the new addr.
//
// example: https://github.com/iris-contrib/examples/tree/master/multiserver_listening2
OptionServerRedirectTo(val string)

// OptionServerVirtual If this server is not really listens to a real host, it mostly used in order to achieve testing without system modifications
OptionServerVirtual(val bool)

// OptionServerVListeningAddr, can be used for both virtual = true or false,
// if it's setted to not empty, then the server's Host() will return this addr instead of the ListeningAddr.
// server's Host() is used inside global template helper funcs
// set it when you are sure you know what it does.
//
// Default is empty ""
OptionServerVListeningAddr(val string)

// OptionServerVScheme if setted to not empty value then all template's helper funcs prepends that as the url scheme instead of the real scheme
// server's .Scheme returns VScheme if  not empty && differs from real scheme
//
// Default is empty ""
OptionServerVScheme(val string)

// OptionServerName the server's name, defaults to "iris".
// You're free to change it, but I will trust you to don't, this is the only setting whose somebody, like me, can see if iris web framework is used
OptionServerName(val string)

```    

View all configuration fields and options by navigating to the [kataras/iris/configuration.go source file](https://github.com/kataras/iris/blob/master/configuration.go)

[Book](https://kataras.gitbooks.io/iris/content/configuration.html) & [Examples](https://github.com/iris-contrib/examples) are updated (website docs will be updated soon).

## 4.1.6 -> 4.1.7

- **CHANGED**: Use of the standard `log.Logger` instead of the `iris-contrib/logger`(colorful logger), these changes are reflects some middleware, examples and plugins, I updated all of them, so don't worry.

So, [iris-contrib/middleware/logger](https://github.com/iris-contrib/middleware/tree/master/logger) will now NO need to pass other Logger instead, instead of: `iris.Use(logger.New(iris.Logger))` use -> `iris.Use(logger.New())` which will use the iris/instance's Logger.

- **ADDED**: `context.Framework()` which returns your Iris instance (typeof `*iris.Framework`), useful for the future(Iris will give you, soon, the ability to pass custom options inside an iris instance).

## 4.1.5 -> 4.1.6

- Align with [go-sessions](https://github.com/kataras/go-sessions), no front-end changes, however I think that the best time to make an upgrade to your local Iris is right now.

## 4.1.4 -> 4.1.5

- Remove unused Plugin's custom callbacks, if you still need them in your project use this instead: https://github.com/kataras/go-events

## 4.1.3 -> 4.1.4

Zero front-end changes. No real improvements, developers can ignore this update.

- Replace the iris sessions with a new cross-framework package, [go-sessions](https://github.com/kataras/go-sessions). Same front-end API, sessions examples are compatible, configuration of `kataras/iris/config/sessions.go` is compatible. `kataras/context.SessionStore` is now `kataras/go-sessions.Session` (normally you, as user, never used it before, because of automatically session getting by `context.Session()`)

- `GzipWriter` is taken, now, from the `kataras/go-fs` package which has improvements versus the previous implementation.


## 4.1.2 -> 4.1.3

Zero front-end changes. No real improvements, developers can ignore this update.

- Replace the template engines with a new cross-framework package, [go-template](https://github.com/kataras/go-websocket). Same front-end API, examples and iris-contrib/template are compatible.

## 4.1.1 -> 4.1.2

Zero front-end changes. No real improvements, developers can ignore this update.

- Replace the main and underline websocket implementation with [go-websocket](https://github.com/kataras/go-websocket). Note that we still need the [ris-contrib/websocket](https://github.com/iris-contrib/websocket) package.
- Replace the use of iris-contrib/errors with [go-errors](https://github.com/kataras/go-errors), which has more features

## 4.0.0 -> 4.1.1

- **NEW FEATURE**: Basic remote control through SSH, example [here](https://github.com/iris-contrib/examples/blob/master/ssh/main.go)
- **NEW FEATURE**: Optionally `OnError` foreach Party (by prefix, use it with your own risk), example [here](https://github.com/iris-contrib/examples/blob/master/httperrors/main.go#L37)
- **NEW**: `iris.Config.Sessions.CookieLength`, You're able to customize the length of each sessionid's cookie's value. Default (and previous' implementation) is 32.
- **FIX**: Websocket panic on non-websocket connection[*](https://github.com/kataras/iris/issues/367)
- **FIX**: Multi websocket servers client-side source route panic[*](https://github.com/kataras/iris/issues/365)
- Better gzip response managment


## 4.0.0-alpha.5 -> 4.0.0

- **Feature request has been implemented**: Add layout support for Pug/Jade, example [here](https://github.com/iris-contrib/examples/tree/master/template_engines/template_pug_2).
- **Feature request has been implemented**: Forcefully closing a Websocket connection, `WebsocketConnection.Disconnect() error`.

- **FIX**: WebsocketConnection.Leave() will hang websocket server if .Leave was called manually when the websocket connection has been closed.
- **FIX**: StaticWeb not serving index.html correctly, align the func with the rest of Static funcs also, [example](https://github.com/iris-contrib/examples/tree/master/static_web) added.

Notes: if you compare it with previous releases (13+ versions before v3 stable), the v4 stable release was fast, now we had only 6 versions before stable, that was happened because many of bugs have been already fixed and we hadn't new bug reports and secondly, and most important for me, some third-party features are implemented mostly by third-party packages via other developers!


## 4.0.0-alpha.4 -> 4.0.0-alpha.5

- **NEW FEATURE**: Letsencrypt.org integration[*](https://github.com/kataras/iris/issues/220)
   - example [here](https://github.com/iris-contrib/examples/blob/master/letsencrypt/main.go)
- **FIX**: (ListenUNIX adds :80 to filename)[https://github.com/kataras/iris/issues/321]
- **FIX**: (Go-Bindata + ctx.Render)[https://github.com/kataras/iris/issues/315]
- **FIX** (auto-gzip doesn't really compress data in latest code)[https://github.com/kataras/iris/issues/312]


## 4.0.0-alpha.3 -> 4.0.0-alpha.4


**The important** , is that the [book](https://kataras.gitbooks.io/iris/content/) is finally updated!

If you're **willing to donate** click [here](DONATIONS.md)!


- `iris.Config.Gzip`, enables gzip compression on your Render actions, this includes any type of render, templates and pure/raw content. If you don't want to enable it globaly, you could just use the third parameter on context.Render("myfileOrResponse", structBinding{}, iris.RenderOptions{"gzip": true}). It defaults to false


-  **Added** `config.Server.Name` as requested


**Fix**
- https://github.com/kataras/iris/issues/301

**Sessions changes **

- `iris.Config.Sessions.Expires` it was time.Time, changed to time.Duration, which defaults to 0, means unlimited session life duration, if you change it then the correct date is setted on client's cookie but also server destroys the session automatically when the duration passed, this is better approach, see [here](https://github.com/kataras/iris/issues/301)


## 4.0.0-alpha.2 -> 4.0.0-alpha.3

**New**

A **Response Engine** gives you the freedom to create/change the render/response writer for

- `context.JSON`
- `context.JSONP`
- `context.XML`
- `context.Text`
- `context.Markdown`
- `context.Data`
- `context.Render("my_custom_type",mystructOrData{}, iris.RenderOptions{"gzip":false,"charset":"UTF-8"})`
- `context.MarkdownString`
- `iris.ResponseString(...)`


**Fix**
- https://github.com/kataras/iris/issues/294
- https://github.com/kataras/iris/issues/303


**Small changes**

- `iris.Config.Charset`, before alpha.3 was `iris.Config.Rest.Charset` & `iris.Config.Render.Template.Charset`, but you can override it at runtime by passinth a map `iris.RenderOptions` on the `context.Render` call .
- `iris.Config.IsDevelopment`, before alpha.1 was `iris.Config.Render.Template.IsDevelopment`


**Websockets changes**

No need to import the `github.com/kataras/iris/websocket` to use the `Connection` iteral, the websocket moved inside `kataras/iris` , now all exported variables' names have the prefix of `Websocket`, so the old `websocket.Connection` is now `iris.WebsocketConnection`.


Generally, no other changes on the 'frontend API', for response engines examples and how you can register your own to add more features on existing response engines or replace them, look [here](https://github.com/iris-contrib/response).

**BAD SIDE**: E-Book is still pointing on the v3 release, but will be updated soon.

## 4.0.0-alpha.1 -> 4.0.0-alpha.2

**Sessions were re-written **

- Developers can use more than one 'session database', at the same time, to store the sessions
- Easy to develop a custom session database (only two functions are required (Load & Update)), [learn more](https://github.com/iris-contrib/sessiondb/blob/master/redis/database.go)
- Session databases are located [here](https://github.com/iris-contrib/sessiondb), contributions are welcome
- The only frontend deleted 'thing' is the: **config.Sessions.Provider**
- No need to register a database, the sessions works out-of-the-box
- No frontend/API changes except the `context.Session().Set/Delete/Clear`, they doesn't return errors anymore, btw they (errors) were always nil :)
- Examples (master branch) were updated.

```sh
$ go get github.com/iris-contrib/sessiondb/$DATABASE
```

```go
db := $DATABASE.New(configurationHere{})
iris.UseSessionDB(db)
```


> Note: Book is not updated yet, examples are up-to-date as always.


## 3.0.0 -> 4.0.0-alpha.1

[logger](https://github.com/iris-contrib/logger), [rest](https://github.com/iris-contrib/rest) and all [template engines](https://github.com/iris-contrib/template) **moved** to the [iris-contrib](https://github.com/iris-contrib).

- `config.Logger` -> `iris.Logger.Config`
- `config.Render/config.Render.Rest/config.Render.Template` -> **Removed**
- `config.Render.Rest` -> `rest.Config`
- `config.Render.Template` -> `$TEMPLATE_ENGINE.Config` except Directory,Extensions, Assets, AssetNames,
- `config.Render.Template.Directory` -> `iris.UseTemplate($TEMPLAET_ENGINE.New()).Directory("./templates", ".html")`
- `config.Render.Template.Assets` -> `iris.UseTemplate($TEMPLAET_ENGINE.New()).Directory("./templates",".html").Binary(assetFn func(name string) ([]byte, error), namesFn func() []string)`

- `context.ExecuteTemplate` -> **Removed**, you can use the `context.Response.BodyWriter()` to get its writer and execute html/template engine manually, but this is useless because we support the best support for template engines among all other (golang) web frameworks
- **Added** `config.Server.ReadBufferSize & config.Server.WriteBufferSize` which can be passed as configuration fields inside `iris.ListenTo(config.Server{...})`, which does the same job as `iris.Listen`
- **Added** `iris.UseTemplate($TEMPLAET_ENGINE.New()).Directory("./templates", ".html")` to register a template engine, now iris supports multi template engines, each template engine has its own file extension, no big changes on context.Render except the last parameter:
- `context.Render(filename string, binding interface{}, layout string{})` -> `context.Render(filename string, binding interface{}, options ...map[string]interface{})  | context.Render("myfile.html", myPage{}, iris.Map{"gzip":true,"layout":"layouts/MyLayout.html"}) |`

E-book and examples are not yet updated, no big changes.


## 3.0.0-rc.4 -> 3.0.0-pre.release

- `context.PostFormValue` -> `context.FormValueString`, old func stays until the next revision
- `context.PostFormMulti` -> `context.FormValues` , old func stays until the next revision

- Added `context.VisitAllCookies(func(key,value string))` to visit all your cookies (because `context.Request.Header.VisitAllCookie` has a bug(I can't fix/pr it because the author is away atm))
- Added `context.GetFlashes` to get all available flash messages for a particular request
- Fix flash message removed after the first `GetFlash` call in the same request

**NEW FEATURE**: Built'n support for multi listening servers per iris station, secondary and virtual servers with one-line using the `iris.AddServer` & `iris.Go` to start all servers.

- `iris.SecondaryListen` -> `iris.AddServer`, old func stays until the next revision
- Added `iris.Servers` with this field you can manage your servers very easy
- Added `iris.AddServer/iris.ListenTo/iris.Go`, but funcs like `Listen/ListenTLS/ListenUNIX` will stay forever
- Added `config.Server.Virtual(bool), config.Server.RedirectTo(string) and config.Server.MaxRequestBodySize(int64)`
- Added `iris.Available (channel bool)`
- `iris.HTTPServer` -> `iris.Servers.Main()` to get the main server, which is always the last registered server (if more than one used), old field removed
- `iris.Config.MaxRequestBodySize` -> `config.Server.MaxRequestBodySize`, old field removed

**NEW FEATURE**: Build'n support for your API's end-to-end tests

- Added `tester := iris.Tester(*testing.T)` , look inside: [http_test.go](https://github.com/kataras/iris/blob/master/http_test.go) & [./context_test.go](https://github.com/kataras/iris/blob/master/context_test.go) for `Tester` usage, you can also look inside the [httpexpect's repo](https://github.com/gavv/httpexpect/blob/master/example/iris_test.go) for extended examples with Iris.



## 3.0.0-rc.3 -> 3.0.0-rc.4

**NEW FEATURE**: **Handlebars** template engine support with all Iris' view engine's functions/helpers support, as requested [here](https://github.com/kataras/iris/issues/239):
- `iris.Config.Render.Template.Layout = "layouts/layout.html"`
- `config.NoLayout`
- **dynamic** optional layout on `context.Render`
- **Party specific** layout
- `iris.Config.Render.Template.Handlebars.Helpers["myhelper"] = func()...`
- `{{ yield }} `
- `{{ render }}`
- `{{ url "myroute" myparams}}`
- `{{ urlpath "myroute" myparams}}`

For a complete example please, click [here](https://github.com/iris-contrib/examples/tree/master/templates_handlebars).

**NEW:** Iris **can listen to more than one server per station** now, as requested [here](https://github.com/kataras/iris/issues/235).
For example you can have https with SSL/TLS and one more server http which navigates to the secure location.
Take a look [here](https://github.com/kataras/iris/issues/235#issuecomment-229399829) for an example of this.


**FIXES**
- Fix  `sessions destroy`
- Fix  `sessions persistence on subdomains` (as RFC2109 commands but you can disable it with `iris.Config.Sessions.DisableSubdomainPersistence = true`)


**IMPROVEMENTS**
- Improvements on `iris run` && `iris create`, note that the underline code for hot-reloading moved to [rizla](https://github.com/kataras/rizla).



## 3.0.0-rc.2 -> 3.0.0-rc.3

**Breaking changes**
- Move middleware & their configs to the  [iris-contrib/middleware](https://github.com/iris-contrib/middleware) repository
- Move all plugins & their configs to the [iris-contrib/plugin](https://github.com/iris-contrib/plugin) repository
- Move the graceful package to the [iris-contrib/graceful](https://github.com/iris-contrib/graceful) repository
- Move the mail package & its configs to the [iris-contrib/mail](https://github.com/iris-contrib/mail) repository

Note 1: iris.Config.Mail doesn't not logger exists, use ` mail.Config` from the `iris-contrib/mail`, and ` service:= mail.New(configs); service.Send(....)`.

Note 2: basicauth middleware's context key changed from `context.GetString("auth")` to ` context.GetString("user")`.

Underline changes, libraries used by iris' base code:
- Move the errors package to the [iris-contrib/errors](https://github.com/iris-contrib/errors) repository
- Move the tests package to the [iris-contrib/tests](https://github.com/iris-contrib/tests) repository (Yes, you should make PRs now with no fear about breaking the Iris).

**NEW**:
- OAuth, OAuth2 support via plugin (facebook,gplus,twitter and 25 more), gitbook section [here](https://kataras.gitbooks.io/iris/content/plugin-oauth.html), plugin [example](https://github.com/iris-contrib/examples/blob/master/plugin_oauth_oauth2/main.go), low-level package example [here](https://github.com/iris-contrib/examples/tree/master/oauth_oauth2) (no performance differences, it's just a working version of [goth](https://github.com/markbates/goth) which is converted to work with Iris)

- JSON Web Tokens support via [this middleware](https://github.com/iris-contrib/middleware/tree/master/jwt), book section [here](https://kataras.gitbooks.io/iris/content/jwt.html), as requested [here](https://github.com/kataras/iris/issues/187).

**Fixes**:
- [Iris run fails when not running from ./](https://github.com/kataras/iris/issues/215)
- [Fix or disable colors in iris run](https://github.com/kataras/iris/issues/217).


Improvements to the `iris run` **command**, as requested [here](https://github.com/kataras/iris/issues/192).

[Book](https://kataras.gitbooks.io/iris/content/) and [examples](https://github.com/iris-contrib/examples) are **updated** also.

## 3.0.0-rc.1 -> 3.0.0-rc.2

New:
- ` iris.MustUse/MustUseFunc`  - registers middleware for all route parties, all subdomains and all routes.
- iris control plugin re-written, added real time browser request logger
- `websocket.OnError` - Add OnError to be able to catch internal errors from the connection
- [command line tool](https://github.com/kataras/iris/tree/master/iris) - `iris run main.go` runs, watch and reload on source code changes. As requested [here](https://github.com/kataras/iris/issues/192)

Fixes: https://github.com/kataras/iris/issues/184 , https://github.com/kataras/iris/issues/175 .

## 3.0.0-beta.3, 3.0.0-beta.4 -> 3.0.0-rc.1

This version took me many days because the whole framework's underline code is rewritten after many many many 'yoga'. Iris is not so small anymore, so I (tried) to organized it a little better. Note that, today, you can just go to [iris.go](https://github.com/kataras/iris/tree/master/iris.go) and [context.go](https://github.com/kataras/iris/tree/master/context/context.go) and look what functions you can use. You had some 'bugs' to subdomains, mail service, basic authentication and logger, these are fixed also, see below...

All [examples](https://github.com/iris-contrib/examples) are updated, and I tested them one by one.


Many underline changes but the public API didn't changed much, of course this is relative to the way you use this framework, because that:

- Configuration changes: **0**

- iris.Iris pointer -> **iris.Framework** pointer

- iris.DefaultIris -> **iris.Default**
- iris.Config() -> **iris.Config** is field now
- iris.Websocket() -> **iris.Websocket** is field now
- iris.Logger() -> **iris.Logger** is field now
- iris.Plugins() -> **iris.Plugins** is field now
- iris.Server() -> **iris.HTTPServer** is field now
- iris.Rest() -> **REMOVED**

- iris.Mail() -> **REMOVED**
- iris.Mail().Send() -> **iris.SendMail()**
- iris.Templates() -> **REMOVED**
- iris.Templates().RenderString() -> **iris.TemplateString()**

- iris.StaticHandlerFunc -> **iris.StaticHandler**
- iris.URIOf() -> **iris.URL()**
- iris.PathOf() -> **iris.Path()**

- context.RenderString() returned string,error -> **context.TemplateString() returns only string, which is empty on any parse error**
- context.WriteHTML() -> **context.HTML()**
- context.HTML() -> **context.RenderWithStatus()**

Entirely new

-  -> **iris.ListenUNIX(addr string, socket os.Mode)**
-  -> **context.MustRender, same as Render but send response 500 and logs the error on parse error**
-  -> **context.Log(format string, a...interface{})**
-  -> **context.PostFormMulti(name string) []string**
-  -> **iris.Lookups() []Route**
-  -> **iris.Lookup(routeName string) Route**
-  -> **iris.Plugins.On(eventName string, ...func())** and fire all by **iris.Plugins.Call(eventName)**

- iris.Wildcard() **REMOVED**, subdomains and dynamic(wildcard) subdomains can only be registered with **iris.Party("mysubdomain.") && iris.Party("*.")**


Semantic change for static subdomains

**1**

**BEFORE** :
```go
apiSubdomain := iris.Party("api.mydomain.com")
{
//....
}
iris.Listen("mydomain.com:80")
```


**NOW** just subdomain part, no need to duplicate ourselves:
```go
apiSubdomain := iris.Party("api.")
{
//....
}
iris.Listen("mydomain.com:80")
```
**2**

Before you couldn't set dynamic subdomains and normal subdomains at the same iris station, now you can.
**NOW, this is possible**

```go
/* admin.mydomain.com,  and for other subdomains the Party(*.) */

admin := iris.Party("admin.")
{
	// admin.mydomain.com
	admin.Get("/", func(c *iris.Context) {
		c.Write("INDEX FROM admin.mydomain.com")
	})
	// admin.mydomain.com/hey
	admin.Get("/hey", func(c *iris.Context) {
		c.Write("HEY FROM admin.mydomain.com/hey")
	})
	// admin.mydomain.com/hey2
	admin.Get("/hey2", func(c *iris.Context) {
		c.Write("HEY SECOND FROM admin.mydomain.com/hey")
	})
}

// other.mydomain.com, otadsadsadsa.mydomain.com  and so on....
dynamicSubdomains := iris.Party("*.")
{
	dynamicSubdomains.Get("/", dynamicSubdomainHandler)

	dynamicSubdomains.Get("/something", dynamicSubdomainHandler)

	dynamicSubdomains.Get("/something/:param1", dynamicSubdomainHandlerWithParam)
}
```

Minor change for listen


**BEFORE you could just listen to a port**
```go
iris.Listen("8080")
```
**NOW you have set a HOSTNAME:PORT**
```go
iris.Listen(":8080")
```

Relative issues/features:  https://github.com/kataras/iris/issues/166 , https://github.com/kataras/iris/issues/176, https://github.com/kataras/iris/issues/183,  https://github.com/kataras/iris/issues/184


**Plugins**

PreHandle and PostHandle are removed, no need to use them anymore you can take routes by **iris.Lookups()**, but add support for custom event listeners by **iris.Plugins.On("event",func(){})** and fire all callbacks by **iris.Plugins.Call("event")** .

**FOR TESTERS**

**BEFORE** :
```go
api := iris.New()
//...

api.PreListen(config.Server{ListeningAddr: ""})

e := httpexpect.WithConfig(httpexpect.Config{
	Reporter: httpexpect.NewAssertReporter(t),
	Client:   fasthttpexpect.NewBinder(api.ServeRequest),
})

```

**NOW**:

```go
api := iris.New()
//...

e := httpexpect.WithConfig(httpexpect.Config{
	Reporter: httpexpect.NewAssertReporter(t),
	Client:   fasthttpexpect.NewBinder(api.NoListen().Handler),
})


```

## 3.0.0-beta.2 -> 3.0.0-beta.3

- Complete the Jade Template Engine support, {{ render }} and {{ url }} done also.

- Fix Mail().Send

- Iriscontrol plugin: Replace login using session to basic authentication

And other not-too-important fixes

## 3.0.0-beta -> 3.0.0-beta.2

- NEW: Wildcard(dynamic) subdomains, read [here](https://kataras.gitbooks.io/iris/content/subdomains.html)

- NEW: Implement feature request [#165](https://github.com/kataras/iris/issues/165). Routes can now be selected by `a custom name`, and this allows us to use the {{ url "custom-name" "any" "named" "parameters"}}``: For HTML & Amber engines, example [here](https://github.com/iris-contrib/examples/tree/master/templates_9). For PongoEngine, example [here](https://github.com/iris-contrib/examples/tree/master/templates_10_pongo)

- Remove the [x/net/context](https://godoc.org/golang.org/x/net/context), it has been useless after v2.


## 3.0.0-alpha.beta -> 3.0.0-beta


- New iris.API for easy API declaration, read more [here](https://kataras.gitbooks.io/iris/content/using-handlerapi.html), example [there](https://github.com/iris-contrib/examples/tree/master/api_handler_2).


- Add [example](https://github.com/iris-contrib/examples/tree/master/middleware_basicauth_2) and fix the Basic Authentication middleware

## 3.0.0-alpha.6 -> 3.0.0-alpha.beta

- [Implement feature request to add Globals on the pongo2](https://github.com/kataras/iris/issues/145)

- [Implement feature request for static Favicon ](https://github.com/kataras/iris/issues/141)

- Implement a unique easy only-websocket support:



```go
OnConnection(func(c websocket.Connection){})
```

websocket.Connection
```go

// Receive from the client
On("anyCustomEvent", func(message string) {})
On("anyCustomEvent", func(message int){})
On("anyCustomEvent", func(message bool){})
On("anyCustomEvent", func(message anyCustomType){})
On("anyCustomEvent", func(){})

// Receive a native websocket message from the client
// compatible without need of import the iris-ws.js to the .html
OnMessage(func(message []byte){})

// Send to the client
Emit("anyCustomEvent", string)
Emit("anyCustomEvent", int)
Emit("anyCustomEvent", bool)
Emit("anyCustomEvent", anyCustomType)

// Send via native websocket way, compatible without need of import the iris-ws.js to the .html
EmitMessage([]byte("anyMessage"))

// Send to specific client(s)
To("otherConnectionId").Emit/EmitMessage...
To("anyCustomRoom").Emit/EmitMessage...

// Send to all opened connections/clients
To(websocket.All).Emit/EmitMessage...

// Send to all opened connections/clients EXCEPT this client(c)
To(websocket.NotMe).Emit/EmitMessage...

// Rooms, group of connections/clients
Join("anyCustomRoom")
Leave("anyCustomRoom")


// Fired when the connection is closed
OnDisconnect(func(){})

```

- [Example](https://github.com/iris-contrib/examples/tree/master/websocket)
- [E-book section](https://kataras.gitbooks.io/iris/content/package-websocket.html)


We have some base-config's changed, these configs which are defaulted to true renamed to 'Disable+$oldName'
```go

		// DisablePathCorrection corrects and redirects the requested path to the registed path
		// for example, if /home/ path is requested but no handler for this Route found,
		// then the Router checks if /home handler exists, if yes,
		// (permant)redirects the client to the correct path /home
		//
		// Default is false
		DisablePathCorrection bool

		// DisablePathEscape when is false then its escapes the path, the named parameters (if any).
		// Change to true it if you want something like this https://github.com/kataras/iris/issues/135 to work
		//
		// When do you need to Disable(true) it:
		// accepts parameters with slash '/'
		// Request: http://localhost:8080/details/Project%2FDelta
		// ctx.Param("project") returns the raw named parameter: Project%2FDelta
		// which you can escape it manually with net/url:
		// projectName, _ := url.QueryUnescape(c.Param("project").
		// Look here: https://github.com/kataras/iris/issues/135 for more
		//
		// Default is false
		DisablePathEscape bool

		// DisableLog turn it to true if you want to disable logger,
		// Iris prints/logs ONLY errors, so be careful when you enable it
		DisableLog bool

		// DisableBanner outputs the iris banner at startup
		//
		// Default is false
		DisableBanner bool

```


## 3.0.0-alpha.5 -> 3.0.0-alpha.6

Changes:
	- config/iris.Config().Render.Template.HTMLTemplate.Funcs typeof `[]template.FuncMap` -> `template.FuncMap`


Added:
	- iris.AmberEngine [Amber](https://github.com/eknkc/amber). [View an example](https://github.com/iris-contrib/examples/tree/master/templates_7_html_amber)
	- iris.JadeEngine [Jade](https://github.com/Joker/jade). [View an example](https://github.com/iris-contrib/examples/tree/master/templates_6_html_jade)

Book section [Render/Templates updated](https://kataras.gitbooks.io/iris/content/render_templates.html)



## 3.0.0-alpha.4 -> 3.0.0-alpha.5

- [NoLayout support for particular templates](https://github.com/kataras/iris/issues/130#issuecomment-219754335)
- [Raw Markdown Template Engine](https://kataras.gitbooks.io/iris/content/render_templates.html)
- [Markdown to HTML](https://kataras.gitbooks.io/iris/content/render_rest.html) > `context.Markdown(statusCode int, markdown string)` , `context.MarkdownString(markdown string) (htmlReturn string)`
- [Simplify the plugin registration](https://github.com/kataras/iris/issues/126#issuecomment-219622481)

## 3.0.0-alpha.3 -> 3.0.0-alpha.4

Community suggestions implemented:

- [Request: Rendering html template to string](https://github.com/kataras/iris/issues/130)
	> New RenderString(name string, binding interface{}, layout ...string) added to the Context & the Iris' station (iris.Templates().RenderString)
- [Minify Templates](https://github.com/kataras/iris/issues/129)
	> New config field for minify, defaulted to true: iris.Config().Render.Template.Minify  = true
	> 3.0.0-alpha5+ this has been removed because the minify package has bugs, one of them is this: https://github.com/tdewolff/minify/issues/35.



Bugfixes and enhancements:

- [Static not allowing configuration of `IndexNames`](https://github.com/kataras/iris/issues/128)
- [Processing access error](https://github.com/kataras/iris/issues/125)
- [Invalid header](https://github.com/kataras/iris/issues/123)

## 3.0.0-alpha.2 -> 3.0.0-alpha.3

The only change here is a panic-fix on form bindings. Now **no need to make([]string,0)** before form binding, new example:

```go
 //./main.go

package main

import (
	"fmt"

	"github.com/kataras/iris"
)

type Visitor struct {
	Username string
	Mail     string
	Data     []string `form:"mydata"`
}

func main() {

	iris.Get("/", func(ctx *iris.Context) {
		ctx.Render("form.html", nil)
	})

	iris.Post("/form_action", func(ctx *iris.Context) {
		visitor := Visitor{}
		err := ctx.ReadForm(&visitor)
		if err != nil {
			fmt.Println("Error when reading form: " + err.Error())
		}
		fmt.Printf("\n Visitor: %v", visitor)
	})

	fmt.Println("Server is running at :8080")
	iris.Listen(":8080")
}

```

```html

<!-- ./templates/form.html -->
<!DOCTYPE html>
<head>
<meta charset="utf-8">
</head>
<body>
<form action="/form_action" method="post">
<input type="text" name="Username" />
<br/>
<input type="text" name="Mail" /><br/>
<select multiple="multiple" name="mydata">
<option value='one'>One</option>
<option value='two'>Two</option>
<option value='three'>Three</option>
<option value='four'>Four</option>
</select>
<hr/>
<input type="submit" value="Send data" />

</form>
</body>
</html>

```



## 3.0.0-alpha.1 -> 3.0.0-alpha.2

*The e-book was updated, take a closer look [here](https://www.gitbook.com/book/kataras/iris/details)*


**Breaking changes**

**First**. Configuration owns a package now `github.com/kataras/iris/config` . I took this decision after a lot of thought and I ensure you that this is the best
architecture to easy:

- change the configs without need to re-write all of their fields.
	```go
	irisConfig := config.Iris { Profile: true, PathCorrection: false }
	api := iris.New(irisConfig)
	```

- easy to remember: `iris` type takes config.Iris, sessions takes config.Sessions`, `iris.Config().Render` is `config.Render`, `iris.Config().Render.Template` is `config.Template`, `Logger` takes `config.Logger` and so on...

- easy to find what features are exists and what you can change: just navigate to the config folder and open the type you want to learn about, for example `/iris.go` Iris' type configuration is on `/config/iris.go`

- default setted fields which you can use. They are already setted by iris, so don't worry too much, but if you ever need them you can find their default configs by this pattern: for example `config.Template` has `config.DefaultTemplate()`, `config.Rest` has `config.DefaultRest()`, `config.Typescript()` has `config.DefaultTypescript()`, note that only `config.Iris` has `config.Default()`. I wrote that all structs even the plugins have their default configs now, to make it easier for you, so you can do this without set a config by yourself: `iris.Config().Render.Template.Engine = config.PongoEngine` or `iris.Config().Render.Template.Pongo.Extensions = []string{".xhtml", ".html"}`.



**Second**. Template & rest package moved to the `render`, so

		*  a new config field named `render` of type `config.Render` which nests the `config.Template` & `config.Rest`
		-  `iris.Config().Templates` -> `iris.Config().Render.Template` of type `config.Template`
		- `iris.Config().Rest` -> `iris.Config().Render.Rest` of type `config.Rest`

**Third, sessions**.



Configuration instead of parameters. Before `sessions.New("memory","sessionid",time.Duration(42) * time.Minute)` -> Now:  `sessions.New(config.DefaultSessions())` of type `config.Sessions`

- Before this change the cookie's life was the same as the manager's Gc duration. Now added an Expires option for the cookie's life time which defaults to infinitive, as you (correctly) suggests me in the chat community.-

- Default Cookie's expiration date: from 42 minutes -> to  `infinitive/forever`
- Manager's Gc duration: from 42 minutes -> to '2 hours'
- Redis store's MaxAgeSeconds: from 42 minutes -> to '1 year`


**Four**. Typescript, Editor & IrisControl plugins now accept a config.Typescript/ config.Editor/ config.IrisControl as parameter

Bugfixes

- [can't open /xxx/ path when PathCorrection = false ](https://github.com/kataras/iris/issues/120)
- [Invalid content on links on debug page when custom ProfilePath is set](https://github.com/kataras/iris/issues/118)
- [Example with custom config not working ](https://github.com/kataras/iris/issues/115)
- [Debug Profiler writing escaped HTML?](https://github.com/kataras/iris/issues/107)
- [CORS middleware doesn't work](https://github.com/kataras/iris/issues/108)



## 2.3.2 -> 3.0.0-alpha.1

**Changed**
- `&render.Config` -> `&iris.RestConfig` . All related to the html/template are removed from there.
- `ctx.Render("index",...)` -> `ctx.Render("index.html",...)` or any extension you have defined in iris.Config().Templates.Extensions
- `iris.Config().Render.Layout = "layouts/layout"` -> `iris.Config().Templates.Layout = "layouts/layout.html"`
- `License BSD-3 Clause Open source` -> `MIT License`
**Added**

- Switch template engines via `IrisConfig`. Currently, HTMLTemplate is 'html/template'. Pongo is 'flosch/pongo2`. Refer to the Book, which is updated too, [read here](https://kataras.gitbooks.io/iris/content/render.html).
