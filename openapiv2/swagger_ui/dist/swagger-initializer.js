window.onload = function () {
    const servicesUrl = new URL("../services", window.location.href);
    fetch(servicesUrl.toString())
        .then(response => response.json())
        .then(data => {
            const urls = data.services.filter((x) => [
                "grpc.health.v1.Health",
                "kratos.api.Metadata",
                "grpc.reflection.v1alpha.ServerReflection",
            ].indexOf(x) === -1).map((x) => {
                const url = new URL("../service/" + x, window.location.href);
                return {url: url.toString(), name: x}
            });
            // Begin Swagger UI call region
            const ui = SwaggerUIBundle({
                urls: urls,
                dom_id: '#swagger-ui',
                deepLinking: true,
                presets: [
                    SwaggerUIBundle.presets.apis,
                    SwaggerUIStandalonePreset,
                ],
                plugins: [
                    SwaggerUIBundle.plugins.Topbar,
                    SwaggerUIBundle.plugins.DownloadUrl,
                ],
                layout: "StandaloneLayout"
            });
            // End Swagger UI call region

            window.ui = ui;
        });

};