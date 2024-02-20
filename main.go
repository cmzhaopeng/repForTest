package main

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

type WorkflowConfig struct {
	Name       string   `yaml:"name"`
	StartNode  string   `yaml:"startNode"`
	MiddleNode []string `yaml:"middleNode"`
	EndNode    string   `yaml:"endNode"`
}

func main() {
	// Set configuration file and type
	viper.SetConfigFile("config.yaml")
	viper.SetConfigType("yaml")

	// Read the configuration
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("Error reading config:", err)
		return
	}

	// Decode into a WorkflowConfig struct
	var config WorkflowConfig
	err = viper.Unmarshal(&config)
	if err != nil {
		fmt.Println("Error unmarshalling config:", err)
		return
	}

	// Print extracted information
	fmt.Println("Workflow Name:", config.Name)
	fmt.Println("Start Node:", config.StartNode)
	fmt.Println("End Node:", config.EndNode)
	fmt.Println("Middle Node:", config.MiddleNode)

	// print the middleNode string

	for _, path := range config.MiddleNode {

		//split the path with ","  just as "node1, node2, node3, node4"split the patht ever","  jus

		nodes := strings.Split(path, ", ")

		/*	for _, node := range nodes {
				fmt.Println("  Node:", node)
			}
		*/
		fmt.Println(nodes[0], nodes[1], nodes[2], nodes[3])

		//		fmt.Println("  Path:", path)
	}

	// Print middle node paths (using map access)
	/*	for node, paths := range config.MiddleNode {
			fmt.Println("Middle Node:", node)
			for _, path := range paths {
				//fmt.Println("  Path:", path)

			}
		}
	*/

}
