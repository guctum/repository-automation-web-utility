# Repository Automation Web Utility

**Purpose**

Sample web service that performs basic web request functions to be used for Kubernetes training - specifically creating an API service to use in conjunction with Dynatrace monitoring on Kubernetes to see OneAgent pick up this traffic while learning to deploy something on Kubernetes

## Web Requests
**GitHub**
- Create an empty GitHub repository
- Request Params
    - name - repo name
    - token - provide a Personal Access Token generated from GitHub. This token needs minimum of repo permissions, I gave it repo, gist, and delete_repo scopes to cover potential future use cases
        - Tokens shouldn't be passed as a param for API calls, I have it setup this way solely for ease of testing and training
        - To use this in an enterprise environment, this would need to be modified to take an environment variable from Kubernetes - this would be the GitHub token
