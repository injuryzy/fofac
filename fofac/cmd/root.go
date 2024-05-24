/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

//jensenhuang455@gmail.com
//6c72716f741f878b25547bed0bdb716f

var (
	page   int8
	size   int16
	full   bool
	key    string
	email  string
	query  string
	before string
	after  string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "fofac",
	Short: "fofa 是一个fofa 的接口工具，将数据以excel 的方式保存下来",
	Long:  `todo`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {

	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.fofac.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().Int8VarP(&page, "page", "p", 1, "页码")
	rootCmd.Flags().Int16VarP(&size, "size", "s", 10000, "每页条数")
	rootCmd.Flags().BoolVarP(&full, "full", "f", false, "开启一年查询")
	rootCmd.Flags().StringVarP(&key, "key", "k", "", "fofa 用户的key")
	rootCmd.Flags().StringVarP(&email, "email", "e", "", "fofa 用户的邮箱")
	rootCmd.Flags().StringVarP(&query, "query", "q", "", "查询调前，排除after 和 before")
	rootCmd.Flags().StringVarP(&before, "before", "b", "", "在此之前，不包括,时间格式 yyyy-MM-dd")
	rootCmd.Flags().StringVarP(&after, "after", "a", "", "在此之后，不包括,时间格式 yyyy-MM-dd")
}
