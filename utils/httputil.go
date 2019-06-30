/*
@Time : 2019/5/14 9:30 
@Author : Lukebryan
@File : httputil.go
@Software: GoLand
*/
package utils

import (

)	//"bytes"
//"crypto/tls"
//"errors"
//"github.com/spf13/cast"
//"io/ioutil"
//"log"
//"mime/multipart"
//"net/http"

/*
form表单post请求
*/
//func PostFormRequest(path string,datas map[string]interface{},headers map[string]string ) (resultBody string,err error) {
//
//	body := new(bytes.Buffer)
//	w := multipart.NewWriter(body)
//
//	for k,v := range datas{
//		w.WriteField(k,cast.ToString(v))
//	}
//	w.Close()
//
//	//跳过证书验证
//	tr := &http.Transport{
//		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
//	}
//	client := http.Client{Transport:tr}
//	req, err := http.NewRequest("POST",path,body)
//	if err != nil {
//		log.Println(err)
//		return "",err
//	}
//
//	req.Header.Add("Content-Type", w.FormDataContentType())
//	for k,v := range headers{
//		req.Header.Add(k, v)
//	}
//
//	//req.Header.Add("Connection-Type", "Keep-Alive")
//	//req.Header.Add("Accept", "*/*")
//	//req.Header.Add("Authorization", conf.Conf.Authorization)
//
//	//处理返回结果
//	response, err := client.Do(req)
//	if err != nil {
//		log.Print(err)
//		return "",err
//	}
//	defer response.Body.Close()
//
//	if err != nil {
//		log.Print(err)
//		return "",err
//	}
//
//	resultStr, err := ioutil.ReadAll(response.Body)
//	if err != nil {
//		log.Println(err)
//		return "",err
//	}
//
//	if response.Status != "200 OK"{
//		return "",errors.New(string(resultStr))
//	}
//
//	return string(resultStr),nil
//}

/*
原样请求
 */
//func PostRequest(r *http.Request,path string,userID,wxID string) (resultBody string,err error) {
//	body := r.Body
//	//跳过证书验证
//	tr := &http.Transport{
//		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
//	}
//	client := http.Client{Transport:tr}
//	req, err := http.NewRequest("POST",path,body)
//	if err != nil {
//		log.Println(err)
//		return "",err
//	}
//	r.Header.Add("Authorization",GentToken(userID,wxID))
//	req.Header.Add("Content-Type", r.w.FormDataContentType())
//
//	//处理返回结果
//	response, err := client.Do(req)
//	defer response.Body.Close()
//
//	if err != nil {
//		log.Print(err)
//		return "",err
//	}
//
//	resultStr, err := ioutil.ReadAll(response.Body)
//	if err != nil {
//		log.Println(err)
//		return "",err
//	}
//
//	if response.Status != "200 OK"{
//		return "",errors.New(string(resultStr))
//	}
//
//	return string(resultStr),nil
//}