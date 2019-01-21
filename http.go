package common

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	c_url "net/url"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type myRegexp struct {
	*regexp.Regexp
}

/**
* 写入cookie
 */

func InsertCookie(w http.ResponseWriter, doname string, key string, val string, exptime int) {
	name := key                                                     //cookie名称
	value := val                                                    //cookie的值
	path := "/"                                                     //作用目录
	expires := time.Now().Add(time.Second * time.Duration(exptime)) //失效时间，单位秒
	maxAge := exptime                                               //新的失效方式，单位秒
	secure := false                                                 //安全cookie
	httpOnly := true                                                //限制http访问

	computerCookie := &http.Cookie{Name: name, Value: value, Path: path, Domain: doname, Expires: expires, MaxAge: maxAge, Secure: secure, HttpOnly: httpOnly}
	http.SetCookie(w, computerCookie)
}

/**
* 读取cookie
 */
func ReadCookie(r *http.Request, key string) string {
	res := ""
	val, err := r.Cookie(key)
	if err == nil {
		res = val.Value
	}
	return res
}

/**
* 获得计算机标识=md5(注册IP+浏览器+操作系统os+当前时间戳)
* return string 获得计算机唯一标识
 */
func GetComputer(IP string, bb string, oo string) string {
	computer := GetMd5(IP + bb + oo + strconv.FormatInt(time.Now().Unix(), 10))
	return computer
}

/**
* 根据传入的用户访问头，判断浏览器类型
* @use_agent string  用户传递过来的header中的user-agent
* return string,string  返回浏览器类型,操作系统类型
 */

func GetBrowserOS(use_agent string) (string, string) {
	OS := "other"
	if strings.Index(use_agent, "Windows") > -1 || strings.Index(use_agent, "win") > -1 {
		var re = myRegexp{regexp.MustCompile(`[Windows|win] (.*?)[;|)| ].*?`)}
		res := re.FindStringSubmatch(use_agent)
		OS := "Windows"
		if len(res) > 1 {
			winVersion := ""
			if strings.Index(res[1], "95") > -1 {
				winVersion = "95"
			}
			if strings.Index(res[1], "98") > -1 {
				winVersion = "98"
			}
			if strings.Index(res[1], "4.9") > -1 {
				winVersion = "ME"
			}
			if strings.Index(res[1], "NT 4.0") > -1 {
				winVersion = "NT 4.0"
			}
			if strings.Index(res[1], "NT 5.0") > -1 {
				winVersion = "2000"
			}
			if strings.Index(res[1], "NT 5.1") > -1 {
				winVersion = "XP"
			}
			if strings.Index(res[1], "NT 5.2") > -1 {
				winVersion = "Server 2003"
			}
			if strings.Index(res[1], "NT 6.0") > -1 {
				winVersion = "Vista"
			}
			if strings.Index(res[1], "NT 6.1") > -1 {
				winVersion = "7"
			}
			if strings.Index(res[1], "NT 6.2") > -1 {
				winVersion = "8"
			}
			if strings.Index(res[1], "NT 6.3") > -1 {
				winVersion = "Server 2012"
			}
			if strings.Index(res[1], "NT 10.0") > -1 {
				winVersion = "10"
			}
			OS = OS + " " + winVersion
		}
	}
	if strings.Index(use_agent, "Mac") > -1 {
		OS = "Mac OS"
	}
	if strings.Index(use_agent, "Linux") > -1 {
		OS = "Linux"
	}
	if strings.Index(use_agent, "Unix") > -1 {
		OS = "Unix"
	}
	if strings.Index(use_agent, "Sun") > -1 {
		OS = "Sun OS"
	}
	if strings.Index(use_agent, "ibm") > -1 {
		OS = "Ibm"
	}
	if strings.Index(use_agent, "PowerPC") > -1 {
		OS = "PowerPC"
	}
	if strings.Index(use_agent, "AIX") > -1 {
		OS = "AIX"
	}
	if strings.Index(use_agent, "HPUX") > -1 {
		OS = "HPUX"
	}
	if strings.Index(use_agent, "BSD") > -1 {
		OS = "BSD"
	}

	if strings.Index(use_agent, "MSIE") > -1 {
		var re = myRegexp{regexp.MustCompile(`MSIE (.*?)[;|)| ].*?`)}
		res := re.FindStringSubmatch(use_agent)
		browser := "Internet Explorer"
		if len(res) > 1 {
			browser = browser + " " + res[1]
		}
		return browser, OS
	}
	if strings.Index(use_agent, "Firefox") > -1 {
		return "Firefox", OS
	}
	if strings.Index(use_agent, "Chrome") > -1 {
		return "Chrome", OS
	}
	if strings.Index(use_agent, "Opera") > -1 {
		return "Opera", OS
	}
	if strings.Index(use_agent, "Safari") > -1 {
		return "Safari", OS
	}
	return "other", OS
}

