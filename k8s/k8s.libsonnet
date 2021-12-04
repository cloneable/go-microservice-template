{
  Namespace(name):: {
    apiVersion: "v1",
    kind: "Namespace",
    metadata: {
      name: name,
      labels: {
        "kubernetes.io/metadata.name": name,
      },
    },
  },

  Labels(name):: {
    "app.kubernetes.io/name": name,
    "app.kubernetes.io/instance": name,
  },

  ServiceAccount(name, namespace):: {
    apiVersion: "v1",
    kind: "ServiceAccount",
    metadata: {
      name: name,
      namespace: namespace,
      labels: $.Labels(name),
    },
  },

  Service(name, namespace):: {
    apiVersion: "v1",
    kind: "Service",
    metadata: {
      name: name,
      namespace: namespace,
      labels: $.Labels(name),
    },
  },

  ServicePort(port, name, protocol="TCP"):: {
    port: port,
    targetPort: port,
    name: name,
    protocol: protocol,
  },

  Deployment(name, namespace):: {
    apiVersion: "apps/v1",
    kind: "Deployment",
    metadata: {
      name: name,
      namespace: namespace,
      labels: $.Labels(name),
    },
  },
}
