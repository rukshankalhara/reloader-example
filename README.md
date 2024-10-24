# reloader-example

These code samples go paired with [this video](https://youtu.be/1vodLYj9yy8).

## Install using Helm

```bash
helm repo add stakater https://stakater.github.io/stakater-charts
helm repo update
helm install reloader -n reloader --create-namespace stakater/reloader
```

You can find the original installation instructions [here](https://github.com/stakater/Reloader?tab=readme-ov-file#deploying-to-kubernetes)!