/**
* 模拟http请求
* @url		string  需要抓取的url
* @method	string	请求方式，POST，GET，PUT等
* @body		string	需要传递的值
* @charset	string	字符编码
* @contentType string 定义http请求的文档格式，默认string
* @headSet	map[string] string	请求需要带上的头
 */
func HttpRequest(url, method, body, charset, contentType string, headSet map[string]string) (string, int) {
	src := ""
	httpStart := false

	statusCode := 101

	if len(charset) < 1 {
		charset = "utf-8"
	}
	if len(method) < 1 {
		method = "POST"
	}
	sContentType := "text/html"
	if method == "POST" {
		sContentType = "application/x-www-form-urlencoded"
	}
	if len(contentType) < 1 {
		contentType = sContentType
	}

	var req *http.Request
	var err error
	var ioread io.Reader

	if len(body) > 0 {
		ioread = strings.NewReader(body)
	}

	//连续10次尝试
	for i := 0; i < 10; i++ {
		req, err = http.NewRequest(method, url, ioread)
		if err == nil {
			//如果成功了，跳出
			httpStart = true
			break
		}
	}
	//如果10次尝试后，都不能连接
	if err != nil {
		LogsWithFileName("", "http_error", url+err.Error())
		return err.Error(), 503
	}

	//使用完连接后关闭
	if req.Body != nil {
		defer req.Body.Close()
	}

	//执行连接
	tr := &http.Transport{DisableKeepAlives: true,
		Dial: func(netw, addr string) (net.Conn, error) {
			c, err := net.DialTimeout(netw, addr, time.Second*60) //设置建立连接超时
			if err != nil {
				return nil, err
			}
			c.SetDeadline(time.Now().Add(30 * time.Second))     //设置发送接收数据超时
			c.SetReadDeadline(time.Now().Add(30 * time.Second)) //设置发送接收数据超时
			return c, nil
		}}
	client := &http.Client{Transport: tr}

	//只有连接成功后，才会写入头的读取字节流
	if httpStart == true {
		//模拟一个host
		sHost := ""
		u, err := c_url.Parse(url)
		if err == nil {
			sHost = u.Host
		}

		//接收的格式
		req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")
		//连接使用后关闭
		req.Header.Set("Connection", "close")
		//设置host
		req.Header.Set("Host", sHost)
		//前一个网页
		req.Header.Set("Referer", sHost)
		//设置字符集和文档类型
		req.Header.Set("Content-Type", fmt.Sprintf("%s;charset=%s;", contentType, charset))
		req.Header.Set("Cache-Control", "no-cache, no-store, max-age=0, must-revalidate")
		//设置浏览器
		req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/67.0.3396.99 Safari/537.36")
		//
		//req.Header.Set("Accept-Encoding", "gzip")
		req.Header.Set("Accept-Encoding", "")

		//将自定义的head填入
		if len(headSet) > 0 {
			for k, v := range headSet {
				req.Header.Add(k, v)
			}
		}

		resp, err := client.Do(req)
		if err == nil {
			defer resp.Body.Close()
			statusCode = resp.StatusCode

			/*var reader io.ReadCloser
			if resp.Header.Get("Content-Encoding") == "gzip" {
				reader, err = gzip.NewReader(resp.Body)
				if err != nil {
					src = err.Error()
					statusCode = 503
				}
			} else {
				reader = resp.Body
			}*/

			body, err := ioutil.ReadAll(resp.Body)

			if err == nil {
				src = string(body)
			} else {
				LogsWithFileName("", "http_error", url+err.Error())
				src = err.Error()
				statusCode = 503
			}
		} else {
			LogsWithFileName("", "http_error", url+err.Error())
			src = err.Error()
			statusCode = 503
		}
	}
	return src, statusCode
}

