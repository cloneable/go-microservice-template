local k8s = import "k8s.libsonnet";

local Namespace = "dev";

[
  k8s.Namespace(Namespace),

  k8s.ServiceAccount("go-microservice-template", Namespace),

  k8s.Service("go-microservice-template", Namespace) {
    spec: {
      type: "ClusterIP",
      ports: [
        k8s.ServicePort(8080, "rest-api"),
        k8s.ServicePort(9090, "grpc-api"),
        k8s.ServicePort(12345, "service"),
      ],
      selector: k8s.Labels("go-microservice-template"),
    },
  },

  k8s.Deployment("go-microservice-template", Namespace) {
    spec: {
      replicas: 1,
      selector: {
        matchLabels: k8s.Labels("go-microservice-template"),
      },
      template: {
        metadata: {
          labels: k8s.Labels("go-microservice-template"),
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
                  path: "/healthz/alive",
                  port: "rest",
                },
                periodSeconds: 60,
                successThreshold: 1,
                failureThreshold: 3,
                timeoutSeconds: 2,
                initialDelaySeconds: 5,
              },
              readinessProbe: {
                httpGet: {
                  path: "/healthz/ready",
                  port: "rest",
                },
                periodSeconds: 60,
                successThreshold: 1,
                failureThreshold: 3,
                timeoutSeconds: 2,
                initialDelaySeconds: 5,
              },
            },
          ],
        },
      },
    },
  },
]
