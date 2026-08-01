# KUBERNETES

## WHY DO WE NEED K8S WHEN WE HAVE DOCKER?
- Docker is excellent for packagig and running containerized applications. It works well when we are running 10 containers. However, at an organizational level, companies may run hundreds and thousands of containers across multiple servers. Managing them would be hard and that is reason why we need an orchaestration tool like k8s.
- k8s can also manage other things which docker cannotl;
    1. Automatic scaling 
    2. Self healing 
    3. Load balancing 
    4. Service discovery -  so applications can communicate 
    5. rolling updates and rollbacks
    6. scheduling 
    7. high availability
    8. configuration and secret management 