/**
* 模拟http请求
* @url		string  需要抓取的url
* @method	string	请求方式，POST，GET，PUT等
* @body		string	需要传递的值
* @charset	string	字符编码
* @contentType string 定义http请求的文档格式，默认string
* @headSet	map[string] string	请求需要带上的头
 */
func HttpRequestWithRespExactHeader(url, method, body, charset, contentType, exactHeader string, headSet map[string]string) (string, int, string) {
	src := ""
	httpStart := false

	statusCode := 101
	respHeaderExact := ""

	if len(charset) < 1 {
		charset = "utf-8"
	}
	if len(method) < 1 {
		method = "POST"
	}
	sContentType := "text/html"
	if method == "POST" {
		sContentType = "application/x-www-form-urlencoded"
	}
	if len(contentType) < 1 {
		contentType = sContentType
	}

	var req *http.Request
	var err error
	var ioread io.Reader

	if len(body) > 0 {
		ioread = strings.NewReader(body)
	}

	//连续10次尝试
	for i := 0; i < 10; i++ {
		req, err = http.NewRequest(method, url, ioread)
		if err == nil {
			//如果成功了，跳出
			httpStart = true
			break
		}
	}
	//如果10次尝试后，都不能连接
	if err != nil {
		LogsWithFileName("", "http_error", url+err.Error())
		return err.Error(), 503, ""
	}

	//使用完连接后关闭
	if req.Body != nil {
		defer req.Body.Close()
	}

	//执行连接
	tr := &http.Transport{DisableKeepAlives: true,
		Dial: func(netw, addr string) (net.Conn, error) {
			c, err := net.DialTimeout(netw, addr, time.Second*60) //设置建立连接超时
			if err != nil {
				return nil, err
			}
			c.SetDeadline(time.Now().Add(30 * time.Second))     //设置发送接收数据超时
			c.SetReadDeadline(time.Now().Add(30 * time.Second)) //设置发送接收数据超时
			return c, nil
		}}
	client := &http.Client{Transport: tr}

	//只有连接成功后，才会写入头的读取字节流
	if httpStart == true {
		//模拟一个host
		sHost := ""
		u, err := c_url.Parse(url)
		if err == nil {
			sHost = u.Host
		}

		//接收的格式
		req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")
		//连接使用后关闭
		req.Header.Set("Connection", "close")
		//设置host
		req.Header.Set("Host", sHost)
		//前一个网页
		req.Header.Set("Referer", sHost)
		//设置字符集和文档类型
		req.Header.Set("Content-Type", fmt.Sprintf("%s;charset=%s;", contentType, charset))
		req.Header.Set("Cache-Control", "no-cache, no-store, max-age=0, must-revalidate")
		//设置浏览器
		req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/67.0.3396.99 Safari/537.36")
		//
		//req.Header.Set("Accept-Encoding", "gzip")
		req.Header.Set("Accept-Encoding", "")

		//将自定义的head填入
		if len(headSet) > 0 {
			for k, v := range headSet {
				req.Header.Add(k, v)
			}
		}

		resp, err := client.Do(req)
		if err == nil {
			defer resp.Body.Close()
			statusCode = resp.StatusCode
			respHeaderExact = resp.Header.Get(exactHeader)

			/*var reader io.ReadCloser
			if resp.Header.Get("Content-Encoding") == "gzip" {
				reader, err = gzip.NewReader(resp.Body)
				if err != nil {
					src = err.Error()
					statusCode = 503
				}
			} else {
				reader = resp.Body
			}*/

			body, err := ioutil.ReadAll(resp.Body)

			if err == nil {
				src = string(body)
			} else {
				LogsWithFileName("", "http_error", url+err.Error())
				src = err.Error()
				statusCode = 503
			}
		} else {
			LogsWithFileName("", "http_error", url+err.Error())
			src = err.Error()
			statusCode = 503
		}
	}
	return src, statusCode, respHeaderExact
}

