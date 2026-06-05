package proxy

import (
	"caching-proxy/internal/cache"
	"fmt"
	"io"
	"net/http"
	"net/http/httputil"
	"net/url"
)

type ProxyServer struct{
	Cache *cache.CacheSystem
	OriginURL string
}

func (p *ProxyServer) ProxyHandler(w http.ResponseWriter, r *http.Request){
	
	cacheKey:=r.URL.Path

	if r.Method=="GET"{
		cachedResponse, found := p.Cache.Get(cacheKey)
		if found{
			for key, values:=range cachedResponse.Headers{
				for _,val:=range values{
					w.Header().Add(key,val)
				}
			}
			w.Header().Set("X-Cache", "HIT")
			w.WriteHeader(cachedResponse.StatusCode)
			w.Write(cachedResponse.Body)
			return
		}
		p.handleCacheMiss(w,r,cacheKey)
		return
	}

	p.handleDirectForward(w, r)
}

func (p *ProxyServer) handleCacheMiss(w http.ResponseWriter, r *http.Request, key string){
	targetURL:=p.OriginURL+r.URL.Path
	resp,err:=http.Get(targetURL)
	if err!=nil{
		fmt.Println("Network error: ",err)
		w.WriteHeader(http.StatusBadGateway)
		w.Write([]byte("Bad Gateway"))
		return
	}
	defer resp.Body.Close()
	
	body,_:=io.ReadAll(resp.Body)
	c:=cache.CachedResponse{
		Body: body,
		Headers: resp.Header,
		StatusCode: resp.StatusCode,
	}
	p.Cache.Set(key, c)
	w.Header().Set("X-Cache","MISS")
	w.WriteHeader(c.StatusCode)
	w.Write(body)
}

// func (p *ProxyServer) handleDirectForward(w http.ResponseWriter, r *http.Request){
// 	structuredURL, _:=url.Parse(p.OriginURL)
// 	proxy:=httputil.NewSingleHostReverseProxy(structuredURL)

// 	proxy.ModifyResponse=func(resp *http.Response)error{

// 		if (resp.StatusCode==200 || resp.StatusCode==201) && 
// 		(resp.Request.Method=="PUT" || resp.Request.Method=="POST" || 
// 		resp.Request.Method=="DELETE"){
// 			cache.NewCacheSystem().Clear()	
// 			fmt.Println("Mutation deteccted, cache invalidated")
// 		}

// 		proxy.ServeHTTP(w,r)
// 		return nil
// 	}

// }

func (p *ProxyServer) handleDirectForward(w http.ResponseWriter, r *http.Request) {
	structuredURL, err := url.Parse(p.OriginURL)
	if err != nil {
		fmt.Println("Invalid origin URL configured:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	proxy := httputil.NewSingleHostReverseProxy(structuredURL)

	proxy.ModifyResponse = func(resp *http.Response) error {
		isSuccess := resp.StatusCode == http.StatusOK || resp.StatusCode == http.StatusCreated
		isMutation := r.Method == "POST" || r.Method == "PUT" || r.Method == "DELETE"

		if isSuccess && isMutation {
			p.Cache.Clear()
			fmt.Println("Mutation detected, active cache invalidated safely.")
		}

		return nil
	}

	proxy.ServeHTTP(w, r)
}