package network

import (
	"net/http"
	"fmt"
	"io/ioutil"
)
/*
//0故障  -1无效文件  1成功
func LoadRemoteFile(url string) ([]byte,int) {
	resp, err := http.DefaultClient.Get(url)
	if err != nil {
		fmt.Println(err)
		return nil,0
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound || resp.StatusCode == http.StatusForbidden {
		return  nil,-1
	}

	if resp.StatusCode != http.StatusOK {
		fmt.Println(err, resp.StatusCode, url)
		return  nil,0
	}

	clen := resp.Header.Get("Content-Length")




	data,err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return  nil,0
	}

	if clen != "" {
		tlen, _ := strconv.Atoi(clen)
		if len(data) != tlen {
			fmt.Println(len(data), tlen)
			return  nil,0
		}
	}
	return data,1
}
*/
func MustGet(url string) ([]byte,error) {
	for ; ;  {
		resp, err := http.DefaultClient.Get(url)
		if err != nil {
			continue
		}
		defer resp.Body.Close()

		if resp.StatusCode == http.StatusNotFound || resp.StatusCode == http.StatusForbidden{
			fmt.Println("404:",url)
			return nil,fmt.Errorf("404")
		}

		if resp.StatusCode != http.StatusOK  {
			continue
		}

		bytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			continue
		}
		return bytes,err
	}
}