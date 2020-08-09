# What I've done

# Step 1 - Setup

Make sure you have the following software installed:
 - homebrew
 - git
 - docker
 - azure-cli
 - kubectl
 - curl
 - bonus: watch, zsh with k8s auto-complete

Install az cli: `brew install azure-cli`

Login so that we can access the cluster with `kubectl`:

    az login
    az aks get-credentials --resource-group 1step-something --name mr-fields

From the azure "overview" tab, there is a "connect" button that will show the commands above to connect to the kubernetes cluster.

Now we can list stuff:

    kubectl get deployments
    kubectl get pods


# Step 2 - Our App

Build the docker images and tag them with versions:

    docker build -t 1stepsomething/echo-go:v1 ./go-echo/
    docker build -t 1stepsomething/echo-node:v1 ./node-echo/

    docker run -d --name echo-go -p 8081:8000 1stepsomething/echo-go:v1
    docker run -d --name echo-node -p 8082:8001 1stepsomething/echo-node:v1

    docker stop echo-go echo-node

    curl localhost:8081/version
    curl localhost:8082/version

# Step 3 - Docker Images

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
    --docker-username=4f3a1225-b670-421c-9659-61881d21c025 \
    --docker-password=L1UCtOGXg~OjcUpleZnPYtkVWZnEw3pho5
    # note to self: credentials are stored on my password manager

Finally it works:

    kubectl logs echo-go

Now, repeat to create the for the node Pod.


## Bonus

kubectl `create` vs `apply`: 

 * https://stackoverflow.com/questions/47369351/kubectl-apply-vs-kubectl-create
 * both work.
`apply` is incremental and will try to update the resource.
`create` will error if it exists instead
 * both are valid approaches, but we will use `apply` from now on

## Kubernetes Dashboard

# WIP: this is not working yet, i still get access denied

Accessing the web interface of the kubernetes instance:

    az aks browse --resource-group 1step-something  --name mr-fields
    cat ~/.kube/config
    # copy paste the token at the end and use that for login on the dashboard

# Step 4 Deployments

A Deployment provides declarative updates for Pods ReplicaSets.

You describe a desired state in a Deployment, and the Deployment Controller changes the actual state to the desired state at a controlled rate.
You can define Deployments to create new ReplicaSets, or to remove existing Deployments and adopt all their resources with new Deployments.

Create our deployment:

    k apply -f kubefiles/demo-deployment-v1.yml

Play with it:

 * `kubectl describe deployment demo-deployment`

    Notice the: `StrategyType:           RollingUpdate`

    Can also be `Recreate`.
 * change the number of replicas
   * whats did you expect?
   * what did happen?
 * views logs of one container

To clean it up:

    k delete deployment demo-deployment

# Step 5 Network

A service is an abstract way to expose an application running on a set of Pods as a network.

    kubectl apply -f kubefiles/demo-service.yml
    kubectl get services
    curl http://20.50.103.25:80

# Step 6 Update / Rollback

First lets create a new version of our app by repeating the steps above to publish a v2 docker image.
Then, we can edit the deployment yaml and `kubectl apply` the file again.
Thats it!

Things to monitor:

    watch kubectl get pods
    # see each Pod is replaced without downtime

# Next Steps

A lot that we have not covered:

 - DBs
 - environments
 - configuration files
 - monitoring
 - metrics
 - auto scaling
 - etc

 #####################
WIP:

  -- como configurar que docker imgs vao em que replica set
