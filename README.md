
# kfadmin : CLI Tool for Kubeflow Admin

- [kfadmin : CLI Tool for Kubeflow Admin](#kfadmin--cli-tool-for-kubeflow-admin)
    - [Prerequitite](#prerequitite)
    - [Confirmed running environment](#confirmed-running-environment)
    - [Commands](#commands)
        - [Users](#users)
            - [Create User](#create-user)
            - [List Users](#list-users)
            - [Delete User](#delete-user)
            - [Change user password](#change-user-password)
        - [Profile](#profile)
            - [Create profile](#create-profile)
            - [List profile](#list-profile)
            - [Delete profile](#delete-profile)
            - [Change profile owner (NSY)](#change-profile-owner-nsy)
            - [Add user to profile as contributor (NSY)](#add-user-to-profile-as-contributor-nsy)
            - [Remove contributor from (NSY)](#remove-contributor-from-nsy)
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

#### Delete User

`kfadmin delete user -e USER_EMAIL`

#### Change user password

`kfadmin update user password -e USER_EMAIL -p NEW_PASSWORD`

### Profile

#### Create profile

`kfadmin create profile -p PROFILE_NAME -e OWNER_EMAIL`

#### List profile

`kfadmin list profile`

#### Delete profile

`kfadmin delete namespace -p PROFILE_NAME`

#### Change profile owner (NSY)

`kfadmin update profile owner -p PROFILE_NAME -e NEW_OWNER_EMAIL`

#### Add user to profile as contributor (NSY)

`kfadmin add profile contributor -n PROFILE_NAME -e NEW_CONTRIBUTOR_EMAIL`

#### Remove contributor from (NSY)

`kfadmin delete profile contributor -n PROFILE_NAME -e NEW_CONTRIBUTOR_EMAIL`

#### Update profile resourceQuota (TBD)

### Secret

#### Create Generic Secret (NSY)

#### List Secret (NSY)

#### Delete Secret (NSY)

#### Create Docker Registry Secret (NSY)

#### List Docker Registry Secret (NSY)

#### Set Docker Registry Secret as default (TBD)

#### Delete Docker Registry Secret (TBD)

