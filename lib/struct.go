package lib

type JobData struct {
	WebSite            string //网站
	City               string //城市
	JobTitle           string //职位
	Salary             string //薪水
	Address            string //地址
	Empirical          string //经验
	Education          string //学历
	CompanyName        string //公司名
	CompanyType        string //公司类型
	Person             string //公司人数
	RecruitName        string //招骋人
	RecruitPosition    string //招骋人title
	UpdateTime         string //职位更新时间
	FinancingSituation string //融资情况
	CreateTime         string //创建时间
}

type Proxy struct {
	ID        int    //ID
	Url       string //代理
	Template  string //模板
	ParamList string //参数列表
	List      string //代理列表
	CheckStr  string //验证代理通过后的字符
}
