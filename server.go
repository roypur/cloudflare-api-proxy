package main

import("net/http"
       "strings"
       "fmt")

type ApiRequest struct{
    method string
    id string
    dnsName string
    dnsType string
    dnsContent string
    valid bool
}

func main(){
    http.HandleFunc("/", handle)

    listenStr := fmt.Sprintf("127.0.0.1:%d", port)

    err := http.ListenAndServe(listenStr, nil)
    if err != nil{
        fmt.Println(err)
    }
}

func handle(w http.ResponseWriter, req *http.Request){

    loggedin := false

    apiUser := req.Header.Get("Api-User")
    apiPass := req.Header.Get("Api-Pass")


    if (apiUser == userName) && (apiPass == userPass){
        loggedin = true
        w.Header().Set("Content-Type", "application/json; charset=utf-8")
        w.WriteHeader(http.StatusOK)
    }else{
        basicUser, basicPass, ok := req.BasicAuth()

        realmStr := fmt.Sprintf(`Basic realm="%s"`, httpRealm)

        w.Header().Set("WWW-Authenticate", realmStr)
        if ok{
            if (basicUser == userName) && (basicPass == userPass){
                loggedin = true
                w.Header().Set("Content-Type", "application/json; charset=utf-8")
                w.WriteHeader(http.StatusOK)
            }else{

                w.Header().Set("Content-Type", "text/plain; charset=utf-8")
                w.WriteHeader(http.StatusUnauthorized)
            }
        }else{
            w.Header().Set("Content-Type", "text/plain; charset=utf-8")
            w.WriteHeader(http.StatusUnauthorized)
        }
    }

    if loggedin{

        apiRequest := getParams(req)

        if apiRequest.valid{
            var data []byte
            var err error

            if apiRequest.method == "add"{
                data, err = dnsAdd(apiRequest.dnsName, apiRequest.dnsType, apiRequest.dnsContent)
            }else if apiRequest.method == "set"{
                data, err = dnsSet(apiRequest.id, apiRequest.dnsContent)
            }else if apiRequest.method == "get"{
                data, err = dnsGet(apiRequest.id)
            }else if apiRequest.method == "delete"{
                data, err = dnsDelete(apiRequest.id)
            }else if apiRequest.method == "list"{
                data, err = dnsList()
            }

            if err == nil{
                w.Write(data)
            }else{
                errStr := err.Error()
                w.Write([]byte(errStr))
            }
        }
    }else{
        w.Write([]byte("ACCESS DENIED"))
    }
}
func getParams(req *http.Request)(ApiRequest){

    rawQuery := strings.Split(req.URL.Path, "/")
    var paramCount = 0
    for k,v := range rawQuery{
        rawQuery[k] = strings.TrimSpace(v)
        if len(rawQuery[k]) != 0{
            paramCount++
        }
    }
    query := make([]string, paramCount)

    pos := 0
    for _,v := range rawQuery{
        if (len(v) != 0) && (pos < paramCount){
            query[pos] = v
            pos++
        }
    }

    var apiRequest ApiRequest
    apiRequest.valid = false

    if paramCount >= 1{
        apiRequest.method = strings.ToLower(query[0])
    }

    if paramCount == 1{
        if apiRequest.method == "list"{
            apiRequest.valid = true
        }
    }else if paramCount == 2{
        if (apiRequest.method == "delete") || (apiRequest.method == "get"){
            apiRequest.id = query[1]
            apiRequest.valid = true
        }
    }else if paramCount == 3{
        if apiRequest.method == "set"{
            apiRequest.id = query[1]
            apiRequest.dnsContent = query[2]
            apiRequest.valid = true
        }
    }else if paramCount == 4{
        if apiRequest.method == "add"{
            apiRequest.dnsName = strings.ToLower(query[1])
            apiRequest.dnsType = strings.ToUpper(query[2])
            apiRequest.dnsContent = query[3]
            apiRequest.valid = true
        }
    }
    return apiRequest
}
