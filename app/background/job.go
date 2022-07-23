// SPDX-License-Identifier: AGPL-3.0-or-later
package main

import (
	"gogogo/pkg/background"
)

// main
// 预期:
//  + 通过读取配置文件来获取关键配置
//  + 持续运行,后台定期调用某个线程,从某个地方获取资源,进行某些演算,最后存进某个地方
func main() {
	background.Main()
}
