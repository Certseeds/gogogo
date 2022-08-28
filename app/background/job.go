// SPDX-License-Identifier: AGPL-3.0-or-later
package main

import (
	"gogogo/pkg/background"
)

// main
// 预期:
//  + [x] 通过读取配置文件来获取关键配置
//  + [x] 持续运行,后台定期调用
//  + [x] 从某个地方获取资源
//  + [x] 进行某些演算
//  + [x] 最后存进某个地方
func main() {
	background.Main()
}
