package main

import (
	"github.com/goravel/framework/facades"
	"goravel/bootstrap"
)

/*var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}*/

func main() {

	/*go func() {
		router := gin.Default()
		router.Use(cors.Default())

		router.GET("/websockets/process/logs", func(c *gin.Context) {
			conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
			if err != nil {
				return
			}
			defer conn.Close()

			for {
				file, err := os.Open(facades.Storage().Disk("public").Path("logs\\logs.log"))

				if err == nil {
					scanner := bufio.NewScanner(file)
					var lineas []string

					for scanner.Scan() {
						if scanner.Text() != "" {
							lineas = append(lineas, fmt.Sprintf("%s", scanner.Text()))
						}
					}
					fmt.Println(lineas[len(lineas)-1])
					conn.WriteMessage(websocket.TextMessage, []byte(lineas[len(lineas)-1]))

				}
				time.Sleep(time.Second)
			}
		})
		router.Run(":8080")
	}()*/

	// This bootstraps the framework and gets it ready for use.
	bootstrap.Boot()

	//Start http server by facades.Route().
	go func() {
		if err := facades.Route().Run(); err != nil {
			facades.Log().Errorf("Route run error: %v", err)
		}
	}()

	select {}

}
