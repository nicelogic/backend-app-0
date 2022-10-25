#!/bin/bash

kubectl --kubeconfig ../../../env/production/token/admin.conf apply -k ./k8s 
