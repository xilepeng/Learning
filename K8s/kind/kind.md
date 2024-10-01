
## Warning: You are using macOS 12.
We (and Apple) do not provide support for this old version.

``` shell 
brew install go
brew install kind
brew install podman

brew install nerdctl
```




``` go
go install sigs.k8s.io/kind@v0.24.0 && kind create cluster

```