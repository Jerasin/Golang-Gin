# Helper

**How to Debug Dockerfile** 
-  run docker run -it {img_name} /bin/bash

**Generate OpenAPI**

- ref: https://github.com/swaggo/swag/issues/810
```
swag init --parseDependency --parseInternal -g ./main.go -o ./docs
```

  **Generate Dependency Injection**
```
wire gen ~/workspace/go-gin/app/module
```

  **Swagger UI**
```
http://localhost:3000/swagger/index.html.
```