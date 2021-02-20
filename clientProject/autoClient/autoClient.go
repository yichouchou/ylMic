package autoClient

import (
	"fmt"
	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"
	"log"
	"strconv"
	"time"
	"ylMic/clientProject/models"
)

// StartChrome 启动谷歌浏览器headless模式
type StandGoodsList struct {
	StandList []*models.StandGoods
}

var standList *StandGoodsList = nil

func getStandList() *StandGoodsList {
	return standList
}

type GoodsHashMap struct {
	GoodsMap map[string][]*models.Goods
}

var goodsMap *GoodsHashMap = nil

func getGoodsMap() *GoodsHashMap {
	return goodsMap
}

func StartChrome() {

	opts := []selenium.ServiceOption{}

	caps := selenium.Capabilities{

		"browserName": "chrome",
	}

	// 禁止加载图片，加快渲染速度

	imagCaps := map[string]interface{}{

		"profile.managed_default_content_settings.images": 2,
	}

	chromeCaps := chrome.Capabilities{

		Prefs: imagCaps,

		Path: "",

		Args: []string{

			"--headless", // 设置Chrome无头模式

			"--no-sandbox",

			"--user-agent=Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_2) AppleWebKit/604.4.7 (KHTML, like Gecko) Version/11.0.2 Safari/604.4.7", // 模拟user-agent，防反爬

		},
	}

	caps.AddChrome(chromeCaps)

	// 启动chromedriver，端口号可自定义

	service, err := selenium.NewChromeDriverService("./chromedriver", 9516, opts...)

	if err != nil {

		log.Printf("Error starting the ChromeDriver server: %v", err)

	}

	// 调起chrome浏览器

	webDriver, err := selenium.NewRemote(caps, fmt.Sprintf("http://localhost:%d/wd/hub", 9516))

	if err != nil {

		panic(err)

	}

	//目标网站

	targeUrl := "https://qqsg.pc9527.vip"

	// 导航到目标网站

	err = webDriver.Get(targeUrl)

	if err != nil {

		panic(fmt.Sprintf("Failed to load page: %s\n", err))

	}
	areaBut, err := webDriver.FindElement(selenium.ByCSSSelector, "#app > div > div > div > div.header > div.el-select.regionSelect.el-select--mini")

	areaBut.Click()

	time.Sleep(3000)

	areaSelect, _ := webDriver.FindElement(selenium.ByCSSSelector, "body > div.el-select-dropdown.el-popper > div.el-scrollbar > div.el-select-dropdown__wrap.el-scrollbar__wrap > ul")
	areaNames, _ := areaSelect.FindElements(selenium.ByTagName, "li")

	for _, v := range areaNames {
		text, _ := v.Text()
		if text == "得陇" {
			v.Click()
		}
	}

	fmt.Println(areaBut.Text())
	time.Sleep(2)

	//登录框
	loginBut, _ := webDriver.FindElement(selenium.ByCSSSelector, "#app > div > div > div > div.header > div.btnBox > button.el-button.loginBtn.el-button--primary.el-button--mini > span")

	loginBut.Click()

	userNameInput, _ := webDriver.FindElement(selenium.ByCSSSelector,
		"#app > div > div > div > div:nth-child(3) > div > div > div.el-dialog__body > form > div:nth-child(1) > div > div > input")

	passwordInput, _ := webDriver.FindElement(selenium.ByCSSSelector,
		"#app > div > div > div > div:nth-child(3) > div > div > div.el-dialog__body > form > div:nth-child(2) > div > div > input")

	userNameInput.SendKeys("a10725542")

	passwordInput.SendKeys("lzy123...")
	time.Sleep(2)

	loginButton, _ := webDriver.FindElement(selenium.ByCSSSelector,
		"#app > div > div > div > div:nth-child(3) > div > div > div.el-dialog__body > form > div:nth-child(3) > div > button")

	loginButton.Click()
	time.Sleep(3)
	//查看历史--
	//可以不用此步
	history, _ := webDriver.FindElement(selenium.ByCSSSelector,
		"#app > div > div > div > div.header > div.btnBox > button:nth-child(3) > span")
	print(history.Text())

	thingsInput, _ := webDriver.FindElement(selenium.ByCSSSelector,
		"#app > div > div > div > div.header > div:nth-child(5) > div > input")

	standGoodsList := getStandList()
	goodsHashMap := getGoodsMap()
	for _, v := range standGoodsList.StandList {
		standGoodsName := v.StandGoodsName
		thingsInput.SendKeys(standGoodsName)
		time.Sleep(3000)
		thingsWantFound, _ := webDriver.FindElement(selenium.ByCSSSelector,
			"body > div.el-select-dropdown.el-popper > div.el-scrollbar > div.el-select-dropdown__wrap.el-scrollbar__wrap > ul > li > span")
		thingsWantFound.Click()
		time.Sleep(2000)
		thingsInfoList, _ := webDriver.FindElement(selenium.ByCSSSelector, "#pane-booth > div.el-table.el-table--fit.el-table--enable-row-hover.el-table--enable-row-transition.el-table--mini > div.el-table__body-wrapper.is-scrolling-none > table > tbody")
		thingsTrList, _ := thingsInfoList.FindElements(selenium.ByTagName, "tr")
		goods := []*models.Goods{}
		for _, thingsTr := range thingsTrList {
			good := new(models.Goods)
			thingsTd, _ := thingsTr.FindElements(selenium.ByTagName, "td")
			text, _ := thingsTd[7].Text()
			prices, _ := strconv.Atoi(text)
			standPrice, _ := strconv.Atoi(v.StandGoodsPrice)
			if prices <= standPrice {
				good.Area, _ = thingsTd[1].Text()
				good.Price = text
				good.AreaLikeDeLong = "delong"
				good.ShopName, _ = thingsTd[4].Text()
				good.GoodsName, _ = thingsTd[2].Text()
				good.Time, _ = thingsTd[6].Text()
				good.SaleName, _ = thingsTd[3].Text()
				good.Number, _ = thingsTd[5].Text()
				goods = append(goods, good)
			}
		}
		goodsHashMap.GoodsMap[standGoodsName] = goods

	}

	//log.Println(webDriver.PageSource())

	defer service.Stop() // 停止chromedriver

	defer webDriver.Quit() // 关闭浏览器

}
