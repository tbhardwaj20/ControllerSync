for i in {1..30}; do
  cat <<EOF | kubectl apply -f -
apiVersion: core.tbh.dev/v1
kind: Notifier
metadata:
  name: bulk-notifier-$i
  namespace: default
spec:
  message: "Hello from bulk-notifier-$i"
  target: "default-target"
  type: "info"
EOF
done