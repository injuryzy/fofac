package fetch

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/injuryzy/fofac/fofac/excel"
	"github.com/injuryzy/fofac/fofac/log"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

var (
	errorUnauthorized = errors.New("401 Unauthorized, make sure email and apikey is correct.")
	errorForbidden    = errors.New("403 Forbidden, can't access the fofa service normally.")
	error503          = errors.New("503 Service Temporarily Unavailable")
	error502          = errors.New("502 Bad Gateway")
	fields            = "host,ip,port,server,domain,title,country,province,city,icp,protocol"
	FetchResult       = make(map[string]([][]string))
)

type FofaQuery struct {
	Page         int8
	Size         int
	Full         bool
	Key          string
	Email        string
	Query        string
	Before       string
	After        string
	TimeInterval int
	FileName     string
}

// QueryResp 接口请求相应
type QueryResp struct {
	Error           bool       `json:"error"`
	ConsumedFpoint  int        `json:"consumed_fpoint"`
	RequiredFpoints int        `json:"required_fpoints"`
	Size            int        `json:"size"`
	Page            int        `json:"page"`
	Mode            string     `json:"mode"`
	Query           string     `json:"query"`
	Results         [][]string `json:"results"`
	Errmsg          string     `json:"errmsg"`
}

type FofaSearch struct {
	client http.Client
	FofaQuery
}

// QueryResult 没有查询范围，
// 直接查询，返回结果输出
func (f *FofaSearch) QueryResult() {
	q := f.Query

	if f.Before != "" {
		q += "&& before=" + f.Before
	} else if f.After != "" {
		q += "&& after=" + f.After
	}
	base64Query := base64.StdEncoding.EncodeToString([]byte(q))
	if f.Size > 10000 {
		log.Error("查询参数: %s, 查询条件大于10000条，没有查询完毕,", q)
		f.Size = 10000
	}

	queryUrl := fmt.Sprintf("https://fofa.info/api/v1/search/all?email=%s&key=%s&qbase64=%s&size=%d&fields=%s", f.Email, f.Key, base64Query, f.Size, fields)
	resp, err := f.HttpGet(queryUrl)
	if err != nil {
		log.Error("查询参数: %s, err: %v", q, err)
		return
	}
	m := make(map[string][][]string)
	m[q] = resp.Results
	parse, _ := time.Parse(time.DateOnly, f.Before)
	end := parse.AddDate(0, 0, f.TimeInterval).Format(time.DateOnly)
	dir, err2 := os.Getwd()
	if err2 != nil {
		log.Error("获取当前目录失败,%v", err)
	}

	if f.FileName != "" {
		excel.WriteXlsx(m, fmt.Sprintf("%s%s%s.xlsx", dir, string(os.PathSeparator), f.FileName))
		return
	}
	//范围内的数据
	if f.Before != "" && f.After != "" {
		excel.WriteXlsx(m, fmt.Sprintf("%s%s%s-%s.xlsx", dir, string(os.PathSeparator), parse.Format(time.DateOnly), end))
	} else {
		if f.After != "" {
			//之后的数据
			excel.WriteXlsx(m, fmt.Sprintf("%s%s%s.xlsx", dir, string(os.PathSeparator), f.After))
		} else if f.Before != "" {
			//之前的数据
			excel.WriteXlsx(m, fmt.Sprintf("%s%s%s.xlsx", dir, string(os.PathSeparator), f.Before))
		} else {
			//没有时间范围
			excel.WriteXlsx(m, fmt.Sprintf("%s%s%s.xlsx", dir, string(os.PathSeparator), time.Now().Format(time.DateOnly)))
		}

	}

}

// QuerySize 查询条数

func (f *FofaSearch) QuerySize() int {
	q := f.Query

	if f.Before != "" && f.After != "" {
		q += fmt.Sprintf(" && before =%s", f.Before)
	}
	log.Info("查询参数：%s", q)
	base64Query := base64.StdEncoding.EncodeToString([]byte(q))

	queryUrl := fmt.Sprintf("https://fofa.info/api/v1/search/all?email=%s&key=%s&qbase64=%s&size=%d&fields=%s", f.Email, f.Key, base64Query, 1, fields)
	resp, err := f.HttpGet(queryUrl)
	if err != nil {
		log.Error("请求错误", err)
	}
	return resp.Size

}

// QuerySize 查询query 间隔的条数
func (f *FofaSearch) QueryT() {
	size := f.QuerySize()
	parse, _ := time.Parse(time.DateOnly, f.Before)
	f.Before = parse.AddDate(0, 0, -f.TimeInterval).Format(time.DateOnly)
	querySize := f.QuerySize()
	qSize := size - querySize
	if qSize == 0 {
		f.QueryT()
	}
	if qSize > 10000 {
		f.Size = 10000
		log.Info("总条数：%d", qSize)
	} else {
		f.Size = qSize
	}
	f.QueryResult()
	start, _ := time.Parse(time.DateOnly, f.Before)
	end, _ := time.Parse(time.DateOnly, f.After)
	if start.Unix() >= end.Unix() {
		f.QueryT()
	}
}

// HttpGet 将http响应封装
func (f *FofaSearch) HttpGet(u string) (*QueryResp, error) {
	body, err := f.client.Get(u)
	if err != nil {
		return nil, err
	}
	defer body.Body.Close()

	if body.StatusCode == 503 {
		return nil, error503
	} else if body.StatusCode == 502 {
		return nil, error502
	} else if body.StatusCode == 403 {
		return nil, errorForbidden
	}

	content, err := ioutil.ReadAll(body.Body)

	var resp *QueryResp

	json.Unmarshal(content, &resp)
	if resp.Error == true {
		log.Error("接口查询错误 :%s", resp.Errmsg)
	}
	return resp, nil
}
