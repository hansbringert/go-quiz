# go-quiz

A project containing tasks needed for creating an executable and distribute it using brew.

I am using 
    
- goreleaser

## Build a `snapshot`
    
    make snapshot

## Build and distribute a version of the software
    
- Create a new release
    
    
    git tag v0.0.0 
    make release

- Rebuild a release


    git tag -d v0.0.0
    git push --delete origin v0.0.0
    make release
    
  
## Download and install using brew
    
    brew tap hansbringert/core
    brew install go-quiz