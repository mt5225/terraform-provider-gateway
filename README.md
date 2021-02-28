# terraform provider for cloudj api gateway

## test and build with ci

```
fly -t local set-pipeline --pipeline provider-build --config ./ci/build.yml
```

## run mock server

```
cd ./mock
pipenv install
pipenv shell
python mock.py
```
