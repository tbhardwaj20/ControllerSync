# ControllerSync
<img width="1076" height="724" alt="image" src="https://github.com/user-attachments/assets/9b9953c2-9fae-4a0a-b6bd-f470f580b1a0" />


ControllerSync is a Kubernetes custom controller that implements sharding to efficiently manage custom resources across multiple controller replicas. It uses a shard-based consistent hashing approach to distribute the workload evenly and ensure high availability.

## Features

- Sharding of custom resources among multiple controller replicas  
- Automatic resource redistribution when replicas scale up/down or fail  
- Lease-based leader election for each shard  
- Scalable and resilient controller design  
- Works with custom `Notifier` resources  

## Architecture
![Image (1)](https://github.com/user-attachments/assets/c9d1ce1b-b40d-40e9-aad3-087bcadda498)

- **Custom Resource Definitions (CRDs):** Defines the `Notifier` resource managed by the controller.  
- **Controller Rings:** Represents shards and manages resource assignment via consistent hashing.  
- **Leases:** Used for leader election to coordinate shard ownership among replicas.  
- **Controllers:** Multiple replicas coordinate through shard assignment, ensuring balanced processing.

## Getting Started

### Prerequisites

- Kubernetes cluster (tested with Kind)  
- kubectl CLI configured  
- Go environment (for building from source)

### Installation

1. Create a Kind cluster:  
   ```
   kind create cluster
   ```
2. Deploy reusable components:
    ```
    kubectl apply -k https://github.com/timebertt/kubernetes-controller-sharding/config/default?ref=main
    ```
3. Create Custom Resource Definitions:
    ```
    kubectl apply -k https://github.com/timebertt/kubernetes-controller-sharding/config/crds?ref=main
    ```
4. Deploy the controller:
   ```
   kubectl apply -k https://github.com/tbhardwaj20/ControllerSync/config/manager
   ```
5. Create Controller Rings and Notifier resources as needed.

## Usage
- Scale controller replicas to see sharding and resource redistribution in action:

```
kubectl scale deployment controllersync-controller-manager -n controllersync-system --replicas=3
```
- View resource-to-controller assignment:

```
kubectl get notifiers -A --show-labels
```
- Monitor leases for shard ownership:

```
kubectl get lease -n controllersync-system
```

## Development
Clone the repo and build the controller:
```
git clone https://github.com/tbhardwaj20/ControllerSync.git
cd controllersync
make build
```
Run tests and generate manifests as needed.

## References
- [Kubernetes Controller Sharding](https://github.com/timebertt/kubernetes-controller-sharding)

## License
MIT License © 2025 [Tushar Bhardwaj](https://github.com/TuShArBhArDwA)

This project was inspired by and based on the Kubernetes controller sharding concept to achieve scalable custom controllers.

