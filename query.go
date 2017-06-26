package main

import("fmt"
       "io/ioutil"
       "bytes"
       "encoding/json"
       "net/http")


func querySet(id string, content string)(CFRespSingle, error){
    resp, err := queryGet(id)
    var req CFreq = CFreq{}

    var jsonDecoded CFRespSingle

    req.DNStype = resp.Result.DNStype
    req.DNSname = resp.Result.DNSname
    req.DNScontent = content
    req.DNSttl = ttl
    req.Proxied = false

    data, err := json.Marshal(req)

    if err != nil{
        return jsonDecoded, err
    }

    httpResp, err := cfQuery("POST", fmt.Sprintf("/dns_records/%s", id), data)

    if err != nil{
        return jsonDecoded, err
    }
    body, err := ioutil.ReadAll(httpResp.Body)

    if err != nil{
        return jsonDecoded, err
    }

    err = json.Unmarshal(body, &jsonDecoded)

    return jsonDecoded, err

}

func queryGet(id string)(CFRespSingle, error){
    path := fmt.Sprintf("/dns_records/%s", id)

    var jsonDecoded CFRespSingle

    resp, err := cfQuery("GET", path, nil)
    if err != nil{
        return jsonDecoded, err
    }
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil{
        return jsonDecoded, err
    }

    err = json.Unmarshal(body, &jsonDecoded)
    return jsonDecoded, err
}

func queryAdd(dnsName string, dnsType string, dnsContent string)(CFRespSingle, error){
    var jsonDecoded CFRespSingle

    var req CFreq = CFreq{}

    req.DNSname = dnsName
    req.DNStype = dnsType
    req.DNScontent = dnsContent

    req.DNSttl = ttl
    req.Proxied = false

    data, err := json.Marshal(&req)

    if err != nil{
        return jsonDecoded, err
    }

    resp, err := cfQuery("POST", "/dns_records", data)

    if err != nil{
        return jsonDecoded, err
    }
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil{
        return jsonDecoded, err
    }

    err = json.Unmarshal(body, &jsonDecoded)
    return jsonDecoded, err
}

func queryDelete(id string)(CFRespSingle, error){
    var jsonDecoded CFRespSingle
    path := fmt.Sprintf("/dns_records/%s", id)

    resp, err := cfQuery("DELETE", path, nil)
    if err != nil{
        return jsonDecoded, err
    }
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil{
        return jsonDecoded, err
    }

    err = json.Unmarshal(body, &jsonDecoded)
    return jsonDecoded, err
}

func queryList(page int64)(CFRespList, error){
    path := fmt.Sprintf("/dns_records/?page=%d&per_page=%d", page, cfPerPage)

    var jsonDecoded CFRespList

    resp, err := cfQuery("GET", path, nil)
    if err != nil{
        return jsonDecoded, err
    }
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil{
        return jsonDecoded, err
    }

    err = json.Unmarshal(body, &jsonDecoded)
    return jsonDecoded, err
}

func cfQuery(method string, path string, body []byte)(*http.Response, error) {
    client := http.Client{}

    fmt.Println(string(body))

    pathBuf := bytes.NewBufferString(cfEndpoint)
    pathBuf.WriteRune('/')
    pathBuf.WriteString(cfZoneID)
    pathBuf.WriteString(path)

    url := pathBuf.String()

    fmt.Printf("-%s-\n", url)

    var err error
    var req *http.Request
    
    if len(body) > 0{
        buf := bytes.NewBuffer(body)
        fmt.Println("woot")
        req, err = http.NewRequest(method, url, buf)
    }else{
        req, err = http.NewRequest(method, url, nil)
    }

    if err != nil{
        return nil, err
    }

    req.Header.Set("X-Auth-Email", cfEmail)
    req.Header.Set("X-Auth-Key", cfApiKey)
    req.Header.Set("Content-Type", "application/json; charset=utf-8")

    return client.Do(req)
}
