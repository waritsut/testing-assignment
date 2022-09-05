package router

import (
	"fmt"
	"net/http"
	"runtime/debug"
	"sync"

	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"
)

type Context interface {
	BindJson(v interface{}) error
	BindForm(obj interface{}) error
	JSON(statusCode int, v interface{})
	Param(key string) string
}

type MyContext struct {
	ctx *gin.Context
}

func NewMyContext(c *gin.Context) *MyContext {
	return &MyContext{ctx: c}
}

func (c *MyContext) BindJson(v interface{}) error {
	return c.ctx.ShouldBindJSON(v)
}

func (c *MyContext) BindForm(obj interface{}) error {
	return c.ctx.Bind(obj)
}

func (c *MyContext) Param(key string) string {
	return c.ctx.Param(key)
}

func (c *MyContext) JSON(statusCode int, v interface{}) {
	c.ctx.JSON(statusCode, v)
}

func NewGinHandler(handler func(Context)) gin.HandlerFunc {
	return func(c *gin.Context) {
		handler(NewMyContext(c))
	}
}

type MyRouter struct {
	*gin.Engine
}

var _ Mux = (*mux)(nil)

type Mux interface {
	http.Handler

	Group(relativePath string, handlers ...HandlerFunc) RouterGroup
}

type mux struct {
	engine *gin.Engine
}

func (m *mux) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	m.engine.ServeHTTP(w, req)
}

func (m *mux) Group(relativePath string, handlers ...HandlerFunc) RouterGroup {
	return &router{
		group: m.engine.Group(relativePath, wrapHandlers(handlers...)...),
	}
}

type HandlerFunc func(c Context)

type IRoutes interface {
	GET(string, ...HandlerFunc)
	POST(string, ...HandlerFunc)
	PATCH(string, ...HandlerFunc)
	PUT(string, ...HandlerFunc)
}

type RouterGroup interface {
	Group(string, ...HandlerFunc) RouterGroup
	IRoutes
}

func wrapHandlers(handlers ...HandlerFunc) []gin.HandlerFunc {
	funcs := make([]gin.HandlerFunc, len(handlers))
	for i, handler := range handlers {
		handler := handler
		funcs[i] = func(c *gin.Context) {
			ctx := newContext(c)
			defer releaseContext(ctx)

			handler(ctx)
		}
	}

	return funcs
}

var contextPool = &sync.Pool{
	New: func() interface{} {
		return new(MyContext)
	},
}

func newContext(ctx *gin.Context) Context {
	context := contextPool.Get().(*MyContext)
	context.ctx = ctx
	return context
}

func releaseContext(ctx Context) {
	c := ctx.(*MyContext)
	c.ctx = nil
	contextPool.Put(c)
}

type router struct {
	group *gin.RouterGroup
}

func (r *router) Group(relativePath string, handlers ...HandlerFunc) RouterGroup {
	group := r.group.Group(relativePath, wrapHandlers(handlers...)...)
	return &router{group: group}
}

func (r *router) PATCH(relativePath string, handlers ...HandlerFunc) {
	r.group.PATCH(relativePath, wrapHandlers(handlers...)...)
}

func (r *router) POST(relativePath string, handlers ...HandlerFunc) {
	r.group.POST(relativePath, wrapHandlers(handlers...)...)
}

func (r *router) GET(relativePath string, handlers ...HandlerFunc) {
	r.group.GET(relativePath, wrapHandlers(handlers...)...)
}

func (r *router) PUT(relativePath string, handlers ...HandlerFunc) {
	r.group.PUT(relativePath, wrapHandlers(handlers...)...)
}

func NewMyRouter() Mux {
	gin.SetMode(gin.DebugMode)

	mux := &mux{
		engine: gin.New(),
	}

	mux.engine.Use(cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{
			http.MethodHead,
			http.MethodGet,
			http.MethodPost,
			http.MethodPut,
			http.MethodPatch,
			http.MethodDelete,
		},
		AllowedHeaders:     []string{"*"},
		AllowCredentials:   true,
		OptionsPassthrough: true,
	}))

	mux.engine.Use(func(ctx *gin.Context) {
		context := newContext(ctx)
		defer releaseContext(context)

		ctx.Next()

	})

	mux.engine.Use(func(ctx *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				str := fmt.Sprintf("%v %s", err, debug.Stack())
				ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": str})
			}
		}()

	})

	return mux
}
