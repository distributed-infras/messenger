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
hub clone http://github.com/distributed-infras/messenger
cd messenger
git checkout -b feature/<your-feature-version>
# make changes
git commit -am "comment on this feature"
hub fork --remote-name origin
git push origin feature/<your-feature-version>
hub pull-request
```



