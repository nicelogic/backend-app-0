#!/bin/bash

while getopts 'hdbiu:' OPTION; do
	case "$OPTION" in
	h)
		echo "script usage: $(basename \$0) [-h(help)] [-b(build)] [-d(deploy)] [-i(deploy-ingress] [-i(update)]" >&2
		;;
	b)
		echo "begin build docker image"
		pwd="$PWD"
		echo $pwd
		cd ../../devops/build
		python3.9 build.py $pwd
		;;
	d)
		echo "begin deploy k8s"
		pwd="$PWD"
		echo $pwd
		cd ../../devops/deploy
		python3.9 deploy.py $pwd
		;;
	i)
		echo "begin deploy k8s ingress-route"
		pwd="$PWD"
		echo $pwd
		cd ../../devops/deploy
		deploy-ingress.sh $pwd
		;;
	u)
		echo "begin update k8s"
		pwd="$PWD"
		echo $pwd
		cd ../../devops/deploy
		update.sh $pwd
		;;
	?)
		echo "script usage: $(basename \$0) [-h(help)] [-b(build)] [-d(deploy)] [-i(deploy-ingress] [-i(update)]" >&2
		exit 1
		;;
	esac
done
shift "$(($OPTIND - 1))"
