
package main

import (
  "github.com/labstack/echo"
  "github.com/labstack/echo/middleware"
  // bbolt "github.com/etcd-io/bbolt"
	handlers "golang-projects/boltdb_image_server/handlers"
) 

// https://github.com/labstack/echox/tree/master/cookbook/twitter/handler
func main() {
  
  server := echo.New()
  server.Use(middleware.Logger())

  // bbolt read
  server.GET("/get_image", onGetImage)

  // bbolt 
  server.POST("/upload_image", onUploadImage)

  server.Logger.Fatal(server.Start(":8080"))
}

func onGetImage(c echo.Context) error {
  // TODO read from boltdb
  file, err := os.Open("path/to/file")

  // TODO return as http response
  c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

  // return c.HTML(http.StatusOK, "html string")
  // https://echo.labstack.com/guide/response
  // return c.File("path/to/file")
  return c.Stream(http.StatusOK, "image/png", file)
  
}

// https://echo.labstack.com/cookbook/file-upload
func onUploadImage (c echo.Context) error {
  // read form fields
  name := c.FormValue("name")

  //  read form file attachment
  file, err := c.FormFile("file")

  src, err := file.Open()
  defer src.Close()

  // TODO setup boltdb as file destination 

  return c.HTML(http.StatusOK, fmt.Sprint("file upload successful"))
}
