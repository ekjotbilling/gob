package go

import (
	// "encoding/json"
	// "fmt"
	// "io"
	// "io/ioutil"
	// "net/http"
)

// var (
// 	httpRequestType = reflect.TypeOf((*Request)(nil)).Elem()
// )

// 	httpResponseWriterType = reflect.TypeOf((*http.ResponseWriter)(nil)).Elem()
// )


// func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
//   // request := &Request{req}
//   //
// 	// cleanedPath := path.Clean(req.URL.Path)
// 	// method := req.Method
//   //
// 	// res := r.rootRouter.tree.Find(cleanedPath)
// 	// if res == nil {
// 	// 	fmt.Println("404 not found!")
// 	// 	w.WriteHeader(404)
// 	// 	return
// 	// }
//   //
// 	// var reqRoute *route
//   //
// 	// switch res.Value.(type) {
// 	// case *Router:
// 	// 	router := res.Value.(*Router)
// 	// 	if router.rootRoute == nil {
// 	// 		fmt.Println("router 404!")
// 	// 		w.WriteHeader(404)
// 	// 		return
// 	// 	}
// 	// 	reqRoute = router.rootRoute
// 	// case *route:
// 	// 	reqRoute = res.Value.(*route)
// 	// }
//   //
// 	// if reqRoute.Method != method {
// 	// 	fmt.Println("method 404!")
// 	// 	w.WriteHeader(404)
// 	// 	return
// 	// }
//   //
// 	// allRouters := []*Router{}
// 	// currRouter := reqRoute.Router
// 	// for currRouter != nil {
// 	// 	allRouters = append(allRouters, currRouter)
// 	// 	currRouter = currRouter.parent
// 	// }
//   //
// 	// requestContext := reflect.New(allRouters[len(allRouters)-1].contextType)
// 	// for _, fn := range allRouters[len(allRouters)-1].middlewares {
// 	// 	fn.Call([]reflect.Value{
// 	// 		requestContext,
// 	// 		reflect.ValueOf(w),
// 	// 		reflect.ValueOf(req),
// 	// 	})
// 	// }
//   //
// 	// // go down router chain, generate contexts and do the linking
// 	// for i := len(allRouters) - 2; i >= 0; i-- {
// 	// 	router := allRouters[i]
// 	// 	ctx := reflect.New(router.contextType)
//   //
// 	// 	f := reflect.Indirect(ctx).Field(0)
// 	// 	f.Set(requestContext)
//   //
// 	// 	// run middlewares
// 	// 	for _, fn := range router.middlewares {
// 	// 		fn.Call([]reflect.Value{
// 	// 			ctx,
// 	// 			reflect.ValueOf(w),
// 	// 			reflect.ValueOf(req),
// 	// 		})
// 	// 	}
//   //
// 	// 	requestContext = ctx
// 	// }
//   //
// 	// handlerType := reqRoute.Handler.Type()
// 	// numParams := handlerType.NumIn()
//   //
// 	// // FIXME: maybe move this logic to when adding routes ???
// 	// if numParams <= 2 {
// 	// 	var res []reflect.Value
// 	// 	if numParams == 2 {
// 	// 		body, err := ioutil.ReadAll(req.Body)
// 	// 		if err != nil {
// 	// 			panic(err)
// 	// 		}
//   //
// 	// 		reqObj := reflect.New(handlerType.In(1).Elem()).Interface()
// 	// 		if err := json.Unmarshal(body, &reqObj); err != nil {
// 	// 			panic(err)
// 	// 		}
//   //
// 	// 		res = reqRoute.Handler.Call([]reflect.Value{
// 	// 			requestContext,
// 	// 			reflect.ValueOf(reqObj),
// 	// 		})
// 	// 	} else {
// 	// 		res = reqRoute.Handler.Call([]reflect.Value{requestContext})
// 	// 	}
//   //
// 	// 	rsp, rspCode, rerr := res[0].Interface(), int(res[1].Int()), res[2].Interface()
// 	// 	if rerr != nil {
// 	// 		panic(rerr)
// 	// 	}
//   //
// 	// 	data, err := json.Marshal(rsp)
// 	// 	if err != nil {
// 	// 		panic(err)
// 	// 	}
// 	// 	w.WriteHeader(rspCode)
// 	// 	io.WriteString(w, string(data))
// 	// } else if numParams == 3 {
// 	// 	if handlerType.In(1) == httpResponseWriterType && handlerType.In(2).Elem() == httpRequestType {
// 	// 		reqRoute.Handler.Call([]reflect.Value{
// 	// 			requestContext,
// 	// 			reflect.ValueOf(w),
// 	// 			reflect.ValueOf(req),
// 	// 		})
// 	// 	}
// 	// }
// }