# Using [kitex](https://github.com/cloudwego/kitex) in kubernetes

Kitex [kaɪt'eks] is a high-performance and strong-extensibility Golang RPC framework. 
This go module helps you to build multiple services in single kubernetes cluster without 
using extra service discovery facilities.


## how to use 

Client and server in the same kubernetes namespace

```go
import (
    "github.com/adolli/kitex-k8s-native-discovery"
    "github.com/cloudwego/kitex/client"
)

var cli your_rpc_client_code_gen.Client

serviceName := "my.awsome.service"   // (IMPORTANT!) call service by name 
var opts []client.Option
suite := kitex_k8s_native_discovery.NewDiscovery() // initialized this to run client and server in the same namespace
opts = append(opts, client.WithSuite(suite))

cli = your_rpc_client_code_gen.MustNewClient(serviceName, opts...)
```

Client and server in different kubernetes namespaces

```go
import (
    "github.com/adolli/kitex-k8s-native-discovery"
    "github.com/cloudwego/kitex/client"
)

var cli your_rpc_client_code_gen.Client

serviceName := "my.awsome.service"   // (IMPORTANT!) call service by name 
serverNamespace := "server-production"
var opts []client.Option
suite := kitex_k8s_native_discovery.NewDiscoveryWithNamespace(serverNamespace) // <- pass the namespace to initialize  
opts = append(opts, client.WithSuite(suite))

cli = your_rpc_client_code_gen.MustNewClient(serviceName, opts...)
```

## conventions

Your server deployed in kubernetes cluster should do the following steps

- Use a Service object to expose your kitex service
- The Service object should be named as in the form of `a-b-c` corresponding to your service name defined in kitex server
- The Service object should expose a port 80 (check the following example)

If your server is initialized in this way and named `my.awsome.service`

```bash
kitex -module "your_module_name" -service my.awsome.service hello.thrift
```

Then apply a Service object like this

```yaml
apiVersion: v1
kind: Service
metadata:
  name: my-awsome-service # <- (IMPORTANT!) define the Service name, replace all '.' and '_' to '-' 
spec:
  type: ClusterIP
  ports:
    - name: http
      protocol: TCP
      port: 80            # <- (IMPORTANT!) this port number must be 80
      targetPort: 6789    # <- this corresponding to port listened by kitex server
  selector:
    app: awsome-app 
```


Visit [kitex getting started](https://www.cloudwego.io/docs/getting-started/) to get more usages about kitex.
