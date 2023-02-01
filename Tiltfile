k8s_yaml(kustomize("./gitops/overlays/local"))

k8s_resource('wotapi', port_forwards='8080:8080', labels='api')
k8s_resource('postgresql',
  port_forwards='5432:5432',
  objects=[
    'postgresql-config',
    'postgresql-data:persistentvolumeclaim'
  ],
  labels='db'
)

docker_build('stianfro/wotapi', '.', dockerfile='Dockerfile')