for i in {1..30}; do
  kubectl delete notifier bulk-notifier-$i -n default
done
