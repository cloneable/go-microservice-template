local Labels = {
  "app.kubernetes.io/name": "go-microservice-template",
  "app.kubernetes.io/instance": "go-microservice-template",
};

local ServiceAccount(name) = {
  apiVersion: "v1",
  kind: "ServiceAccount",
  metadata: {
    name: name,
    labels: Labels,
  },
};

local Service(name) = {
  apiVersion: "v1",
  kind: "Service",
  metadata: {
    name: name,
    labels: Labels,
  },
};

local Port(port, name, protocol="TCP") = {
  port: port,
  targetPort: port,
  name: name,
  protocol: protocol,
};

local Deployment(name) = {
  apiVersion: "apps/v1",
  kind: "Deployment",
  metadata: {
    name: name,
    labels: Labels,
  },
};

[
  ServiceAccount("go-microservice-template"),

  Service("go-microservice-template") {
    spec: {
      type: "ClusterIP",
      ports: [
        Port(8080, "rest-api"),
        Port(9090, "grpc-api"),
        Port(12345, "service"),
      ],
      selector: Labels,
    },
  },

  Deployment("go-microservice-template") {
    spec: {
      replicas: 1,
      selector: {
        matchLabels: Labels,
      },
      template: {
        metadata: {
          labels:
            Labels,
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
