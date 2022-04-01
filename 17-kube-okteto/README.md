# Kubernetes

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