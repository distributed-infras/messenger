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



