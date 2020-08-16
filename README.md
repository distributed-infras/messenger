# messenger

# sandbox

Purpose: use this to create your sand castles.  Be ready to display
them and showcase your creations.

## example

You can create your personal folders under sandbox.

```sandbox/frankieliu/README.md```

as an example.

## hiearchy and README.md

1. choose a understandable hierachy for your files
2. add at a minimum a top level README.md file
3. document before implementing (not the other way around)

# forking and creating pull requests

```bash
# Clones the repo
hub clone http://github.com/distributed-infras/messenger
cd messenger

# Create a new branch for your work 
git checkout -b feature/<your-feature-name>

# add and commit changes
git commit -am "comment on this feature"

# Line below is optional (advise: don't fork)
# hub fork --remote-name origin

# push your work to origin
git push origin feature/<your-feature-name>

# create a pull-request
hub pull-request
```

# starting off

## Generate a ssh keypair

```bash
ssh-keygen
```

For the most part you can use the default values.  But if you already have
a `~/.ssh/id_rsa` then you should rename it to something else.

Leave the passphrase empty.

## Paste on github

```bash
cat ~/.ssh/<name-you-gave-above>.pub
```

In github, login to your account, click on account settings, go to SSH keys and add your key under `Add an SSH Key`.

Click on `Add key`.

## Test it
To test it you can use

```bash
ssh -T git@github.com
```

It should return with your github user account.

## Change the origin

Copy the the SSH clone of the repository.  And set the origin to this.

```bash
git remote set origin git@github.com:distributed-infras/messenger.git
```

Now you should be able to push using your ssh-key just generated.

