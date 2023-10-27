

### yaml-sorter

## Utility for sorting the contents of a YAML file

sorts all elements in a YAML file in ascending order
Now preserving comments and file structure

## For example:

This YAML

```yaml
foo: bar

# this is comment for test123
test123: rrr

version:
    services:
        # comment for test
        test: 54

# comment for baz
baz:
    env:
      env1: ffd
      env0: dfdfsd
      testenv: testenv
      another: another
      
# this is comment for test2
test2: 3ww
```

will turn into:

```yaml
# comment for baz
baz:
    env:
        another: another
        env0: dfdfsd
        env1: ffd
        testenv: testenv
foo: bar
# this is comment for test123
test123: rrr
# this is comment for test2
test2: 3ww
version:
    services:
        # comment for test
        test: 54
```


## Build in docker
```bash
docker build -t yaml-sorter .
```

## Using in Docker:

Just output to stdout
```bash
docker run -i --rm  yaml-sorter < input.yml 
```
or edit file in place
```bash
docker run -i --rm  yaml-sorter input.yml
```

