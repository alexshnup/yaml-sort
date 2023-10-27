package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"

	"gopkg.in/yaml.v3"
)

func sortMappingNode(node *yaml.Node) {
	if node.Kind != yaml.MappingNode {
		return
	}

	mapping := make(map[string][]*yaml.Node)

	for i := 0; i < len(node.Content); i += 2 {
		keyNode := node.Content[i]
		valueNode := node.Content[i+1]

		// Recursively sort if valueNode is another MappingNode
		sortMappingNode(valueNode)

		key := keyNode.Value
		mapping[key] = []*yaml.Node{keyNode, valueNode}
	}

	keys := make([]string, 0, len(mapping))
	for k := range mapping {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	sortedNodes := []*yaml.Node{}
	for _, k := range keys {
		sortedNodes = append(sortedNodes, mapping[k]...)
	}

	node.Content = sortedNodes
}

func main() {
	var content []byte
	var err error

	// Check if a filename is provided as an argument
	if len(os.Args) < 2 {
		// Read from STDIN
		scanner := bufio.NewScanner(os.Stdin)
		var sb strings.Builder
		for scanner.Scan() {
			sb.WriteString(scanner.Text())
			sb.WriteString("\n")
		}

		if scanner.Err() != nil {
			fmt.Println("Error reading from STDIN:", scanner.Err())
			return
		}
		content = []byte(sb.String())
	} else {
		filename := os.Args[1]

		// Read the YAML file
		content, err = os.ReadFile(filename)
		if err != nil {
			fmt.Println("Error reading file:", err)
			return
		}
	}

	var rootNode yaml.Node
	err = yaml.Unmarshal(content, &rootNode)
	if err != nil {
		fmt.Printf("Error unmarshaling: %v\n", err)
		return
	}

	// Split top-level nodes into simple and nested
	simpleMap := make(map[string][]*yaml.Node)
	nestedMap := make(map[string][]*yaml.Node)
	if len(rootNode.Content) > 0 {
		rootMapping := rootNode.Content[0]
		for i := 0; i < len(rootMapping.Content); i += 2 {
			keyNode := rootMapping.Content[i]
			valueNode := rootMapping.Content[i+1]

			key := keyNode.Value
			switch valueNode.Kind {
			case yaml.ScalarNode:
				simpleMap[key] = []*yaml.Node{keyNode, valueNode}
			default:
				nestedMap[key] = []*yaml.Node{keyNode, valueNode}
				sortMappingNode(valueNode) // Sort the nested mappings
			}
		}
	}

	// Construct ordered simpleNodes and nestedNodes slices
	sortedNodes := constructSortedNodes(simpleMap)
	sortedNodes = append(sortedNodes, constructSortedNodes(nestedMap)...)

	rootNode.Content[0].Content = sortedNodes

	out, err := yaml.Marshal(&rootNode)
	if err != nil {
		fmt.Printf("Error marshaling: %v\n", err)
		return
	}

	// fmt.Println(string(out))
	if len(os.Args) < 2 {
		// Print sorted content to STDOUT
		fmt.Println(string(out))
	} else {
		filename := os.Args[1]
		// Write the sorted content back to the file
		err := os.WriteFile(filename, out, 0644)
		if err != nil {
			fmt.Println("Error writing sorted content to file:", err)
			return
		}
		fmt.Println("YAML content sorted successfully!")
	}
}

func constructSortedNodes(mapping map[string][]*yaml.Node) []*yaml.Node {
	keys := make([]string, 0, len(mapping))
	for k := range mapping {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	sortedNodes := []*yaml.Node{}
	for _, k := range keys {
		sortedNodes = append(sortedNodes, mapping[k]...)
	}
	return sortedNodes
}
