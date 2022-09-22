#!/bin/bash

kubectl --kubeconfig ../../../env/production/token/admin.conf apply -k https://github.com/stakater/Reloader/deployments/kubernetes

