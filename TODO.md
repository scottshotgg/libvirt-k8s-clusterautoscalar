# TODO
x create template xml in Go
x API to create new machines

- create new template HDD to copy that starts up and uses `--nodename`
  - pull new latest worker image, ensure script is there
  - squash: `qemu-img convert -c -O qcow2 k8s-worker-arch.qcow2 k8s-worker-arch-comp.qcow2`
  - grab master connection string and put it in script
  - make script that
    - reaches out for hostname : /whoami
    - connects node to master

- generate exact bindings
- Use external gRPC cloudprovider to construct our own "cloud provider"