/**
* 模拟http 表单请求
* @url		string  需要抓取的url
* @method	string	请求方式，POST，GET，PUT等
* @body		string	需要传递的值
* @charset	string	字符编码
* @contentType string 定义http请求的文档格式，默认string
* @headSet	map[string] string	请求需要带上的头
 */
func HttpFormRequestWithRespExactHeader(url string, method string, pams string, params map[string]string, charset, contentType, exactHeader string, headSet map[string]string) (string, int, string) {
	src := ""
	httpStart := false

	statusCode := 101
	respHeaderExact := ""

	if len(charset) < 1 {
		charset = "utf-8"
	}
	if len(method) < 1 {
		method = "POST"
	}
	sContentType := "text/html"
	if method == "POST" {
		sContentType = "application/x-www-form-urlencoded"
	}
	if len(contentType) < 1 {
		contentType = sContentType
	}

	var req *http.Request
	var err error

	values, _ := c_url.ParseQuery(pams)
	//连续10次尝试
	for i := 0; i < 10; i++ {
		req, err = http.NewRequest(method, url, strings.NewReader(values.Encode()))
		if err == nil {
			//如果成功了，跳出
			httpStart = true
			break
		}
	}
	//req.PostForm=values
	//err = req.ParseForm()
	//
	//fmt.Println(err)

	//如果10次尝试后，都不能连接
	if err != nil {
		LogsWithFileName("", "http_error", url+err.Error())
		return err.Error(), 503, ""
	}
	////如果10次尝试后，都不能连接
	if err != nil {
		LogsWithFileName("", "http_error", url+err.Error())
		return err.Error(), 503, ""
	}

	//使用完连接后关闭
	if req.Body != nil {
		defer req.Body.Close()
	}

	//执行连接
	tr := &http.Transport{DisableKeepAlives: true,
		Dial: func(netw, addr string) (net.Conn, error) {
			c, err := net.DialTimeout(netw, addr, time.Second*60) //设置建立连接超时
			if err != nil {
				return nil, err
			}
			c.SetDeadline(time.Now().Add(30 * time.Second))     //设置发送接收数据超时
			c.SetReadDeadline(time.Now().Add(30 * time.Second)) //设置发送接收数据超时
			return c, nil
		}}
	client := &http.Client{Transport: tr}

	//只有连接成功后，才会写入头的读取字节流
	if httpStart == true {
		//模拟一个host
		sHost := ""
		u, err := c_url.Parse(url)
		if err == nil {
			sHost = u.Host
		}

		//接收的格式
		req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")
		//连接使用后关闭
		req.Header.Set("Connection", "close")
		//设置host
		req.Header.Set("Host", sHost)
		//前一个网页
		req.Header.Set("Referer", sHost)
		//设置字符集和文档类型
		req.Header.Set("Content-Type", fmt.Sprintf("%s;charset=%s;", contentType, charset))
		req.Header.Set("Cache-Control", "no-cache, no-store, max-age=0, must-revalidate")
		//设置浏览器
		req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/67.0.3396.99 Safari/537.36")
		//
		//req.Header.Set("Accept-Encoding", "gzip")
		req.Header.Set("Accept-Encoding", "")

		//将自定义的head填入
		if len(headSet) > 0 {
			for k, v := range headSet {
				req.Header.Add(k, v)
			}
		}
		resp, err := client.Do(req)
		if err == nil {
			defer resp.Body.Close()
			statusCode = resp.StatusCode
			respHeaderExact = resp.Header.Get(exactHeader)

			/*var reader io.ReadCloser
			if resp.Header.Get("Content-Encoding") == "gzip" {
				reader, err = gzip.NewReader(resp.Body)
				if err != nil {
					src = err.Error()
					statusCode = 503
				}
			} else {
				reader = resp.Body
			}*/

			body, err := ioutil.ReadAll(resp.Body)

			if err == nil {
				src = string(body)
			} else {
				LogsWithFileName("", "http_error", url+err.Error())
				src = err.Error()
				statusCode = 503
			}
		} else {
			LogsWithFileName("", "http_error", url+err.Error())
			src = err.Error()
			statusCode = 503
		}
	}
	return src, statusCode, respHeaderExact
}

