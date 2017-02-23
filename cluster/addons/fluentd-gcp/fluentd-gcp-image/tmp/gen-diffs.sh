#!/bin/bash

case $1 in
  old)
    IMAGE=gcr.io/google_containers/fluentd-gcp:1.40
    GEMCMD=td-agent-gem
    ;;
  new)
    IMAGE=gcr.io/stclair-k8s-ext/fluentd-gcp:dev
    GEMCMD=gem
    ;;
  *)
    echo "Expected \$1 to be one of {old,new}"
    exit 1
esac

# apt_pkgs
docker run --rm $IMAGE sh -c "apt list --installed 2> /dev/null | tail -n +2 | cut -f1,2 -d' ' | sort" > apt_pkgs_versioned.txt
cat apt_pkgs_versioned.txt | cut -f1 -d/ > apt_pkgs_unversioned.txt

# gems_versioned.txt
docker run --rm $IMAGE sh -c "$GEMCMD list --local | sort" > gems_versioned.txt
cat gems_versioned.txt | cut -f1 -d' ' > gems_unversioned.txt
