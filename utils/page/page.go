package page

import (
	"github.com/astaxie/beego/orm"
	"strconv"
	"math"
)

const (
	DEFAULT_LINK_ITEM_COUNT = 5
	DEFAULT_FIRST_PAGE_TEXT = "首页"
	DEFAULT_LAST_PAGE_TEXT  = "尾页"
	DEFAULT_PRE_PAGE_TEXT   = "上一页"
	DEFAULT_NEXT_PAGE_TEXT  = "下一页"
)

type PageLinks struct {
	Query               orm.QuerySeter //查询条件
	Currentpage         int    //当前页 ,默认1 每次分页,必须在前台设置新的页数,不设置始终默认1.在控制器中使用方式:cp, _ := this.GetInt("pno")   po.Currentpage = int(cp)
	PageDataSize        int    //每页显示多少条数据
	LinkItemCount       int    //生成A标签的个数 默认10个
	Href                string //A标签的链接地址
	FirstPageText       string //首页文字  默认"首页"
	LastPageText        string //尾页文字  默认"尾页"
	PrePageText         string //上一页文字 默认"上一页"
	NextPageText        string //下一页文字 默认"下一页"
	total               int    //总数
	start               int    //当前页面开始显示数
	end                 int    //当前页面结束显示数
	countPage           int    //可以分多少页
	startData           string //从 XXX 记录
	endData             string //到 XXX 记录
	Offset              int    //偏移量
}

func (p *PageLinks) GetLinks() string {
	p.init()
	p.total = p.getTotal()
	p.getShowPageNumber()
	return p.linkStyleOne()
}

//获取偏移量
func (p *PageLinks) GetOffset() int {
	p.getTotal()
	//偏移量
	p.Offset = p.Currentpage * p.PageDataSize - p.PageDataSize
	//不能超出可取范围
	if p.Offset >= p.total {
		p.Offset = p.total - p.PageDataSize
	}
	return p.Offset
}

//获取总数
func (p *PageLinks) getTotal() int {
	int64, _ := p.Query.Count()
	strNum :=strconv.FormatInt(int64,10)
	total, _:=strconv.Atoi(strNum)
	p.total = total
	return total
}

//分页样式1
func (p *PageLinks) linkStyleOne() string {
	//没有总数
	if p.total == 0 || (p.start == 0 && p.end == 0) {
		return ``
	}

	strHtmlHead := `
					<div class="row">
                        <div class="col-sm-5">
                            <div class="dataTables_info">从 `+ p.startData +` 到  `+ p.endData+` 条记录 总记录数为 `+ strconv.Itoa(p.total) +` 条</div></div>
                        <div class="col-sm-7">
                            <div class="dataTables_paginate paging_full_numbers">
                                <ul class="pagination">

	`

	//显示首页按钮
	if p.Currentpage != 1 {
		if p.Currentpage >= p.countPage {
			p.Currentpage = p.countPage
		}

		prePageNum := strconv.Itoa(p.Currentpage - 1)
		strHtmlHead += `
									<li class="paginate_button first" >
										<a href="`+ p.Href +`">`+ p.FirstPageText +`</a>
									</li>
									<li class="paginate_button previous">
										<a href="`+ p.Href + "/"+ prePageNum +`">`+ p.PrePageText +`</a>
									</li>
		`
	}

	strHtmlCenter := ``

	for i:=p.start; i <= p.end; i++ {
		var active string
		var disabled string
		var href string
		href = p.Href + "/" + strconv.Itoa(i)
		if i == p.Currentpage {
			active = "active"
			disabled = "disabled"
			href = "javascript:void(0)"
		}
		//超出页码数
		if i == p.end && p.Currentpage > p.end {
			active = "active"
			disabled = "disabled"
			href = "javascript:void(0)"
		}
		strHtmlCenter += `
									<li class="paginate_button `+ active +` `+ disabled +` ">
										 <a href="`+ href +`">`+ strconv.Itoa(i) +`</a>
									</li>
		`
	}

	strHtmlTail := ``

	//显示 尾页按钮
	if p.Currentpage < p.countPage {
		nextPageNum := strconv.Itoa(p.Currentpage + 1)
		strHtmlTail += `
									<li class="paginate_button next">
										<a href="`+ p.Href + "/" + nextPageNum +`">`+ p.NextPageText +`</a>
									</li>
									<li class="paginate_button last">
										<a href="`+ p.Href + "/" + strconv.Itoa(p.countPage) +`">`+ p.LastPageText +`</a>
									</li>
		`
	}

	strHtmlTail += `

								</ul>
                            </div>
                        </div>
                    </div>
	`

	return strHtmlHead + strHtmlCenter + strHtmlTail
}

