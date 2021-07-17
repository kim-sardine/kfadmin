
# kfadmin : CLI Tool for Kubeflow Admin

- [kfadmin : CLI Tool for Kubeflow Admin](#kfadmin--cli-tool-for-kubeflow-admin)
  - [Prerequitite](#prerequitite)
  - [Running environment](#running-environment)
  - [Commands](#commands)
    - [Get](#get)
      - [Get Dex Static Users](#get-dex-static-users)
      - [Get Profiles](#get-profiles)
      - [Get Contributors (TBU)](#get-contributors-tbu)
    - [Create](#create)
      - [Create Dex Static User](#create-dex-static-user)
      - [Create profile](#create-profile)
      - [Create contributor using existing user and profile](#create-contributor-using-existing-user-and-profile)
    - [Update](#update)
      - [Change Dex Static User's password](#change-dex-static-users-password)
      - [Change profile owner](#change-profile-owner)
    - [Delete](#delete)
      - [Delete Dex Static User](#delete-dex-static-user)
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

## Running environment

- kubectl : v1.19.3
- Kubernetes cluster : v1.19.3
- kfctl : 1.2.0 (kfctl_istio_dex.v1.2.0.yaml)

## Commands

> NSY : Not Supported Yet
>
> TBD : To Be Decided

### Get

#### Get Dex Static Users

> Only when you're using `staticPasswords`

`kfadmin get staticusers`

#### Get Profiles

- `kfadmin get profiles`
  - Get all profiles
- `kfadmin get profiles -e EMAIL` (TBU)
  - Get profiles that containing a specific user

#### Get Contributors (TBU)

- `kfadmin get contributors -p PROFILE`
  - Get all contributors included in a specific profile

### Create

#### Create Dex Static User

> Only when you're using `staticPasswords`

`kfadmin create staticuser -e USER_EMAIL -p PASSWORD`

- flags
    - `-r`, `--restart-dex`
        - Restart `dex` deployment after updating ConfigMap to reflect changes
        - Without this option, you have to run `kubectl rollout restart deployment dex -n auth` to manually reflect changes
    - `-y` (TBD)

#### Create profile

`kfadmin create profile -p PROFILE_NAME -e OWNER_EMAIL`

#### Create contributor using existing user and profile

`kfadmin create contributor -p PROFILE_NAME -e USER_EMAIL`

### Update

#### Change Dex Static User's password

`kfadmin update staticuser password -e USER_EMAIL -p NEW_PASSWORD`

#### Change profile owner

`kfadmin update profile owner -p PROFILE_NAME -e NEW_OWNER_EMAIL`

### Delete

#### Delete Dex Static User

`kfadmin delete staticuser -e USER_EMAIL`

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
