package main

import (
	"testing"
)

func Test_GMOF_Cai_Pai_Month(t *testing.T) {
	url := "http://zhs.mof.gov.cn/zhuantilanmu/caipiaoguanli/201712/t20171221_2786493.html"

	caipiaoSpider := init_GMOF_CaiPiao_Month_HTMLSpider(url)
	ctx, _ := caipiaoSpider.Setup(nil)
	caipiaoSpider.Spin(ctx)
}