/**
* 模拟https的POST请求
* @url		string	需要发送请求的url
* @method	string	请求方式，POST，GET，PUT等
* @body		string	需要传递的值
* @charset	string	字符编码
* @contentType string 定义http请求的文档格式，默认string
* @sslCert	[]byte	证书主体
* @sslKey	[]byte	证书密钥
* @headSet	map[string] string	请求需要带上的头
 */
func HttpsRequest(url, method, body, charset, contentType string, sslCert []byte, sslKey []byte, headSet map[string]string) (string, int) {
	src := ""
	statusCode := 101
	httpStart := true

	_tlsConfig := &tls.Config{
		InsecureSkipVerify: true,
	}
	if len(sslCert) > 1 {
		//加载安全证书
		cert, err := tls.X509KeyPair(sslCert, sslKey)

		if err != nil {
			return "证书加载失败", statusCode
		}

		pool := x509.NewCertPool()
		pool.AppendCertsFromPEM(sslCert)
		_tlsConfig = &tls.Config{
			Certificates:       []tls.Certificate{cert},
			InsecureSkipVerify: true,
		}
	}

	if len(charset) < 1 {
		charset = "utf-8"
	}
	if len(method) < 1 {
		method = "POST"
	}
	sContentType := "text/html"
	if method == "POST" {
		sContentType = "application/x-www-form-urlencoded"
	}
	if len(contentType) < 1 {
		contentType = sContentType
	}

	var req *http.Request
	var err error
	var ioread io.Reader

	if len(body) > 0 {
		ioread = strings.NewReader(body)
	}

	//连续10次尝试
	for i := 0; i < 10; i++ {
		req, err = http.NewRequest(method, url, ioread)
		if err == nil {
			//如果成功了，跳出
			httpStart = true
			break
		}
	}
	//如果10次尝试后，都不能连接
	if err != nil {
		LogsWithFileName("", "http_error", url+"\n"+err.Error())
		return err.Error(), 503
	}

	//使用完连接后关闭
	if req.Body != nil {
		defer req.Body.Close()
	}

	tr := &http.Transport{TLSClientConfig: _tlsConfig,
		DisableKeepAlives: true,
		Dial: func(netw, addr string) (net.Conn, error) {
			c, err := net.DialTimeout(netw, addr, time.Second*30) //设置建立连接超时
			if err != nil {
				return nil, err
			}
			c.SetDeadline(time.Now().Add(5 * time.Second)) //设置发送接收数据超时
			return c, nil
		}}
	client := &http.Client{Transport: tr}

	//只有连接成功后，才会写入头的读取字节流
	if httpStart == true {
		//模拟一个host
		sHost := ""
		u, err := c_url.Parse(url)
		if err == nil {
			sHost = u.Host
		}
		//接收的格式
		req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")
		//连接使用后关闭
		req.Header.Set("Connection", "close")
		//设置host
		req.Header.Set("Host", sHost)
		//前一个网页
		req.Header.Set("Referer", sHost)
		//设置字符集和文档类型
		req.Header.Set("Content-Type", fmt.Sprintf("%s;charset=%s;", contentType, charset))
		//设置浏览器
		req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/42.0.2311.135 Safari/537.36")
		for k, v := range headSet {
			req.Header.Add(k, v)
		}
		resp, err := client.Do(req)
		if err == nil {

			defer resp.Body.Close()
			statusCode = resp.StatusCode

			/*var reader io.ReadCloser
			if resp.Header.Get("Content-Encoding") == "gzip" {
				reader, err = gzip.NewReader(resp.Body)
				if err != nil {
					src = err.Error()
					statusCode = 503
				}
			} else {
				reader = resp.Body
			}*/

			contents, err := ioutil.ReadAll(resp.Body)
			if err == nil {
				src = string(contents)
			} else {
				LogsWithFileName("", "http_error", url+"\n"+err.Error())
				src = err.Error()
				statusCode = 503
			}
		} else {
			LogsWithFileName("", "http_error", url+"\n"+err.Error())
			src = err.Error()
			statusCode = 503
		}
	}
	return src, statusCode
}

