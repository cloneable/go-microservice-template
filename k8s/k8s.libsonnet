{
  Labels(name):: {
    "app.kubernetes.io/name": name,
    "app.kubernetes.io/instance": name,
  },

  ServiceAccount(name):: {
    apiVersion: "v1",
    kind: "ServiceAccount",
    metadata: {
      name: name,
      labels: $.Labels(name),
    },
  },

  Service(name):: {
    apiVersion: "v1",
    kind: "Service",
    metadata: {
      name: name,
      labels: $.Labels(name),
    },
  },

  Port(port, name, protocol="TCP"):: {
    port: port,
    targetPort: port,
    name: name,
    protocol: protocol,
  },

  Deployment(name):: {
    apiVersion: "apps/v1",
    kind: "Deployment",
    metadata: {
      name: name,
      labels: $.Labels(name),
    },
  },
}
