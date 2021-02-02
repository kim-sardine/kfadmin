
# kfadmin : CLI Tool for Kubeflow Admin

- [kfadmin : CLI Tool for Kubeflow Admin](#kfadmin--cli-tool-for-kubeflow-admin)
    - [Prerequitite](#prerequitite)
    - [Confirmed running environment](#confirmed-running-environment)
    - [Commands](#commands)
        - [Users](#users)
            - [Create User](#create-user)
            - [List Users](#list-users)
            - [Update user password](#update-user-password)
            - [Delete User](#delete-user)
        - [Namespace (Kubeflow Profile)](#namespace-kubeflow-profile)
            - [Create namespace (NSY)](#create-namespace-nsy)
            - [Update namespace owner (NSY)](#update-namespace-owner-nsy)
            - [Update namespace resourceQuota (TBD)](#update-namespace-resourcequota-tbd)
            - [Add user to namespace as contributor (NSY)](#add-user-to-namespace-as-contributor-nsy)
            - [Remove contributor from (NSY)](#remove-contributor-from-nsy)
            - [List namespace (NSY)](#list-namespace-nsy)
            - [Delete namespace (NSY)](#delete-namespace-nsy)
        - [Secret](#secret)
            - [Create Generic Secret (NSY)](#create-generic-secret-nsy)
            - [List Secret (NSY)](#list-secret-nsy)
            - [Delete Secret (NSY)](#delete-secret-nsy)
            - [Create Docker Registry Secret (NSY)](#create-docker-registry-secret-nsy)
            - [List Docker Registry Secret (NSY)](#list-docker-registry-secret-nsy)
            - [Set Docker Registry Secret as default (TBD)](#set-docker-registry-secret-as-default-tbd)
            - [Delete Docker Registry Secret (TBD)](#delete-docker-registry-secret-tbd)

## Prerequitite

- `kubectl`
- kubernetes context
    - commands like `kubectl get nodes` should be working.

## Confirmed running environment

- kubectl : 1.19.3
- Kubernetes cluster : 1.19.3
- kfctl : 1.2.0 (kfctl_istio_dex.v1.2.0.yaml)

## Commands

> NSY : Not Supported Yet

> TBD : To Be Decided

### Users

#### Create User

`kfadmin create user -e USER_EMAIL -p PASSWORD`

#### List Users

`kfadmin list user`

#### Update user password

`kfadmin update user password -e USER_EMAIL -p NEW_PASSWORD`

#### Delete User

`kfadmin delete user -e USER_EMAIL`

### Namespace (Kubeflow Profile)

#### Create namespace (NSY)

`kfadmin create namespace -n NAMESPACE_NAME -e OWNER_EMAIL`

#### Update namespace owner (NSY)

`kfadmin update namespace owner -n NAMESPACE_NAME -e NEW_OWNER_EMAIL`

#### Update namespace resourceQuota (TBD)

`kfadmin update namespace quota -n NAMESPACE_NAME -e NEW_OWNER_EMAIL`

#### Add user to namespace as contributor (NSY)

`kfadmin add namespace contributor -n NAMESPACE_NAME -e NEW_CONTRIBUTOR_EMAIL`

#### Remove contributor from (NSY)

`kfadmin delete namespace contributor -n NAMESPACE_NAME -e NEW_CONTRIBUTOR_EMAIL`

#### List namespace (NSY)

`kfadmin list namespace`

#### Delete namespace (NSY)

`kfadmin delete namespace -n NAMESPACE_NAME`

### Secret

#### Create Generic Secret (NSY)

#### List Secret (NSY)

#### Delete Secret (NSY)

#### Create Docker Registry Secret (NSY)

#### List Docker Registry Secret (NSY)

#### Set Docker Registry Secret as default (TBD)

#### Delete Docker Registry Secret (TBD)