/**
* 模拟https的POST请求
* @url		string	需要发送请求的url
* @method	string	请求方式，POST，GET，PUT等
* @body		string	需要传递的值
* @charset	string	字符编码
* @contentType string 定义http请求的文档格式，默认string
* @sslCert	[]byte	证书主体
* @sslKey	[]byte	证书密钥
* @headSet	map[string] string	请求需要带上的头
 */
func HttpsRequestWithRespExactHeader(url, method, body, charset, contentType, exactHeader string, sslCert []byte, sslKey []byte, headSet map[string]string) (string, int, string) {
	src := ""
	statusCode := 101
	httpStart := true
	respHeaderExact := ""

	_tlsConfig := &tls.Config{
		InsecureSkipVerify: true,
	}
	if len(sslCert) > 1 {
		//加载安全证书
		cert, err := tls.X509KeyPair(sslCert, sslKey)

		if err != nil {
			return "证书加载失败", statusCode, ""
		}

		pool := x509.NewCertPool()
		pool.AppendCertsFromPEM(sslCert)
		_tlsConfig = &tls.Config{
			Certificates:       []tls.Certificate{cert},
			InsecureSkipVerify: true,
		}
	}

	if len(charset) < 1 {
		charset = "utf-8"
	}
	if len(method) < 1 {
		method = "POST"
	}
	sContentType := "text/html"
	if method == "POST" {
		sContentType = "application/x-www-form-urlencoded"
	}
	if len(contentType) < 1 {
		contentType = sContentType
	}

	var req *http.Request
	var err error
	var ioread io.Reader

	if len(body) > 0 {
		ioread = strings.NewReader(body)
	}

	//连续10次尝试
	for i := 0; i < 10; i++ {
		req, err = http.NewRequest(method, url, ioread)
		if err == nil {
			//如果成功了，跳出
			httpStart = true
			break
		}
	}
	//如果10次尝试后，都不能连接
	if err != nil {
		LogsWithFileName("", "http_error", url+"\n"+err.Error())
		return err.Error(), 503, respHeaderExact
	}

	//使用完连接后关闭
	if req.Body != nil {
		defer req.Body.Close()
	}

	tr := &http.Transport{TLSClientConfig: _tlsConfig,
		DisableKeepAlives: true,
		Dial: func(netw, addr string) (net.Conn, error) {
			c, err := net.DialTimeout(netw, addr, time.Second*30) //设置建立连接超时
			if err != nil {
				return nil, err
			}
			c.SetDeadline(time.Now().Add(5 * time.Second)) //设置发送接收数据超时
			return c, nil
		}}
	client := &http.Client{Transport: tr}

	//只有连接成功后，才会写入头的读取字节流
	if httpStart == true {
		//模拟一个host
		sHost := ""
		u, err := c_url.Parse(url)
		if err == nil {
			sHost = u.Host
		}
		//接收的格式
		req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")
		//连接使用后关闭
		req.Header.Set("Connection", "close")
		//设置host
		req.Header.Set("Host", sHost)
		//前一个网页
		req.Header.Set("Referer", sHost)
		//设置字符集和文档类型
		req.Header.Set("Content-Type", fmt.Sprintf("%s;charset=%s;", contentType, charset))
		//设置浏览器
		req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/42.0.2311.135 Safari/537.36")
		for k, v := range headSet {
			req.Header.Add(k, v)
		}
		resp, err := client.Do(req)
		if err == nil {

			defer resp.Body.Close()
			statusCode = resp.StatusCode
			respHeaderExact = resp.Header.Get(exactHeader)
			/*var reader io.ReadCloser
			if resp.Header.Get("Content-Encoding") == "gzip" {
				reader, err = gzip.NewReader(resp.Body)
				if err != nil {
					src = err.Error()
					statusCode = 503
				}
			} else {
				reader = resp.Body
			}*/

			contents, err := ioutil.ReadAll(resp.Body)
			if err == nil {
				src = string(contents)
			} else {
				LogsWithFileName("", "http_error", url+"\n"+err.Error())
				src = err.Error()
				statusCode = 503
			}
		} else {
			LogsWithFileName("", "http_error", url+"\n"+err.Error())
			src = err.Error()
			statusCode = 503
		}
	}
	return src, statusCode, respHeaderExact
}

