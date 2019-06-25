package support

import . "github.com/ahmetb/go-linq"

/**
* @author zhangsheng
* @date 2019/6/20
 */

func GroupBy(list interface{}, key func(interface{}) interface{}, item func(interface{}) interface{}) map[string][]interface{} {
	group := From(list).GroupBy(key, item)
	result := make(map[string][]interface{})
	group.ForEach(func(it interface{}) {
		g := it.(Group)
		result[g.Key.(string)] = g.Group
	})
	return result
}
