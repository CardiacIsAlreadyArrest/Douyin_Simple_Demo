// Code generated by hertz generator.

package publish

import (
	"github.com/Yra-A/Douyin_Simple_Demo/cmd/api/biz/mw/jwt"
	"github.com/cloudwego/hertz/pkg/app"
)

func rootMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _douyinMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _publishMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _actionMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _publishactionMw() []app.HandlerFunc {
	// your code...
	return []app.HandlerFunc{
		jwt.JwtMiddleware.MiddlewareFunc(),
	}
}

func _listMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _publishlistMw() []app.HandlerFunc {
	// your code...
	return []app.HandlerFunc{
		jwt.JwtMiddleware.MiddlewareFunc(),
	}
}
