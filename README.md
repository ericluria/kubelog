# kubelog

## Prerequisites
```
brew install kubectl

// Make sure to read the brew output for instructions on finishing fzf setup
brew install fzf
```

## Usage
Create a symlink:
```
ln -s $(pwd)/kubelog /usr/local/bin/kubelog
```

Use it!
```
// Start a container
kubectl run nginx --image nginx

// Tail the logs
kubelog
```

