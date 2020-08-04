# What I've done

# Step 1 - Setup

Create a kubernetes cluster on azure.

Install az cli: `brew install azure-cli`

Login so that we can access the cluster with `kubectl`:

    az login
    az aks get-credentials --resource-group 1step-something --name mr-fields

From the azure "overview" tab, there is a "connect" button that will show the commands above to connect to the kubernetes cluster.

Now we can list stuff:

    kubectl get deployments
    kubectl get pods


# Step 2 - Services

Build the docker images and tag them with versions:

    docker build -t 1stepsomething/echo-go:v1 ./go-echo/
    docker build -t 1stepsomething/echo-node:v1 ./node-echo/

    docker run -d --name echo-go -p 8081:8000 1stepsomething/echo-go:v1
    docker run -d --name echo-node -p 8082:8001 1stepsomething/echo-node:v1

    docker stop echo-go echo-node

    curl localhost:8081/version
    curl localhost:8082/version

# Step 3 - Docker Registry

When we create the  azure kubernetes cluster it creates a docker registry for us.

From the web UI we can get the credentials to login:

    docker login 25friday.azurecr.io
    # user and pass from the azure web UI

Then we tag and push our images:

    docker tag 1stepsomething/echo-go:v1 25friday.azurecr.io/echo-go:v1
    docker tag 1stepsomething/echo-node:v1 25friday.azurecr.io/echo-node:v1
    docker push 25friday.azurecr.io/echo-go:v1
    docker push 25friday.azurecr.io/echo-node:v1

# Step 3 - Kubectl

Create namespace from the terminal:

    kubectl create namespace demo
    kubectl get namespaces

    kubectl get namespace/demo -o yaml
    kubectl delete namespace/demo

Create a file instead for the namespace creation:

    kubectl create -f kubefiles/demo-namespace.yml

Switch to our namespace so that we don't need to always pass it as a command line argument.
All commands are now targeted only at this namespace.

    kubectl config set-context --current --namespace=demo

Create a pod and deploy it:

    kubectl apply -f kubefiles/go-pod.yml

What is going on?

    kubectl get pods
    kubectl describe pods echo-go

This will fail because:

 * `https://docs.microsoft.com/en-us/azure/container-registry/container-registry-auth-kubernetes`

So we need to:

    kubectl create secret docker-registry <secret-name> \
    --namespace <namespace> \
    --docker-server=<container-registry-name>.azurecr.io \
    --docker-username=<service-principal-ID> \
    --docker-password=<service-principal-password>

    kubectl create secret docker-registry acr-pull-for-namespace-demo \
    --namespace demo \
    --docker-server=25friday.azurecr.io \
    --docker-username=foo \
    --docker-password=bar
    # note to self: credentials are stored on my password manager

Finally it works:

    kubectl logs echo-go

Now, repeat to create the for the node pod.

TODO: how to get urls/ip to access the pods ??

## Bonus

kubectl `create` vs `apply`: 

 * https://stackoverflow.com/questions/47369351/kubectl-apply-vs-kubectl-create
 * both work.
`apply` is incremental and will try to update the resource.
`create` will error if it exists instead

## Kubernetes Dashboard

Accessing the web interface of the kubernetes instance:

    az aks browse --resource-group 1step-something  --name mr-fields
    cat ~/.kube/config
    # copy paste the token at the end and use that for login on the dashboard
    # WIP: this is not working yet, i still get access denied


# Step Replica Set ?? or deployments ??