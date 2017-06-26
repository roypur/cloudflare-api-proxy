package main

import("encoding/json")

type CFreq struct{
    DNStype string `json:"type,omitempty"`
    DNSname string `json:"name,omitempty"`
    DNScontent string `json:"content,omitempty"`
    DNSttl int64 `json:"ttl"`
    Proxied bool `json:"proxied"`
}

type CFRespList struct{
    Result []APIresult `json:"result,omitempty"`
    Info CFinfo `json:"result_info,omitempty"`
}

type CFRespSingle struct{
    Result APIresult `json:"result,omitempty"`
    Success bool `json:"success"`
}

type APIresult struct{
    ID string `json:"id,omitempty"`
    DNStype string `json:"type,omitempty"`
    DNSname string `json:"name,omitempty"`
    DNScontent string `json:"content,omitempty"`
    Modified string `json:"modified_on,omitempty"`
    Created string `json:"created_on,omitempty"`
    DNSttl int64 `json:"ttl"`
}

type CFinfo struct{
    Page int64 `json:"page"`
    PerPage int64 `json:"per_page"`
    TotalPages int64 `json:"total_pages"`
    Count int64 `json:"count"`
    TotalCount int64 `json:"total_count"`
}

func dnsList()([]byte, error){
    resp, err := queryList(1)

    if err != nil{
        return nil, err
    }

    recordCount := resp.Info.TotalCount
    pageCount := resp.Info.TotalPages

    var entries []APIresult = make([]APIresult, recordCount)

    var pos int64 = 0

    var i int64 = 0
    for ;i<pageCount; i++{

        for _,v := range resp.Result{
            if pos < recordCount{
                entries[pos] = v
            }
            pos++
        }
        if i < (pageCount - 1){
            resp, err = queryList(i+2)
        }
    }

    return json.MarshalIndent(&entries, "", "    ")
}

func dnsGet(id string)([]byte, error){
    resp, err := queryGet(id)
    if err != nil{
        return nil, err
    }

    return json.MarshalIndent(&resp, "", "    ")
}

func dnsAdd(dnsName, dnsType, dnsContent string)([]byte, error){
    resp, err := queryAdd(dnsName, dnsType, dnsContent)
    if err != nil{
        return nil, err
    }

    return json.MarshalIndent(&resp, "", "    ")
}

func dnsSet(id, content string)([]byte, error){
    resp, err := querySet(id, content)
    if err != nil{
        return nil, err
    }

    return json.MarshalIndent(&resp, "", "    ")
}

func dnsDelete(id string)([]byte, error){
    resp, err := queryDelete(id)
    if err != nil{
        return nil, err
    }

    return json.MarshalIndent(&resp, "", "    ")
}
