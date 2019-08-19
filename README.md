
### reference boilerplate
basic
https://github.com/Amatobahn/echo-boilerplate/blob/master/server/server.go

with jwt auth
https://github.com/SeptiyanAndika/go-echo-boilerplate/blob/master/utils/authorizer.go

### TODO
- [] create flutter client that crops and uploads an image 

### Basic Schema

- KeyId bson.ObjectId
- Payload string
- CreatedAt time.Time 

### Image Schema

- KeyId string `bson.ObjectId`
- Payload blob.binary `image/jpeg/png`


### Aws3 service client
  a := &handler.Aws3{}
  a.Bucket = bucket
  a.Region = region
  a.Init(bucket, region)
  name := shortuuid.New()
  // read file from local dir
  image_byte, err := ioutil.ReadFile("test.png")
  if err != nil {log.Fatal(err)}
  err = a.PutImage(image_byte, name) // Aws3.PutObject(data *os.File, fileName string)
  if err != nil {
    log.Fatal("AWS S3 put error : ",err)
  }
  obj, err := a.GetImage(name)
  if err != nil {
    log.Fatal("AWS S3 get error : ",err)
  }
  log.Println("SUCCESS S3 RESULT ",obj)