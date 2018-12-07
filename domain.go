package common

import (
	"strings"
)

/**
* 传入一个域名，获取该域名的顶级域名
* @param		domain		域名
* @return  string	标准的host
 */
func GetTopDomain(domain string) string {
	resDomain := domain

	//不是本机才进行下面操作
	if len(domain) > 0 && strings.Count(domain, "127.0.0.1") < 1 {
		//ip的正则表达式
		ipMatch := `^(25[0-5]|2[0-4][0-9]|[0-1]{1}[0-9]{2}|[1-9]{1}[0-9]{1}|[1-9])\.(25[0-5]|2[0-4][0-9]|[0-1]{1}[0-9]{2}|[1-9]{1}[0-9]{1}|[1-9]|0)\.(25[0-5]|2[0-4][0-9]|[0-1]{1}[0-9]{2}|[1-9]{1}[0-9]{1}|[1-9]|0)\.(25[0-5]|2[0-4][0-9]|[0-1]{1}[0-9]{2}|[1-9]{1}[0-9]{1}|[0-9])$`
		ipCheck := RegMatch(ipMatch, domain)
		if len(ipCheck) == 1 {
			//如果只是ip地址，顶级域名 = ip
			resDomain = domain
		} else {
			//按逗号切割
			tmpSlicp := strings.Split(domain, ".")
			//超过2段，表示是二级域名
			if len(tmpSlicp) > 2 {
				n := len(tmpSlicp)
				resDomain = tmpSlicp[n-2] + "." + tmpSlicp[n-1]
			}
		}
	}
	return resDomain
}
