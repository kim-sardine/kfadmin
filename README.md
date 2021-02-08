
# kfadmin : CLI Tool for Kubeflow Admin

- [kfadmin : CLI Tool for Kubeflow Admin](#kfadmin--cli-tool-for-kubeflow-admin)
    - [Prerequitite](#prerequitite)
    - [Confirmed running environment](#confirmed-running-environment)
    - [Commands](#commands)
        - [Users (Only for Dex)](#users-only-for-dex)
            - [Create Dex User](#create-dex-user)
            - [List Dex Users](#list-dex-users)
            - [Delete Dex User](#delete-dex-user)
            - [Change Dex User's password](#change-dex-users-password)
        - [Profile](#profile)
            - [Create profile](#create-profile)
            - [List profile](#list-profile)
            - [Delete profile](#delete-profile)
            - [Add user to profile as contributor](#add-user-to-profile-as-contributor)
            - [Change profile owner (NSY)](#change-profile-owner-nsy)
            - [Remove contributor from profile](#remove-contributor-from-profile)
            - [Update profile resourceQuota (TBD)](#update-profile-resourcequota-tbd)
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
- kubeconfig file (eg. `~/.kube/config`)

> commands like `kubectl get nodes` should be working.

## Confirmed running environment

- kubectl : 1.19.3
- Kubernetes cluster : 1.19.3
- kfctl : 1.2.0 (kfctl_istio_dex.v1.2.0.yaml)

## Commands

> NSY : Not Supported Yet
>
> TBD : To Be Decided

### Users (Only for [Dex](https://www.kubeflow.org/docs/started/k8s/kfctl-istio-dex/))

#### Create Dex User

`kfadmin create user -e USER_EMAIL -p PASSWORD`

#### List Dex Users

`kfadmin list user`

#### Delete Dex User

`kfadmin delete user -e USER_EMAIL`

#### Change Dex User's password

`kfadmin update user password -e USER_EMAIL -p NEW_PASSWORD`

### Profile

#### Create profile

`kfadmin create profile -p PROFILE_NAME -e OWNER_EMAIL`

#### List profile

`kfadmin list profile`

#### Delete profile

`kfadmin delete namespace -p PROFILE_NAME`

#### Add user to profile as contributor

`kfadmin add profile contributor -p PROFILE_NAME -e NEW_CONTRIBUTOR_EMAIL`

#### Change profile owner (NSY)

`kfadmin update profile owner -p PROFILE_NAME -e NEW_OWNER_EMAIL`

#### Remove contributor from profile

`kfadmin delete profile contributor -p PROFILE_NAME -e NEW_CONTRIBUTOR_EMAIL`

#### Update profile resourceQuota (TBD)

### Secret

#### Create Generic Secret (NSY)

#### List Secret (NSY)

#### Delete Secret (NSY)

#### Create Docker Registry Secret (NSY)

#### List Docker Registry Secret (NSY)

#### Set Docker Registry Secret as default (TBD)

#### Delete Docker Registry Secret (TBD)

