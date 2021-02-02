
# kfadmin : CLI Tool for Kubeflow Admin

- [kfadmin : CLI Tool for Kubeflow Admin](#kfadmin--cli-tool-for-kubeflow-admin)
    - [Prerequitite](#prerequitite)
    - [Commands](#commands)
        - [Users](#users)
            - [Create User](#create-user)
            - [Delete User (NSY)](#delete-user-nsy)
            - [Set user password (NSY)](#set-user-password-nsy)
            - [List Users (NSY)](#list-users-nsy)
        - [Namespace (Kubeflow Profile)](#namespace-kubeflow-profile)
            - [Create namespace (NSY)](#create-namespace-nsy)
            - [Change owner (NSY)](#change-owner-nsy)
            - [Add user to namespace as contributor (NSY)](#add-user-to-namespace-as-contributor-nsy)
            - [List namespace (NSY)](#list-namespace-nsy)
            - [Delete namespace (NSY)](#delete-namespace-nsy)
        - [Secret](#secret)
            - [Create Generic Secret (NSY)](#create-generic-secret-nsy)
            - [List Secret (NSY)](#list-secret-nsy)
            - [Delete Secret (NSY)](#delete-secret-nsy)
            - [Create Docker Registry Secret (NSY)](#create-docker-registry-secret-nsy)
            - [List Docker Registry Secret (NSY)](#list-docker-registry-secret-nsy)
            - [Delete Docker Registry Secret (NSY)](#delete-docker-registry-secret-nsy)

## Prerequitite

- `kubectl`-available environment. (Environment that can run `kubectl get nodes`)
- Kubernetes v1.17.0+ (maybe?) (TBU)

## Commands

> NSY : Not Supported Yet!!

### Users

#### Create User

`kfadmin create user -e USER_EMAIL -p PASSWORD`

#### Delete User (NSY)

`kfadmin delete user -e USER_EMAIL`

#### Set user password (NSY)

`kfadmin set user password -e USER_EMAIL -p NEW_PASSWORD`

#### List Users (NSY)

`kfadmin list user`

### Namespace (Kubeflow Profile)

#### Create namespace (NSY)

`kfadmin create namespace -n NAMESPACE_NAME -e OWNER_EMAIL`

#### Change owner (NSY)

`kfadmin set namespace owner -n NAMESPACE_NAME -e NEW_OWNER_EMAIL`

#### Add user to namespace as contributor (NSY)

`kfadmin add namespace contributor -n NAMESPACE_NAME -e NEW_CONTRIBUTOR_EMAIL`

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

#### Delete Docker Registry Secret (NSY)

