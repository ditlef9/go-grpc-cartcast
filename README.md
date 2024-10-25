
# Go gRPC - CartCast

This is a demonstration of Go gRPC protocol. 
The demo has a server and client part that sends data to each other.

The client is a futuristic shopping cart that has a GPS-system and a small screen on it.
When the cart sends its coordinates to a server, which then sends back a commercial for
the items near that shelf, 


<pre>
|------------------|                    |--------------|
|   Client         |                    |    Server    |
|  Shopping cart   | -- Coordinates --> |              |
|                  | <-- Commercial --  |              |
|------------------|                    |--------------|
</pre>


## 🔥 Cool Features


## 🚀 Running Locally (Windows)

1. **Install Go:** https://golang.org/doc/install

2. **Install TDM-CCC:** Download 64+32-bit MinGW-w64 edition (https://jmeubank.github.io/tdm-gcc/download/) (Requires restart)

3. **Install Protoc:** Download latest zip, example `protoc-28.3-win64.zip` from https://github.com/protocolbuffers/protobuf/releases<br>
   Extract the zip file to `C:\Program Files\Protoc`.<br>
   Open Settings > System > About > Advanced system settings > Enviroment variables > 
   System variables: Path [Edit] > [New] > `C:\Program Files\Protoc\bin`<br>
   Restart Windows<br>
   Run `protoc --version` to check that it is installed.

4. **Add openssl to Path:**
   Open Settings > System > About > Advanced system settings > Environment variables >
   System variables: Path [Edit] > [New] > `C:\Program Files\Git\usr\bin\`<br>
   Restart Windows<br>
   Run `openssl --version` to check that it is installed.

5. **Install Evans CLI:** Download latest zip, example `evans_windows_amd64.tar.gz` from https://github.com/ktr0731/evans/releases/tag/
   Extract the zip file to `C:\Program Files\Evans`.<br>
   Open Settings > System > About > Advanced system settings > Enviroment variables >
   System variables: Path [Edit] > [New] > `C:\Program Files\Evans`<br>
   Restart Windows<br>
   Run `evans --version` to check that it is installed.

6. **Install Docker:** Download and install: https://docs.docker.com/get-started/get-docker/<br>
   Run ` docker --version` and ` docker-compose --version` to check that it is installed.<br>
   Run `Get-Service -Name *docker*`

7. **Clone repository:** the repository and navigate to the project folder.

8. **Install Go plugins for the protocol compiler:**
```bash
go get go.mongodb.org/mongo-driver/v2/mongo
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

9. **Install dependencies:** Run `go mod tidy` to install dependencies.


## 🛠️ Deploy to Google Cloud Run

1. **Authenticate with GCP:** `gcloud auth login`
2. **Build and deploy:**
    ```bash
    gcloud builds submit --tag gcr.io/[PROJECT_ID]/product-api
    gcloud run deploy --image gcr.io/[PROJECT_ID]/product-api --platform managed
    ```

3. **Access API:** The API will be available via the generated GCP URL.


## Run App

Open three terminals:

1. **Database: **<br>
```
cd cartcast
docker-compose up
```

If error then stop and delete containers:
```
docker stop $(docker ps -a -q)
docker rm $(docker ps -a -q)
```

2. **Server**<br>
```
.\make-cartcast.ps1
.\bin\cartcast\server.exe
```

3. **Client**<br>

```
.\bin\cartcast\client.exe
```

4. **Check DB**<br>
http://localhost:8081 (admin:pass)<br>
cartcastdb > cartcastdb

## 💻 Developer Notes

**Initializing the Project:**
```bash
go mod init ekeberg.com/go-grpc-carcast
```


## 📖 License

This project is licensed under the
[Apache License 2.0](https://www.apache.org/licenses/LICENSE-2.0).

```
Copyright 2024 github.com/ditlef9

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
```