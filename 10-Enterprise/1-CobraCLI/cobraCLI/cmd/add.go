/*
Copyright © 2022 Liam Hogan

*/
package cmd

import (
        "fmt"
        "strconv"
        "github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
        Use:   "add",
        Short: "A brief description of your command",
        Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
        Run: func(cmd *cobra.Command, args []string) {
                sum := 0
        for _ , args := range args {
            num, err :=  strconv.Atoi(args)
        
        if err != nil {
            fmt.Println(err)
        }
        sum = sum + num
    }
        fmt.Println("result of addition is", sum)

        },
}

func init() {
        rootCmd.AddCommand(addCmd)

        // Here you will define your flags and configuration settings.

        // Cobra supports Persistent Flags which will work for this command
        // and all subcommands, e.g.:
        // addCmd.PersistentFlags().String("foo", "", "A help for foo")

        // Cobra supports local flags which will only run when this command
        // is called directly, e.g.:
        // addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