//获取当前页需要显示的分页
func (p *PageLinks) getShowPageNumber() {
	//未填写每一页显示 数据量
	if p.PageDataSize == 0 {
		p.start = 0
		p.end   = 0
	}

	//总数等于0
	if p.total == 0 || p.total <= p.PageDataSize {
		p.start = 0
		p.end   = 0
		return
	}

	//int to string
	strTotal := strconv.Itoa(p.total)
	strPageDataSize := strconv.Itoa(p.PageDataSize)

	//string to float64
	float64Total, _ := strconv.ParseFloat(strTotal, 64)
	float64PageDataSize, _ := strconv.ParseFloat(strPageDataSize, 64)

	//数据总量 除以 每页最大分页数 = 有多少页
	float64CountPage := math.Ceil(float64Total / float64PageDataSize)
	p.countPage = int(float64CountPage)

	showPageNums := make([]int, 0)

	//期望 显示的最大分页数
	showMaxPageNum := p.Currentpage + p.LinkItemCount

	//起始页码数
	var index int = p.Currentpage

	//期望页码数 超出实际最大页码数 只返回实际最大页码数
	if showMaxPageNum > p.countPage {
		showMaxPageNum = p.countPage
		//最大码数 - LinkItemCount
		index = showMaxPageNum - p.LinkItemCount
	}

	//防止负页数
	if index <= 0 {
		index = 1
	}

	for i := index; i <= showMaxPageNum; i++  {
		showPageNums = append(showPageNums, i)
	}

	//非页尾 靠左选择 LinkItemCount 个数
	if p.Currentpage < p.countPage && len(showPageNums) > p.LinkItemCount {
		showPageNums = showPageNums[0:p.LinkItemCount]
	}

	//页尾 选择 靠右选择
	if p.Currentpage >= p.countPage && len(showPageNums) > p.LinkItemCount {
		showPageNums = showPageNums[len(showPageNums) - p.LinkItemCount:]
	}

	p.start = showPageNums[0]
	p.end   = showPageNums[len(showPageNums) - 1]

	p.startData = strconv.Itoa(p.Currentpage * p.PageDataSize - p.PageDataSize + 1)
	p.endData   = strconv.Itoa(p.Currentpage * p.PageDataSize)
}

//初始化
func (p *PageLinks) init()  {
	//当前请求页
	if p.Currentpage == 0 {
		p.Currentpage = 1
	}

	//默认分页数
	if p.LinkItemCount == 0 {
		p.LinkItemCount = DEFAULT_LINK_ITEM_COUNT
	}

	//首页
	if p.FirstPageText == "" {
		p.FirstPageText = DEFAULT_FIRST_PAGE_TEXT
	}

	//上一页
	if p.PrePageText == "" {
		p.PrePageText = DEFAULT_PRE_PAGE_TEXT
	}

	//下一页
	if p.NextPageText == "" {
		p.NextPageText = DEFAULT_NEXT_PAGE_TEXT
	}

	//尾页
	if p.LastPageText == "" {
		p.LastPageText = DEFAULT_LAST_PAGE_TEXT
	}
}