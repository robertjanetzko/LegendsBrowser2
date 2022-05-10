package main

import (
	"embed"
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"runtime"

	"github.com/pkg/profile"
	"github.com/robertjanetzko/LegendsBrowser2/backend/model"
	"github.com/robertjanetzko/LegendsBrowser2/backend/server"
	"github.com/robertjanetzko/LegendsBrowser2/backend/templates"
	"github.com/spf13/cobra"
)

//go:embed static
var static embed.FS

var (
	f, c, subUri string
	l, p, d, s   *bool
	port         *int
)

var rootCmd = &cobra.Command{
	Use:   "legendsbrowser",
	Short: "A Legends Browser for Dwarf Fortress",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(`
      __                              __        ____                                   
     / /   ___  ____  ___  ____  ____/ /____   / __ )_________ _      __________  _____
    / /   / _ \/ __ \/ _ \/ __ \/ __  / ___/  / __  / ___/ __ \ | /| / / ___/ _ \/ ___/
   / /___/  __/ /_/ /  __/ / / / /_/ (__  )  / /_/ / /  / /_/ / |/ |/ (__  )  __/ /    
  /_____/\___/\__, /\___/_/ /_/\__,_/____/  /_____/_/   \____/|__/|__/____/\___/_/     
             /____/                                                                    ` + "\n ")

		if *p {
			defer profile.Start(profile.ProfilePath(".")).Stop()
			go func() {
				http.ListenAndServe(":8081", nil)
			}()
		}

		config, err := server.LoadConfig(c)
		if err != nil {
			log.Fatal(err)
		}

		templates.DebugTemplates = config.DebugTemplates

		server.DebugJSON = *d
		config.Port = *port
		config.ServerMode = *s
		config.SubUri = subUri

		var world *model.DfWorld

		if *l {
			f = config.LastFile
		}

		if len(f) > 0 {
			w, err := model.Parse(f, nil)
			if err != nil {
				fmt.Println(err)
			} else {
				runtime.GC()
				world = w
			}
		}

		err = server.StartServer(config, world, static)
		if err != nil {
			log.Fatal(err)
		}
	},
}

func main() {
	cobra.MousetrapHelpText = ""
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&f, "world", "w", "", "path to legends.xml")
	rootCmd.PersistentFlags().StringVarP(&c, "config", "c", "", "config file")
	rootCmd.PersistentFlags().StringVarP(&subUri, "subUri", "u", "", "run on /<subUri>")
	l = rootCmd.PersistentFlags().BoolP("last", "l", false, "open last file")
	p = rootCmd.PersistentFlags().BoolP("profile", "P", false, "start profiling")
	d = rootCmd.PersistentFlags().BoolP("debug", "d", false, "show debug data")
	s = rootCmd.PersistentFlags().BoolP("serverMode", "s", false, "run in server mode (disables file chooser)")
	port = rootCmd.PersistentFlags().IntP("port", "p", 58881, "use specific port")
}