/**
* 模拟http请求，携带cookie
* @url		string  需要抓取的url
* @method	string	请求方式，POST，GET，PUT等
* @body		string	需要传递的值
* @charset	string	字符编码
* @contentType string 定义http请求的文档格式，默认string
* @headSet	map[string] string	请求需要带上的头
 */
func HttpRequestByCookie(url, method, body, charset, contentType string, headSet map[string]string, cook []*http.Cookie) (string, int, []*http.Cookie) {
	src := ""
	httpStart := false

	statusCode := 101

	if len(charset) < 1 {
		charset = "utf-8"
	}
	if len(method) < 1 {
		method = "POST"
	}
	sContentType := "text/html"
	if method == "POST" {
		sContentType = "application/x-www-form-urlencoded"
	}
	if len(contentType) < 1 {
		contentType = sContentType
	}

	var req *http.Request
	var err error
	var ioread io.Reader

	if len(body) > 0 {
		ioread = strings.NewReader(body)
	}

	cookies := []*http.Cookie{} //新建一个cookie对象

	//连续10次尝试
	for i := 0; i < 10; i++ {
		req, err = http.NewRequest(method, url, ioread)
		if err == nil {
			//如果成功了，跳出
			httpStart = true
			break
		}
	}
	//如果10次尝试后，都不能连接
	if err != nil {
		LogsWithFileName("", "http_error", url+err.Error())
		return err.Error(), 503, cookies
	}

	//使用完连接后关闭
	if req.Body != nil {
		defer req.Body.Close()
	}

	//只有连接成功后，才会写入头的读取字节流
	if httpStart == true {
		//模拟一个host
		sHost := ""
		u, err := c_url.Parse(url)
		if err == nil {
			sHost = u.Host
		}

		//接收的格式
		req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")
		//连接使用后关闭
		req.Header.Set("Connection", "close")
		//设置host
		req.Header.Set("Host", sHost)
		//前一个网页
		req.Header.Set("Referer", sHost)
		//设置字符集和文档类型
		req.Header.Set("Content-Type", fmt.Sprintf("%s;charset=%s;", contentType, charset))
		req.Header.Set("Cache-Control", "no-cache, no-store, max-age=0, must-revalidate")
		//设置浏览器
		req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/42.0.2311.135 Safari/537.36")

		//将自定义的head填入
		if len(headSet) > 0 {
			for k, v := range headSet {
				req.Header.Add(k, v)
			}
		}

		//如果有cook的值，加载请求中
		if cook != nil {
			for _, value := range cook {
				req.AddCookie(value)
			}
		}

		//执行连接
		tr := &http.Transport{DisableKeepAlives: true,
			Dial: func(netw, addr string) (net.Conn, error) {
				c, err := net.DialTimeout(netw, addr, time.Second*30) //设置建立连接超时
				if err != nil {
					return nil, err
				}
				c.SetDeadline(time.Now().Add(5 * time.Second)) //设置发送接收数据超时
				return c, nil
			}}
		cookies = req.Cookies() //获取cookie结果
		client := &http.Client{Transport: tr}
		resp, err := client.Do(req)

		if err == nil {
			defer resp.Body.Close()
			statusCode = resp.StatusCode

			cookies = resp.Cookies() //失败后再获取cookie结果

			/*var reader io.ReadCloser
			if resp.Header.Get("Content-Encoding") == "gzip" {
				reader, err = gzip.NewReader(resp.Body)
				if err != nil {
					src = err.Error()
					statusCode = 503
				}
			} else {
				reader = resp.Body
			}*/

			body, err := ioutil.ReadAll(resp.Body)

			if err == nil {
				src = string(body)
			} else {
				LogsWithFileName("", "http_error", url+err.Error())
				src = err.Error()
				statusCode = 503
			}
		} else {
			LogsWithFileName("", "http_error", url+err.Error())
			src = err.Error()
			statusCode = 503
		}
	}
	return src, statusCode, cookies
}
