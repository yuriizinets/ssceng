# Core Features

This part of documentation holds core library features.  
This features highly relies on rendering lifecycle and covers all low-level operations.

## Page rendering

The most important part of the library.  
Low-level function, responsible for page rendering directly into `io.Writer`.  
Under the hood, executes full rendering [lifecycle](/concepts/#rendering-lifecycle). Library has more high-level wrappers with context setters and another features, but all of them rely on this function. Accepts 2 parameters - page pointer and `io.Writer`.  

First of all, let's create a page structure

`page.index.go`

```go
package main

import (
    "github.com/yuriizinets/kyoto"
)

type PageIndex struct {}
```

As a requirement, each page must have an html template builder method.  
Please note that providing `kyoto.Funcs()` is not required, but highly recommended as far as it provides some library features.

`page.index.go`

```go
...

func (*PageIndex) Template() *template.Template {
    return template.Must(template.New("page.index.html").Funcs(kyoto.Funcs()).ParseGlob("*.html"))
}
```

After creating page structure, it's time to create template

`page.index.html`

```html
<html>
    <head>
        <title>kyoto page</title>
    </head>
    <body>
        ...
    </body>
</html>
```

Now you can use rendering function

```go
func ExampleHandler(rw http.ResponseWriter, r *http.Request) {
    RenderPage(rw, &PageIndex{})
}
```

::: details Complete example with component
Let's add an example component to make this example more complete. This component will generate random number and will hold that number as content.
Please note that component template definition must to much with actual component structure name.

`component.rand.go`

```go
package main

import (
    "crypto/rand"
    "strconv"
)

type ComponentRand struct {
    Content string
}

func (c *ComponentRand) Init() {
    c.Content = strconv.Itoa(rand.Intn(1000))
}
```

`component.rand.html`

```html
{{ define "ComponentRand" }}
    <div>Random number: {{ .Content }}</div>
{{ end }}
```

After component creation, let's register and include it into page.  
Check [Lifecycle integration](/concepts/#lifecycle-integration) section for detailed documentation.

`page.index.go`

```go
package main

import (
    "github.com/yuriizinets/kyoto"
)

type PageIndex struct {
    Rand kyoto.Component
}

func (p *PageIndex) Init() {
    p.Rand = kyoto.RegC(p &ComponentRand{})
}
```

`page.index.html`

```html
<html>
    <head>
        <title>kyoto page</title>
    </head>
    <body>
        {{ template "ComponentRand" .Rand }}
    </body>
</html>
```

That's it! Now you have component instance, included into lifecycle and rendered on the page.

:::

## Context management

You can use `kyoto.SetContext`, `kyoto.GetContext` and `kyoto.DelContext` for managing your context.  

Context uses page instance as namespace for correct concurrency handling on requests level (page instance is creating for each new request).
Context can be used for passing additional state (f.e.`http.Request`, `gin.Context`) which can be accessed inside of lifecycle methods, like `Init` or `Async`.  
**It's important to cleanup context with `kyoto.DelContext(p, "")` after page processing to avoid memory leaks!**

Example of usage:

```go
func IndexPageHandler(rw http.ResponseWriter, r *http.Request) {
    p := &PageIndex{}
    kyoto.SetContext(p, "internal:r", r)
    kyoto.SetContext(p, "internal:rw", rw)
    kyoto.RenderPage(rw, )
    kyoto.DelContext(p, "")
}
...
func (p *PageIndex) Init() {
    r := kyoto.GetContext(p, "internal:r").(*http.Request)
    rw := kyoto.GetContext(p, "internal:rw").(http.ResponseWriter)
    ...
}
...
```

Most of the component methods have overload option with page argument. In this way you don't need to store page pointer in the component itself. Check full [interfaces specification](/concepts/#interfaces) in the [Concepts](/concepts) section.  
Example of overloaded asynchronous method:  

```go
...

func (*ComponentExample) Async(p kyoto.Page) error {
    r := kyoto.GetContext(p, "internal:r").(*http.Request)
}

...
```

## Component lifecycle

This section extends [Lifecycle integration](/concepts/#lifecycle-integration) documentation with examples.

### Init

`Init` method is triggering on initialization lifecycle step.  

Usage:  

```go
...

func (*ComponentExample) Init() {
    // Do what you want here
}

...
```

This method have overload option with page argument

```go
...

func (*ComponentExample) Init(p kyoto.Page) {
    // Do what you want here
}

...
```

### Async

`Async` method is triggering on asynchronous operations lifecycle step. You need to return an error in case of async operation failure.  
> Very useful in case of time-consuming operations, when you need to fetch data from external API or database.  

Usage:  

```go
...

func (*ComponentExample) Async() error {
    // Do what you want here
    return nil
}

...
```

This method have overload option with page argument

```go
...

func (*ComponentExample) Async(p kyoto.Page) error {
    // Do what you want here
    return nil
}

...
```

### AfterAsync

`AfterAsync` method is triggering after asynchronous operations lifecycle step.  

Usage:

```go
...

func (*ComponentExample) AfterAsync() {
    // Do what you want here
}

...
```

This method have overload option with page argument

```go
...

func (*ComponentExample) AfterAsync(p kyoto.Page) {
    // Do what you want here
}

...
```

## Handler factory

High-level function for creating own page handler with context setting.  
Takes 2 parameters: page pointer and `map[string]interface{}` containing context items (`SetContext` will be called for each item in map)

Usage:

```go
func pagehandler(p kyoto.Page) http.HandlerFunc {
    return func(rw http.ResponseWriter, r *http.Request) {
        kyoto.PageHandlerFactory(p, map[string]interface{}{
            "internal:rw": rw,
            "internal:r":  r,
        })(rw, r)
    }
}

...

func main() {
    ...
    mux.HandleFunc("/", pagehandler(&PageIndex{}))
    ...
}
```