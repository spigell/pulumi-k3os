# Simple k3os cluster

## Craate
```bash
$ ./vagrant_prepare.sh
$ npm i
$ pulumi stack init test
$ pulumi up
```

## Destroy
```bash
$ pulumi destroy
$ vagrant destroy --force
```
