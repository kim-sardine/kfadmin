
# kfadmin : CLI Tool for Kubeflow Admin

- [kfadmin : CLI Tool for Kubeflow Admin](#kfadmin--cli-tool-for-kubeflow-admin)
    - [Prerequitite](#prerequitite)
    - [Confirmed running environment](#confirmed-running-environment)
    - [Commands](#commands)
        - [Get](#get)
            - [Get Dex Users](#get-dex-users)
            - [Get profile](#get-profile)
            - [Get secrets (NSY)](#get-secrets-nsy)
        - [Create](#create)
            - [Create Dex User](#create-dex-user)
            - [Create profile](#create-profile)
            - [Create contributor using existing user and profile](#create-contributor-using-existing-user-and-profile)
            - [Create environment variable using opaque secret (NSY)](#create-environment-variable-using-opaque-secret-nsy)
        - [Update](#update)
            - [Change Dex User's password](#change-dex-users-password)
            - [Change profile owner](#change-profile-owner)
        - [Delete](#delete)
            - [Delete Dex User](#delete-dex-user)
            - [Delete profile](#delete-profile)
            - [Remove contributor from profile](#remove-contributor-from-profile)
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

`kfadmin get dex-users`

#### Get profile

`kfadmin get profiles`

#### Get secrets (NSY)

### Create

#### Create Dex User

`kfadmin create dex-user -e USER_EMAIL -p PASSWORD`

- flags
    - `--restart-dex`
        - Restart `dex` deployment after updating ConfigMap to reflect changes
        - Without this option, you have to run `kubectl rollout restart deployment dex -n auth` to manually reflect changes
    - `-y` (TBD)

#### Create profile

`kfadmin create profile -p PROFILE_NAME -e OWNER_EMAIL`

#### Create contributor using existing user and profile

`kfadmin create contributor -p PROFILE_NAME -e USER_EMAIL`

#### Create environment variable using opaque secret (NSY)

### Update

#### Change Dex User's password

`kfadmin update dex-user password -e USER_EMAIL -p NEW_PASSWORD`

#### Change profile owner

`kfadmin update profile owner -p PROFILE_NAME -e NEW_OWNER_EMAIL`

### Delete

#### Delete Dex User

`kfadmin delete dex-user -e USER_EMAIL`

#### Delete profile

`kfadmin delete profile -p PROFILE_NAME`

#### Remove contributor from profile

`kfadmin delete contributor -p PROFILE_NAME -e USER_EMAIL`

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
