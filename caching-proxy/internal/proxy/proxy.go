package proxy

import (
	"caching-proxy/internal/cache"
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
	resp,_:=http.Get(targetURL)
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

func (p *ProxyServer) handleDirectForward(w http.ResponseWriter, r *http.Request){
	structuredURL, _:=url.Parse(p.OriginURL)
	proxy:=httputil.NewSingleHostReverseProxy(structuredURL)
	proxy.ServeHTTP(w,r)
}