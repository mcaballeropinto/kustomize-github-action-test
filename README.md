# Test for issue when running kustomize on a GitHub action
### Description
- When running a [kustomize](https://github.com/kubernetes-sigs/kustomize) configuration that uses a remote resource, it fails with this error:
  ```
  Error: accumulating resources: accumulateFile "accumulating resources from 'git@github.com:mcaballeropinto/kustomize-github-action-test//remote-k8s': evalsymlink failure on '/home/runner/work/kustomize-github-action-test/kustomize-github-action-test/k8s/git@github.com:mcaballeropinto/kustomize-github-action-test/remote-k8s' : lstat /home/runner/work/kustomize-github-action-test/kustomize-github-action-test/k8s/git@github.com:mcaballeropinto: no such file or directory", loader.New "no 'git' program on path: exec: \"git\": executable file not found in $PATH"
  ```
- The error states that it can't find the git executable

### Repo
- You can find a result of the GitHub action execution [here](https://github.com/mcaballeropinto/kustomize-github-action-test/runs/1063731515?check_suite_focus=true)
- Alternatively, you can trigger a new action by pushing against this repo

### Observations
- This issue did not reproduce before 09/01/2020
- This issue does not reproduce locally
- kustomize is built using Go and the line of code that throws the error is [this one](https://github.com/kubernetes-sigs/kustomize/blob/1dc22e3f0d2badac45befa5d58bbb72b7b0fd9e2/api/internal/git/cloner.go#L21)
- There is a Go script in this repo that executes that same function, the file is `test.go`. This script is executed as part of the GitHub test action and it does not return an error
- The test GitHub action also calls `git --version` successfully and finds the path to git using `whereis`
