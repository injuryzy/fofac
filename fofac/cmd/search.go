/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/injuryzy/fofac/fofac/fetch"
	"github.com/injuryzy/fofac/fofac/log"
	"github.com/spf13/cobra"
)

var (
	page         int8
	size         int
	full         bool
	key          string
	email        string
	query        []string
	before       string
	after        string
	timeInterval int
)

// searchCmd represents the search command
var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "fofac 是一个fofa 的接口工具",
	Long:  `fofac 是一个可以查询范围的接口恐惧,将数据以excel 的方式保存下来`,
	Run: func(cmd *cobra.Command, args []string) {

		if len(query) == 0 {
			log.Error("查询参数: %s, 查询条件不能为空", query)
			return
		}
		var q string
		for i, s := range query {
			if i == len(query)-1 {
				q += fmt.Sprintf("%s", s)
			} else {
				q += fmt.Sprintf("%s", s) + " && "
			}
		}
		search := fetch.FofaSearch{
			FofaQuery: fetch.FofaQuery{
				Page:         page,
				Size:         size,
				Full:         full,
				Key:          key,
				Email:        email,
				Query:        q,
				Before:       before,
				After:        after,
				TimeInterval: timeInterval,
			},
		}
		fmt.Print(search)

		if search.After != "" && search.Before != "" {
			// 已有范围的方式查询
			search.QueryT()
		} else {
			// 直接查询
			search.QueryResult()
		}
	},
}

func init() {
	rootCmd.AddCommand(searchCmd)

	searchCmd.Flags().StringVarP(&key, "key", "k", "", "fofa 用户的key")
	searchCmd.MarkFlagRequired("key")
	searchCmd.Flags().StringVarP(&email, "email", "e", "", "fofa 用户的邮箱")
	searchCmd.MarkFlagRequired("email")
	searchCmd.Flags().Int8VarP(&page, "page", "p", 1, "页码")
	searchCmd.Flags().IntVarP(&size, "size", "s", 10000, "每页条数")
	searchCmd.Flags().BoolVarP(&full, "full", "f", false, "开启一年查询")
	searchCmd.Flags().StringArrayVarP(&query, "query", "q", []string{}, "查询调前，排除after 和 before")
	searchCmd.Flags().StringVarP(&before, "before", "b", "", "在此之前，不包括,时间格式 yyyy-MM-dd")
	searchCmd.Flags().StringVarP(&after, "after", "a", "", "在此之后，不包括,时间格式 yyyy-MM-dd")
	searchCmd.Flags().IntVarP(&timeInterval, "timeInterval", "t", 1, "查询时间间隔，单位为天，默认为1天")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// searchCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// searchCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
