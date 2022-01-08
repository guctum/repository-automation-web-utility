# DT K8s Training - Custom Web Utils Service

**Purpose**

Sample web service that performs basic web request functions to be used for Kubernetes training - specifically creating an API service to use in conjunction with Dynatrace monitoring on Kubernetes to see OneAgent pick up this traffic while learning to deploy something on Kubernetes

## Web Requests
**GitHub**
- Create an empty GitHub repository
- Request Params
    - name - repo name
    - token - provide a Personal Access Token generated from GitHub. This token needs minimum of repo permissions, I gave it repo, gist, and delete_repo scopes to cover potential future use cases
        - Tokens shouldn't be passed as a param for an API call, eventually this will be converted to an env variable on the kubernetes cluster - just went with this approach for testing as I build this out