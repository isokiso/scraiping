package main

import (
    "github.com/sclevine/agouti"
    "log"
    "time"
    "regexp"
    "strings"
    "./myconf"
    //"fmt"
    //"strconv"
)

func main() {
    // ブラウザはChromeを指定して起動
    driver := agouti.ChromeDriver(agouti.Browser("chrome"))
    if err := driver.Start(); err != nil {
        log.Fatalf("Failed to start driver:%v", err)
    }
    defer driver.Stop()

    page, err := driver.NewPage()
    if err != nil {
        log.Fatalf("Failed to open page:%v", err)
    }
    // ログインページに遷移
    if err := page.Navigate("https://pos.toshin.com/SSO1/Examiner/SSOLogin/nhqlogin.aspx"); err != nil {
        log.Fatalf("Failed to navigate:%v", err)
    }
    // ID, Passの要素を取得し、値を設定
    identity := page.FindByID("txtUserID")
    password := page.FindByID("txtPassword")
    identity.Fill(myconf.Myid)
    password.Fill(myconf.Mypassword)
    time.Sleep(3 * time.Second)
    // formをサブミット
    if err := page.FindByName("ibLogin").Click(); err != nil {
        log.Fatalf("Failed to login:%v", err)
    }
    page.FindByID("dgMenu__ctl2_lbLink").Click()


    //MacOSだとボタンが表示されないので以下でurlを取得
    content, err := page.HTML()
    content = strings.Replace(content, "\n", "", -1)
    r := regexp.MustCompile(`<frame src="http.*?"`)//<frame src="http.から"までを探して取得
    first := r.FindStringSubmatch(content)
    zero := strings.Replace(first[0], `<frame src="`, "",1)
    url := strings.Replace(zero, `"`, "",1)
    page.Navigate(url)


    log.Println(page.FindByXPath(`//*[@id="GridView1"]/tbody/tr[2]/td[1]/a`).Click())
    //page.FindByID("lblSearch").Click()
    //or  log.Println(page.FindByXPath(`//*[@id="lblSearch"]`).Click())
    time.Sleep(18*time.Second)
/*
    content, err = page.HTML()
    content = strings.Replace(content, "\n", "", -1)
    s := regexp.MustCompile(`<span id="SearchFeedBack" style="color:Red;font-family:Arial;font-size:Small;">.*?件取得しました。`)
    uno := s.FindStringSubmatch(content)
    if len(uno) != 0{
      dos := strings.Replace(uno[0], `<span id="SearchFeedBack" style="color:Red;font-family:Arial;font-size:Small;">`, "",1)
      number := strings.Replace(dos, `件取得しました。`, "",1)
      fmt.Println(number)
      num, _ := strconv.Atoi(number)
      for i := 2; i < 2+num; i++ {
        log.Println(page.FindByXPath(`//*[@id="GridView1"]/tbody/tr[` + strconv.Itoa(i) + `]/td[11]/a`).Click())
      }
    }

*/
    //time.Sleep(1 * time.Second)
}
