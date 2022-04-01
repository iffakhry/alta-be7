# Kubernetes
Orchestration
![orchestration-1.png]
![orchestration-2.png]

## Preparation

- Install kubectl
- Register di [Okteto](https://cloud.okteto.com/#/login)
- Download [Lens](https://k8slens.dev/) (optional)

## Check

```bash
kubectl config get-contexts

#pastikan sudah muncul cloud okteto
```

## How to Deploy resources

```bash
kubectl -n <namespace> apply -f <file>
```

## See resources

```bash
kubectl -n <namespace> get <resource>
```

## Delete resources

```bash
kubectl -n <namespace> delete <resource>
```

## See list of object
```bash
    kubectl get <kind>
```

# How to run go-app
```bash
    # go to directory go-app
    cd go-app

    # apply
    kubectl -n iffakhry apply -f mydeployment.yaml
    kubectl -n iffakhry apply -f myservice.yaml
    kubectl -n iffakhry apply -f myingress.yaml

    # check ingress
    kubectl get ingress
    # should show the ingress. example iffakhry.cloud.okteto.net
    # and you can try on postman https://iffakhry.cloud.okteto.net/hello

```