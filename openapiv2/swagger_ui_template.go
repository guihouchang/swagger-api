/*
 * @Author: guihouchang guihouchang@163.com
 * @Date: 2025-10-29 15:59:57
 * @LastEditors: guihouchang guihouchang@163.com
 * @LastEditTime: 2025-10-29 16:09:48
 * @FilePath: /swagger-api/openapiv2/swagger_ui_template.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package openapiv2

import (
	"fmt"
	"strings"
)

// generateSwaggerUIHTML 生成动态的Swagger UI HTML内容
func generateSwaggerUIHTML(pathPrefix string, apiBasePath string) string {
	// 确保pathPrefix以/开头但不以/结尾
	if pathPrefix == "" {
		pathPrefix = "/q/swagger-ui"
	}
	if !strings.HasPrefix(pathPrefix, "/") {
		pathPrefix = "/" + pathPrefix
	}
	if strings.HasSuffix(pathPrefix, "/") {
		pathPrefix = strings.TrimSuffix(pathPrefix, "/")
	}

	return fmt.Sprintf(`<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="UTF-8">
    <title>Swagger UI</title>
    <link rel="stylesheet" type="text/css" href="%s/swagger-ui.css" />
    <link rel="icon" type="image/png" href="%s/favicon-32x32.png" sizes="32x32" />
    <link rel="icon" type="image/png" href="%s/favicon-16x16.png" sizes="16x16" />
    <style>
      html
      {
        box-sizing: border-box;
        overflow: -moz-scrollbars-vertical;
        overflow-y: scroll;
      }

      *,
      *:before,
      *:after
      {
        box-sizing: inherit;
      }

      body
      {
        margin:0;
        background: #fafafa;
      }
    </style>
  </head>

  <body>
    <div id="swagger-ui"></div>

    <script src="%s/swagger-ui-bundle.js" charset="UTF-8"> </script>
    <script src="%s/swagger-ui-standalone-preset.js" charset="UTF-8"> </script>
    <script>
    window.onload = function() {
      // Begin Swagger UI call region
      const ui = SwaggerUIBundle({
        url: '%s/openapi.json',
        dom_id: '#swagger-ui',
        deepLinking: true,
        presets: [
          SwaggerUIBundle.presets.apis,
          SwaggerUIStandalonePreset
        ],
        plugins: [
          SwaggerUIBundle.plugins.DownloadUrl
        ],
        layout: "StandaloneLayout"
      });
      // End Swagger UI call region

      window.ui = ui;
    };
  </script>
  </body>
</html>`, pathPrefix, pathPrefix, pathPrefix, pathPrefix, pathPrefix, pathPrefix)
}
