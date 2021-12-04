local k8s = import "k8s.libsonnet";

[
  k8s.ServiceAccount("go-microservice-template"),

  k8s.Service("go-microservice-template") {
    spec: {
      type: "ClusterIP",
      ports: [
        k8s.Port(8080, "rest-api"),
        k8s.Port(9090, "grpc-api"),
        k8s.Port(12345, "service"),
      ],
      selector: k8s.Labels("go-microservice-template"),
    },
  },

  k8s.Deployment("go-microservice-template") {
    spec: {
      replicas: 1,
      selector: {
        matchLabels: k8s.Labels("go-microservice-template"),
      },
      template: {
        metadata: {
          labels:
            k8s.Labels("go-microservice-template"),
        },
        spec: {
          serviceAccountName: "go-microservice-template",
          securityContext: {},
          containers: [
            {
              name: "go-microservice-template",
              securityContext: {},
              image: "cloneable/go-microservice-template:dev",
              imagePullPolicy: "Always",
              ports: [
                {
                  name: "rest",
                  containerPort: 8080,
                  protocol: "TCP",
                },
                {
                  name: "grpc",
                  containerPort: 9090,
                  protocol: "TCP",
                },
                {
                  name: "service",
                  containerPort: 12345,
                  protocol: "TCP",
                },
              ],
              livenessProbe: {
                httpGet: {
                  path: "/healthz",
                  port: "rest",
                },
              },
              readinessProbe: {
                httpGet: {
                  path: "/healthz",
                  port: "rest",
                },
              },
            },
          ],
        },
      },
    },
  },
]
