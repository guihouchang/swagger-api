package openapiv2

import (
	"net/http"
	"net/url"
	"strings"

	"github.com/go-kratos/grpc-gateway/v2/protoc-gen-openapiv2/generator"
	"github.com/go-kratos/kratos/v2/api/metadata"
	"github.com/go-kratos/kratos/v2/transport/http/binding"
	"github.com/gorilla/mux"
	_ "github.com/guihouchang/swagger-api/openapiv2/swagger_ui/statik" // import statik static files
	"github.com/rakyll/statik/fs"
)

func NewHandler(handlerOpts ...HandlerOption) http.Handler {
	opts := &options{
		// Compatible with default UseJSONNamesForFields is true
		generatorOptions: []generator.Option{generator.UseJSONNamesForFields(true)},
		pathPrefix:       "/q/swagger-ui", // 默认路径前缀
	}

	for _, o := range handlerOpts {
		o(opts)
	}

	service := New(nil, opts.generatorOptions...)
	r := mux.NewRouter()

	// 从pathPrefix中提取API基础路径
	apiBasePath := "/q"
	if opts.pathPrefix != "/q/swagger-ui" {
		// 如果是自定义路径前缀，提取基础路径
		// 例如：/user/q/swagger-ui -> /user/q
		if idx := strings.LastIndex(opts.pathPrefix, "/swagger-ui"); idx > 0 {
			apiBasePath = opts.pathPrefix[:idx]
		}
	}

	r.HandleFunc(apiBasePath+"/services", func(w http.ResponseWriter, r *http.Request) {
		services, err := service.ListServices(r.Context(), &metadata.ListServicesRequest{})
		if err != nil {
			w.WriteHeader(500)
			w.Write([]byte(err.Error()))
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		w.Write([]byte(services.String()))
	}).Methods("GET")

	r.HandleFunc(apiBasePath+"/services/{service}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		serviceName := vars["service"]

		in := metadata.GetServiceDescRequest{
			Name: serviceName,
		}

		if r.URL.RawQuery != "" {
			values, err := url.ParseQuery(r.URL.RawQuery)
			if err != nil {
				w.WriteHeader(400)
				w.Write([]byte(err.Error()))
				return
			}

			if err := binding.BindQuery(values, &in); err != nil {
				w.WriteHeader(400)
				w.Write([]byte(err.Error()))
				return
			}
		}

		content, err := service.GetServiceOpenAPI(r.Context(), &in, false)
		if err != nil {
			w.WriteHeader(500)
			w.Write([]byte(err.Error()))
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		w.Write([]byte(content))
	}).Methods("GET")

	// 静态文件处理器 - 处理CSS、JS等静态资源
	statikFS, err := fs.New()
	if err != nil {
		panic(err)
	}
	staticServer := http.FileServer(statikFS)

	// 为静态资源创建更具体的路由，避免与openapi.json和根路径冲突
	r.PathPrefix(opts.pathPrefix + "/swagger-ui").Handler(http.StripPrefix(opts.pathPrefix, staticServer))
	r.PathPrefix(opts.pathPrefix + "/favicon").Handler(http.StripPrefix(opts.pathPrefix, staticServer))

	// 添加openapi.json路由
	r.HandleFunc(opts.pathPrefix+"/openapi.json", func(w http.ResponseWriter, r *http.Request) {
		// 这里可以返回一个默认的OpenAPI规范或者重定向到具体的服务
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		w.Write([]byte(`{"openapi":"3.0.0","info":{"title":"API Documentation","version":"1.0.0"},"paths":{}}`))
	}).Methods("GET")

	// 添加动态HTML处理
	r.HandleFunc(opts.pathPrefix+"/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		w.WriteHeader(200)
		w.Write([]byte(generateSwaggerUIHTML(opts.pathPrefix, apiBasePath)))
	}).Methods("GET")

	return r
}
