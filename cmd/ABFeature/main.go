package main

import (
	"fmt"
	"os"

	"github.com/TerrenceHo/ABFeature"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	logo = `
   _____  ______________________            __                        
  /  _  \ \_   _____/\_   _____/___ _____ _/  |_ __ _________   ____  
 /  /_\  \ |    __)   |    __)/ __ \\__  \\   __\  |  \_  __ \_/ __ \ 
/    |    \|     \    |     \\  ___/ / __ \|  | |  |  /|  | \/\  ___/ 
\____|__  /\___  /    \___  / \___  >____  /__| |____/ |__|    \___  >
        \/     \/         \/      \/     \/                        \/
	`
)

var (
	Version string
)

var mainCmd = &cobra.Command{
	Use:   "abfeature",
	Short: "Feature management made easy.",
	Long:  "Feature management server instance, consuming feature requests to determine if a user is eligible for certain features. If both a flag and configuration variable is set, then the flag takes precedence.",
	Run: func(cmd *cobra.Command, args []string) {
		serve()
	},
}

func main() {
	flags := mainCmd.Flags()

	// Define flags
	flags.BoolP("version", "v", false, "print version of abfeature.")
	flags.StringP("file", "f", "./config/config.yaml", "set path to configuration file.")
	flags.StringP("port", "p", "31337", "set HTTP port of abfeature will run on, between 1024-65535.")
	flags.BoolP("debug", "d", true, "run abfeature in debug mode.")
	flags.Bool("hidebanner", false, "hide banner of router.")
	flags.String("database-engine", "sqlite", "set database engine used for data persistence.")
	flags.String("database-name", "", "set name of database.")
	flags.String("database-user", "", "set user of database.")
	flags.String("database-password", "", "set password to access database.")
	flags.String("database-port", "", "set port database is running on.")
	flags.String("database-host", "", "set host, local or remote database is running on.")

	// Use BindPFlags to pass cobra flags into viper
	viper.BindPFlag("VERSION", flags.Lookup("version"))
	viper.BindPFlag("FILE", flags.Lookup("file"))
	viper.BindPFlag("PORT", flags.Lookup("port"))
	viper.BindPFlag("DEBUG", flags.Lookup("debug"))
	viper.BindPFlag("HIDEBANNER", flags.Lookup("hidebanner"))
	viper.BindPFlag("DATABASE.ENGINE", flags.Lookup("database-engine"))
	viper.BindPFlag("DATABASE.DBNAME", flags.Lookup("database-name"))
	viper.BindPFlag("DATABASE.USER", flags.Lookup("database-user"))
	viper.BindPFlag("DATABASE.PASSWORD", flags.Lookup("database-password"))
	viper.BindPFlag("DATABASE.PORT", flags.Lookup("database-port"))
	viper.BindPFlag("DATABASE.HOST", flags.Lookup("database-host"))

	mainCmd.Execute()
}

func serve() {
	v := viper.GetViper()
	if v.GetBool("VERSION") {
		fmt.Printf("abfeature v%s\n", Version)
		os.Exit(0)
	}

	if !viper.GetBool("HIDEBANNER") {
		fmt.Println(logo)
	}

	filepath := v.GetString("FILE")
	// Read in configuration file
	if filepath == "" {
		fmt.Println("No configuration file found. Defaulting to flags.")
	}
	fmt.Println("Configuration file:", filepath)

	v.SetConfigFile(filepath)
	err := v.ReadInConfig()
	if err != nil {
		fmt.Println("Configuration file could not be found. Defaulting to flags.")
	}

	ABFeature.Start(v)
}
