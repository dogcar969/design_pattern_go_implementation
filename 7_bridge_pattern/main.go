package main

import "fmt"

type Content interface {
	getContent() string
}

type Json struct {
	json string
}

func (json Json) getContent()string{
return json.json
}

type Image struct {
	filePath string
	image string
}

func (image Image) getContent()string{
	return image.image
	}

func (image *Image) Load() {
	image.image = "content in "+image.filePath 
	fmt.Println("Load from",image.filePath)
}

type Connect interface {
	wrapper(Content)
	send()
}

type Tcp struct {
	header map[string]string
	pack map[string]string
}

func (tcp *Tcp) wrapper(content Content){
	pack := make(map[string]string)
	for k,v := range tcp.header {
		pack[k] = v
	}
	pack["Content"] = content.getContent()
	tcp.pack = pack
}

func (tcp Tcp) send() {
	// 发到控制台
	fmt.Println("tcp receive")
	fmt.Println("key : value")
	for k,v := range tcp.pack {
		fmt.Println(k," : ",v)
	}
}

type Http struct {
	headers map[string]string
	request    bool  // true为请求，false为响应
	statusCode uint8 // 响应状态码，如果是请求为0
	pack       map[string]string
	method     string
}

func (http *Http) getContent() string {
	res := "Http pack\n"
	for k,v := range http.pack {
		res += k + " : " + v + "\n"
	}
	return res
}

func (http *Http) wrapper(content Content) {
	pack := make(map[string]string)
	pack["content"] = content.getContent()
	for k,v := range http.headers {
		pack[k] = v
	}
	if http.request {
		pack["request"] = "true"
	} else {
		pack["request"] = "false"
	}
	pack["method"] = http.method
	pack["statusCode"] = string(http.statusCode)
	http.pack = pack
}

func (http Http) send() {
	// doNothing
}

type Bridge struct {

}

func (bridge Bridge) send(content Content,conn Connect) {
	image,ok := content.(Image) // 图像需要特别添加load步骤
	if ok {
		image.Load()
		if tcp,ok := conn.(*Tcp);ok {
			tcp.wrapper(image)
			tcp.send()
		} else if http,ok := conn.(*Http);ok {
			http.wrapper(image)
			tcp := Tcp{header: make(map[string]string)}
			tcp.header["contentType"] = "Http"
			bridge.send(http,&tcp)
		}
		return 
	}
	if tcp,ok := conn.(*Tcp);ok {
		tcp.wrapper(content)
		tcp.send()
	} else if http,ok := conn.(*Http);ok {
		http.wrapper(content)
		tcp := Tcp{header: make(map[string]string)}
		tcp.header["contentType"] = "Http"
		bridge.send(http,&tcp)
	}
}

func main() {
	tcp := Tcp{header: make(map[string]string)}
	http := Http{headers: make(map[string]string)}
	tcp.header["Source Port"] = "8080"
	// 其他tcp头...
	http.headers["access-control-allow-origin"] = "*"
	// 其他http头...
	bridge := Bridge{}
	bridge.send(Image{filePath: "dog.png"},&tcp)
	bridge.send(Image{filePath: "dog.png"},&http)
	bridge.send(Json{json: "{\"msg\":\"pong\"}"},&tcp)
	bridge.send(Json{json: "{\"msg\":\"pong\"}"},&http)
}