

## yaml-sorter

[![Go](https://github.com/alexshnup/yaml-sort/actions/workflows/go.yml/badge.svg)](https://github.com/alexshnup/yaml-sort/actions/workflows/go.yml)

### Utility for sorting the contents of a YAML file

sorts all elements in a YAML file in ascending order
Now preserving comments and file structure



### For example:

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


### Build in docker
```bash
docker build -t yaml-sorter .
```

### Using in Docker:

Just innput to stdin, and output to stdout
```bash
docker run -i --rm  yaml-sorter < input.yml 
#or
docker run -i --rm  yaml-sorter < input.yml > out.yml
```
or edit file in place
```bash
docker run -i --rm  yaml-sorter input.yml
```


## P.S.
Using yaml.Node for Enhanced YAML Handling
When working with YAML files in Go, the common approach many developers (including myself) take is to unmarshal the YAML content into a native map structure or a custom struct. This works well for many use cases, but it does have its limitations, especially when it comes to preserving and manipulating certain characteristics of the original YAML file, such as comments, order of keys, or intricate nesting.

### The Old Way: Unmarshaling into map or Structs
The standard way of handling YAML typically involves the Unmarshal and Marshal functions. This method is straightforward:

```go
Copy code
data := `
key1: value1
key2: value2
`
m := make(map[string]string)
err := yaml.Unmarshal([]byte(data), &m)
```
While this approach is convenient and intuitive, it's abstracted. You lose out on some of the YAML-specific intricacies. Comments get discarded, the order of keys may not be preserved, and certain manipulations can become complex or clunky.

### The New (Better) Way: Using yaml.Node
Enter yaml.Node. This is a more direct representation of the YAML content. It allows you to work with the YAML in a way that preserves its structure and characteristics. With yaml.Node, you have a tree-like structure that closely mirrors the YAML content, making it possible to preserve comments, order of keys, and perform more nuanced manipulations.

Here's a simple example of unmarshaling into a yaml.Node:

```go
Copy code
var node yaml.Node
err := yaml.Unmarshal([]byte(data), &node)
```
Working directly with yaml.Node structures might seem a bit more complex at first, especially if you're accustomed to the map or struct-based approach. However, the benefits in terms of flexibility and fidelity to the original YAML content make it worth the effort.

### Conclusion
While unmarshaling YAML content into maps or structs is still suitable for many scenarios, if you require greater precision, manipulation capabilities, or the preservation of comments and key order, consider diving into yaml.Node. It's a powerful tool that can greatly enhance how you work with YAML in Go.