
# kfadmin : CLI Tool for Kubeflow Admin

- [kfadmin : CLI Tool for Kubeflow Admin](#kfadmin--cli-tool-for-kubeflow-admin)
    - [Prerequitite](#prerequitite)
    - [Confirmed running environment](#confirmed-running-environment)
    - [Commands](#commands)
        - [Get](#get)
            - [Get Dex Users](#get-dex-users)
            - [Get profile](#get-profile)
            - [Get Secret (NSY)](#get-secret-nsy)
            - [Get Docker Registry Secret (NSY)](#get-docker-registry-secret-nsy)
        - [Create](#create)
            - [Create Dex User](#create-dex-user)
            - [Create profile](#create-profile)
            - [Create Generic Secret (NSY)](#create-generic-secret-nsy)
            - [Create Docker Registry Secret (NSY)](#create-docker-registry-secret-nsy)
        - [Add](#add)
            - [Add user to profile as contributor](#add-user-to-profile-as-contributor)
        - [Update](#update)
            - [Change Dex User's password](#change-dex-users-password)
            - [Change profile owner](#change-profile-owner)
            - [Update profile resourceQuota (TBD)](#update-profile-resourcequota-tbd)
            - [Set Docker Registry Secret as default (TBD)](#set-docker-registry-secret-as-default-tbd)
        - [Delete](#delete)
            - [Delete Dex User](#delete-dex-user)
            - [Delete profile](#delete-profile)
            - [Delete Secret (NSY)](#delete-secret-nsy)
            - [Remove contributor from profile](#remove-contributor-from-profile)
            - [Delete Docker Registry Secret (TBD)](#delete-docker-registry-secret-tbd)
    - [Auto Completion](#auto-completion)
        - [bash](#bash)
        - [zsh](#zsh)
        - [fish](#fish)
        - [PowerShell](#powershell)

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

### Get

#### Get Dex Users

`kfadmin get users`

#### Get profile

`kfadmin get profiles`

#### Get Secret (NSY)

#### Get Docker Registry Secret (NSY)

### Create

#### Create Dex User

`kfadmin create user -e USER_EMAIL -p PASSWORD`

- flags
    - `--restart-dex` (NSY)
        - Restart `dex` deployment after updating ConfigMap to reflect changes
        - Without this option, you have to run `kubectl rollout restart deployment dex -n auth` to manually reflect changes
    - `-y` (TBD)

#### Create profile

`kfadmin create profile -p PROFILE_NAME -e OWNER_EMAIL`

#### Create Generic Secret (NSY)

#### Create Docker Registry Secret (NSY)

### Add

#### Add user to profile as contributor

`kfadmin add profile contributor -p PROFILE_NAME -e NEW_CONTRIBUTOR_EMAIL`

### Update

#### Change Dex User's password

`kfadmin update user password -e USER_EMAIL -p NEW_PASSWORD`

#### Change profile owner

`kfadmin update profile owner -p PROFILE_NAME -e NEW_OWNER_EMAIL`

#### Update profile resourceQuota (TBD)

#### Set Docker Registry Secret as default (TBD)

### Delete

#### Delete Dex User

`kfadmin delete user -e USER_EMAIL`

#### Delete profile

`kfadmin delete namespace -p PROFILE_NAME`

#### Delete Secret (NSY)

#### Remove contributor from profile

`kfadmin delete profile contributor -p PROFILE_NAME -e NEW_CONTRIBUTOR_EMAIL`

#### Delete Docker Registry Secret (TBD)

## Auto Completion

### bash

```bash
$ source <(kfadmin completion bash)

# To load completions for each session, execute once:
# Linux:
$ kfadmin completion bash > /etc/bash_completion.d/kfadmin
# macOS:
$ kfadmin completion bash > /usr/local/etc/bash_completion.d/kfadmin
```

### zsh

```bash
# If shell completion is not already enabled in your environment,
# you will need to enable it.  You can execute the following once:
$ echo "autoload -U compinit; compinit" >> ~/.zshrc

# To load completions for each session, execute once:
$ kfadmin completion zsh > "${fpath[1]}/_kfadmin"

# You will need to start a new shell for this setup to take effect.
```

### fish

```bash
$ kfadmin completion fish | source

# To load completions for each session, execute once:
$ kfadmin completion fish > ~/.config/fish/completions/kfadmin.fish
```

### PowerShell

```bash
PS> kfadmin completion powershell | Out-String | Invoke-Expression

# To load completions for every new session, run:
PS> kfadmin completion powershell > kfadmin.ps1
# and source this file from your PowerShell profile.
```
