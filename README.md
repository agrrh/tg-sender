# tg-sender

Universal Telegram sender daemon.

- Consumes events from NATS JetStream topic
- Sends messages to Telegram API

## Sample ArgoCD application

```yaml
---

apiVersion: v1
kind: Secret
metadata:
  name: tg-token
  namespace: tg-my-sender
type: Opaque
data:
  token: base64-encoded-tg-token

---

kind: Application
apiVersion: argoproj.io/v1alpha1
metadata:
  name: tg-my-sender
  namespace: argocd
spec:
  project: default
  source:
    repoURL: https://github.com/agrrh/tg-sender
    targetRevision: master
    path: helm/
    helm:
      parameters:
        - name: app.name
          value: my-bot               # TODO: Change
        - name: app.nats.addr
          value: nats.namespace:4222  # TODO: Change
        - name: app.nats.prefix
          value: my-bot               # TODO: Change
  destination:
    namespace: tg-my-sender
    server: https://kubernetes.default.svc
  syncPolicy:
    automated:
      selfHeal: true
      prune: true
```
