/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"net/http"
	"os"
	"io"
	"github.com/spf13/cobra"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "A brief description of your command",
	Long: `This get command will call Github repository in order to return the desired Gopher`,
	Run: func(cmd *cobra.Command, args []string) {
		var gopherName = `dr-who`
		if len(args)>=1 && args[0] != ""{
			gopherName =args[0]

		}
		URL:="https://github.com/scraly/gophers/raw/main/"+gopherName +".png"
		fmt.Println("Try to get '"+gopherName+"'Gopher...")
		response,err:=http.Get(URL)
		if err!=nil{
			fmt.Println(err)
		}
		defer response.Body.Close()
		if response.StatusCode ==200{
			//create The file 
			out,err:=os.Create(gopherName + ".png")
			if err!=nil{
				fmt.Println(err)
			}
			defer out.Close()

			_,err=io.Copy(out,response.Body)
			if err!=nil{
				fmt.Println(err)
			}
			fmt.Println("Perfect! Just saved in " + out.Name() +"!")

		}else{
			fmt.Println("Error: "+gopherName +"not exists! :-(")
		}
	},
}

func init() {
	rootCmd.AddCommand(getCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
