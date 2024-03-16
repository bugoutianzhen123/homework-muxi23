package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"sync"
)

type SchoolId struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Data    struct {
		Item []struct {
			SchoolID int `json:"school_id"`
		} `json:"item"`
	} `json:"data"`
}

type SchoolInfo struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Data    struct {
		SchoolName string `json:"name"`
		Email      string `json:"email"`
		Site       string `json:"site"`
		CityName   string `json:"city_name"`
	} `json:"data"`
}

type SchoolData struct {
	SchoolName string `json:"name"`
	Email      string `json:"email"`
	Site       string `json:"site"`
	CityName   string `json:"city_name"`
}

var (
	//url  = "https://www.gaokao.cn/school/search"
	wg   sync.WaitGroup
	lock sync.Mutex
	num  int = 0 //成功数量
	all  []SchoolData
)

// 获取学校id
func getid(page int, client *http.Client) []struct {
	SchoolID int `json:"school_id"`
} {
	url := fmt.Sprintf("https://api.zjzw.cn/web/api/?keyword=&page=%d&province_id=&ranktype=&request_type=1&size=20&top_school_id=[2461,436]&type=&uri=apidata/api/gkv3/school/lists&signsafe=9326e2339790781062a5aac6ac933f66", page)
	//创建请求
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalf("Failed to create request: %v", err)
	}
	// 发送请求
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()
	// 读取响应内容
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Failed to read response body: %v", err)
	}

	// 解析 JSON 数据
	var data SchoolId
	err = json.Unmarshal(body, &data)
	if err != nil {
		log.Fatalf("Failed to unmarshal JSON: %v", err)
	}

	// 打印响应内容
	//fmt.Println("Response Body:", data.Data.Item)
	return data.Data.Item
}

// 获取学校信息
func getschoolinfo(id int, client *http.Client, ch1 chan int) {
	ch1 <- 1

	url := fmt.Sprintf("https://static-data.gaokao.cn/www/2.0/school/%d/info.json", id)
	//创建请求
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalf("Failed to create request: %v", err)
	}
	// 发送请求
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()
	// 读取响应内容
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Failed to read response body: %v", err)
	}

	// 解析 JSON 数据
	var data SchoolInfo
	err = json.Unmarshal(body, &data)
	if err != nil {
		log.Fatalf("Failed to unmarshal JSON: %v", err)
	}

	// 打印响应内容
	//fmt.Println("Response Body:", data)

	lock.Lock()
	//存入全局变量
	all = append(all, data.Data)
	//成功数加一
	num++
	fmt.Printf("已爬取数量：%d\n", num)
	lock.Unlock()
	wg.Done()
	<-ch1
}

// 存储
func into(data []SchoolData) {
	// 打开文件
	file, err := os.Create("school_data.json")
	if err != nil {
		log.Fatalf("Failed to create file: %v", err)
	}
	defer file.Close()

	// 将数据编码为 JSON 格式并写入文件
	encoder := json.NewEncoder(file)
	err = encoder.Encode(data)
	if err != nil {
		log.Fatalf("Failed to encode data: %v", err)
	}

	fmt.Println("Data has been written to school_data.json")
}

func main() {
	ch1 := make(chan int, 10)
	// 创建一个 HTTP 客户端
	client := &http.Client{}
	//a := getid(2, client)
	//fmt.Println(a)
	//getschoolinfo(140, client, ch1)
	//开始
	for i := 1; i < 11; i++ {
		schoolid := getid(i, client)
		for _, id := range schoolid {
			//fmt.Println(id.SchoolID)
			go getschoolinfo(id.SchoolID, client, ch1)
			wg.Add(1)
		}
	}
	//等待结束
	wg.Wait()
	//储存
	into(all)
}
