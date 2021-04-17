package toolfunc

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"isotime"
	"math"
	"netutil"
	"os"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"testing"
	"time"
)

func TestBoscriptSeprate(t *testing.T) {
	fmt.Println(BoscriptSeperate("{a=[23432],b=[234324]}"))
	fmt.Println(BoscriptSeperate("{a=[23432],{b=[234324]}}"))
	fmt.Println(BoscriptSeperate("{a=[23432],[(adsf)sdfd(dlkfd)]}"))
	fmt.Println(BoscriptSeperate("{a=[23432],[(adsf)sdfd(dlkfd),(lsdfj)sdklf(sdf)]}"))
	fmt.Println(BoscriptSeperate("{a=[23432],{(adsf)sdfd(dlkfd),(lsdfj)sdklf(sdf)},sdflkf=23434}"))
	fmt.Println(BoscriptSeperate("{a=[23432],(sdfdsf=3244)kdsfdf{(adsf)sdfd(dlkfd),(lsdfj)sdklf(sdf)},sdflkf=23434}"))
	fmt.Println(BoscriptSeperate("{a=[23432],kdsfdf{(adsf)sdfd(dlkfd),(lsdfj)sdklf(sdf)},sdflkf=23434}"))
	fmt.Println(BoscriptSeperate("{a=[23432],[(adsf)sdfd(dlkfd),(lsdfj)sdklf(sdf)]dsfdsf[(sdfdsf)sdf(sdf),(dsfd)sdfd(sdfd)],sdfd=234}"))
}

func TestOrderFileNextKey(t *testing.T) {
	dd := []byte("连")
	fmt.Println("dd:", dd)
	//should 36830
	fmt.Println(U8FirstUChar(dd))
	fmt.Println(Contain([]int{2, 4}, 3))
	fmt.Println(SeperateUseList("sdalfk(dsf)sdfdsf(sdfdsf)sdfdsf", []string{"(", ")"}))
}

func TestFindValInRange(t *testing.T) {
	rng1 := []float64{0.1, 0.2, 0.2, 0.3, 0.3, 0.4}
	if RangeIndex(rng1, 0.2) != 1 {
		panic("error")
	}
	if RangeIndex(rng1, 0.25) != 1 {
		panic("error")
	}
	if RangeIndex(rng1, 0.1) != 0 {
		panic("error")
	}
	fmt.Println("0.4")
	if RangeIndex(rng1, 0.4) != -1 {
		panic("error")
	}
	fmt.Println("0.39")
	if RangeIndex(rng1, 0.39) != 2 {
		panic("error")
	}
	fmt.Println("0.05")
	if RangeIndex(rng1, 0.05) != -1 {
		panic("error")
	}
	rng1 = []float64{0.1, 0.2, 0.2, 0.3, 1, 1.4}
	if RangeIndex(rng1, float64(1)) != 2 {
		panic("error")
	}
	if RangeIndex(rng1, float64(0.9)) != -1 {
		panic("error")
	}
}

func testlen(dd []byte) {
	fmt.Println(len(dd))
}

func TestAttachmentParse(t *testing.T) {
	ctt, head, _, httpcode, _ := netutil.UrlGet("http://www.baike.com/wiki/辐射", []string{}, false, []string{"User-Agent", "Mozilla/5.0 (Windows; U; Windows NT 5.1; zh-CN; rv:1.9.1.6) Gecko/20091201 Firefox/3.5.6;Mozilla/5.0 (Windows NT 6.3; Win64; x64; rv:64.0) Gecko/20100101 Firefox/64.0"}, nil, 0, 0, nil)
	ioutil.WriteFile("ctt.txt", ctt, 0666)
	//head.Add("Content-Disposition", "attachment; filename*=UTF-8''weird%20%23%20%e2%82%ac%20%3D%20%7B%20%7D%20%3B%20filename.txt")
	fmt.Println("httpcode:", httpcode, head)
	fmt.Println(HttpAttachmentName(head))

	head.Add("Content-Disposition", "attachment; filename=weird%20%23%20%e2%82%ac%20%3D%20%7B%20%7D%20%3B%20filename.txt")
	fmt.Println("httpcode:", httpcode, head)
	fmt.Println(HttpAttachmentName(head))
	fmt.Println(FilePathAddDirSuffix("C:\\32434.txt", "kkkk"))

	fmt.Println(DomainMatch("http://www.163.com/", "http://mss.163.com/"))
	fmt.Println(DomainMatch("http://www.164.com.cn/", "http://mss.163.com.cn/"))
	fmt.Println(GetUrlDomain("http://www.164.com.cn/"))
	fmt.Println(GetUrlDomain("http://lebb.cc/"))

	fmt.Println(RangeIndex([]int64{0, 1, 1, 2, 2, 3, 3, 4, 4, 5, 5, 6, 6, 7, 7, 8, 8, 9, 9, 10, 10, 11, 11, 12, 12, 13, 13, 14, 14, 15, 15, 16, 16, 17, 17, 18, 18, 19}, int64(18)))
	fmt.Println(ToAbsolutePath(".././../aabb.cc/ddd.txt"))
	fmt.Println(ToAbsolutePath("C:/sdfdf/aabb.cc/ddd.txt"))
	fmt.Println(ToAbsolutePath("/uuuu/ddd.txt"))
	fmt.Println(ToAbsolutePath("wwww/ddd.txt"))
	fmt.Println("is english:", IsEnglish([]byte(`











Yahoo Sports Tourney Pick'em













Prosecutors in Palm Beach County, Florida, have offered New England Patriots owner Robert Kraft a chance to avoid prosecution for soliciting prostitution and having his charges dropped, the Wall Street Journal reported on Tuesday.

Called deferred prosecution, it would require that Kraft admit that he could be proved guilty at trial.

Kraft is not the only man charged who has been presented with this offer, the WSJ reports.

It is unclear whether Kraft will accept the agreement.

The proposed deferred prosecution deal comes with caveats: completion of an education course about prostitution, 100 hours of community service, screening for sexually transmitted diseases and payment of some court costs.

But Kraft and his attorney would also review the evidence in the case and agree that if it were to go to trial, the state would be able to prove guilt.

Kraft has been charged with two counts of soliciting prostitution; he was one of several men charged after a months-long investigation by Jupiter police, who were looking into alleged human trafficking.

[Best bracket wins $1M: Enter our free contest now! | Printable bracket]

A statement from a Kraft spokesman at the time of the charges being announced proclaimed his innocence.

Jupiter police say they have Kraft on surveillance video entering Orchids of Asia day spa, paying for services and receiving oral sex on two occasions, including on the day of the AFC championship game.

If Kraft accepted the deal and charges are dropped, he could still be punished by the NFL under the personal conduct policy. The league has reinforced that the policy applies to all members of the league, including team owners.
`)))
	fmt.Println(UrlToRegex("http://www.baidu.com/aaa/23434?wd=%AF%3G%LL&jsdfjl=32434"))
	fmt.Println(UrlToRegex("http://www.mzjt.com.cn/plus/"))
	fmt.Println("public ip:", GetPulicIPByIp138())
	urlendata := UrlDataEncode("%s/jkdf_*(极度疯狂")
	fmt.Println(urlendata)
	fmt.Println(UrlDataDecode(urlendata))
	urlen := UrlEncode("http://ww.的咖啡机.co%s/jk=d&f=_*(极度疯狂%SS")
	fmt.Println(urlen)
	fmt.Println(UrlDecode(urlen))

	ctt1, ctt1e := ioutil.ReadFile("toolfunc.go")
	ctt2, ctt2e := ReadFile("toolfunc.go")
	if ctt1e == nil && ctt2e == nil {
		if bytes.Compare(ctt1, ctt2) != 0 {
			panic("read file error!")
		}
	}

	attrbitlenbt, _ := ReadFile("pages/pagesdb/_attrbitlen")
	attrbitlen2, _ := strconv.ParseInt(string(attrbitlenbt), 10, 32)
	if 16 != attrbitlen2 {
		panic("attribute bit length error!")
	}
	fmt.Println(GetUrlDomain("http://www.baidu.com/dsfdsa/sdfdsf"))
	fmt.Println(GetNoTagHomeUrl("http://www.baidu.com/dsfdsa/sdfdsf"))
	fmt.Println(GetUrlDomain("https://www.baidu.net.cn/dsfdsa/sdfdsf"))
	fmt.Println(GetNoTagHomeUrl("https://www.baidu.net.cn/dsfdsa/sdfdsf"))
	fmt.Println(StdHomeUrl("https://www.baidu.net.cn/dsfdsa/sdfdsf"))
	fmt.Println(StdHomeUrl("https://www.baidu.net.cn"))
	fmt.Println(UrlEncode("https://www.baidu.net.cn/sdfd_%sdlkfd.html"))
	fmt.Println(UrlDataEncode("https://www.baidu.net.cn/sdfd_%sdlkfd.html"))

	var bigm sync.Map
	for i := 0; i < 3000; i++ {
		bigm.Store(string(RandAlpha(4, 24)), RandAlpha(4, 24))
	}
	t1 := time.Now().Nanosecond()
	fmt.Println(len(StringBytesMapToBytes(bigm)))
	fmt.Println(float32(time.Now().Nanosecond()-t1) / 1e9)
	fmt.Println(GetUrlDomain("https://www.huaxianer.com/zhishi/86/"))

	fmt.Println(FilePathReplaceDir("sdfad/aaa/3433.txt", "bbb"))

	fmt.Println(FilePathAppendDir(`H:\cangtianweb\pages\pagesdb\index\_curdir`, "3"))
	fmt.Println(UrlDataEncode("42	区\n\r委组织部 \x00/sdflds"))
	aa := []string{"kldsafjlds", "9394", "8932jk"}
	StringQuickSortAsc(aa, 0, 2)
	fmt.Println(aa)
	fmt.Println(ToAbsolutePath("H:\\324\\324324.23432"))
	fmt.Println(CurDir())
	fmt.Println(IsoTimeToSec("00:00:33"))
	fmt.Println(TimeZoneSec(CountryZones()["CN"][0]))
	fmt.Println(HttpTimeToTime("08 May 2018 06:17:00 GMT", CountryZones()["CN"][0]))
	fmt.Println(Float64RoundToStr(Round(Float64FromStr("174426")/Float64FromStr("306819"), 3), 3))
	fmt.Println((time.Now().Local().Unix() - TodayLocalSec()) % (24 * 3600))
	fmt.Println(time.Now().Local().Add(-time.Duration(TodayLocalSec())*time.Second).Unix() % (24 * 3600))
	tt := time.Now().Local()
	fmt.Println(tt.Unix(), tt.Local().Unix())
	fmt.Println(Float32RoundToStr(3.13, 5))
	fmt.Println(IsNumber([]byte("-3.324E+3.6")))
	fmt.Println(IsIPV4([]byte("3.6.34.4")))
	fmt.Println(IsIsoDateTime([]byte("1904-07-08T23:44:23")))

	numbt := []byte("45.3454E-3*Sin(34)")
	inds := regexp.MustCompile("(?ism)^[+-]?[0-9]+(\\.[0-9]+)?([Ee][+-]?[0-9]+(\\.[0-9]+)?)?").FindAllSubmatchIndex(numbt, -1)
	for i := 0; i < len(inds[0]); i += 2 {
		if inds[0][i] == -1 {
			continue
		}
		fmt.Println(string(numbt[inds[0][i]:inds[0][i+1]]))
	}

	var stack sync.Map
	var expri int
	val, vale := DoPartCalc([]byte("3+8"), &expri, nil, &stack, false)
	if val.(string) != "11" {
		panic("error")
	}
	expri = 0
	val, vale = DoPartCalc([]byte("3*(8-5)"), &expri, nil, &stack, false)
	fmt.Println(val, vale)
	if val.(string) != "9" {
		panic("error")
	}
	expri = 0
	val, vale = DoPartCalc([]byte("3*(8-5)-4"), &expri, nil, &stack, false)
	fmt.Println(val, vale)
	if val.(string) != "5" {
		panic("error")
	}
	expri = 0
	val, vale = DoPartCalc([]byte("3*(8-5)/(4-(5-4))"), &expri, nil, &stack, false)
	fmt.Println(val, vale)
	if val.(string) != "3" {
		panic("error")
	}
	expri = 0
	val, vale = DoPartCalc([]byte("3*(8-5)/(4-(5-4)"), &expri, nil, &stack, false)
	fmt.Println(val, vale)
	if vale != false {
		panic("error")
	}
	expri = 0
	val, vale = DoPartCalc([]byte("3e24+2e23"), &expri, nil, &stack, false)
	fmt.Println(val, vale)
	if val.(string) != "3.2e+24" {
		panic("error")
	}
	expri = 0
	val, vale = DoPartCalc([]byte("3*(8-5)/(4-(5-4)"), &expri, nil, &stack, true)
	fmt.Println(val, vale)
	if vale != false {
		panic("error")
	}

	dd := GetTagContent("abcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyz", "c", "f")
	fmt.Println(dd)

	fmt.Println(GetAndTruncateTagContent("abcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyz", "c", "f"))

	dd2 := GetRecursiveTagContent("<body><div dd=aa><div kd=3243><div dkf=34></div><div kdf=3434></div></div><div kkg=4545></div></div></body>", "<div dd=aa", "<div", "</div>")
	fmt.Println(dd2)

	CreateDir("sdfds/dsafdsa/sdafsdafdsa/dsafs/333", 0666)

	fmt.Println(StdFileName("a:aa?/??|\\|bbb\\|/?<>dd d>c*c\r\n\tc"))

	fmt.Println(Count([]byte("a  "), []byte(" ")))
	fmt.Println(Count([]byte("a b c d e "), []byte(" ")))
	fmt.Println(string(Truncate([]byte("a b c d "), []byte(" "), 3)))
	fmt.Println(string(Truncate([]byte("a b c d e "), []byte(" "), 5)))

	fmt.Println(IsoDateMilisecTimeCompare("2019-12-03 12:33:44.123", "2019-12-03 12:33:44.124"))

	aaa := []string{"", "", "", "a", "dsfdsaf", "", "", "dsfdf", "dsieidkkdd", "", "", "dsfdsafdsfdsaf"}
	aaabt := StringListToBytes(aaa)
	fmt.Println(StringListFromBytes(aaabt))

	MoveFile("toolfunc.test.exe", "toolfunc.test.exe2")

	ImageResizeBig("M:\\errorimage\\weibor_icon副本.png", "abc.jpg", 100)
	fmt.Println("kdkkdkkdt")
	fmt.Println(GetRepeatCountList([]string{"z", "a", "a", "a", "b", "b", "b", "d", "c", "c", "y"}))
	fmt.Println("dddd", string(RandUpperNum(10, 10)))

	fmt.Println(GetFileModifiedMicroTime("ctt.txt"))
	fmt.Println(GetFileModifiedNanoTime("ctt.txt"))
	fmt.Println(isotime.MicrosecToIsoDateTime(GetFileModifiedMicroTime("ctt.txt")))
	fmt.Println(isotime.DayToDate(int32(GetFileModifiedMicroTime("ctt.txt") / (24 * 3600000000))))

	fmt.Println(FileSha1("doc.go"))

	kvbuf := NewKeyValueBuf(nil)
	kvbuf.Append("dsfsdf", "dsfdsf")
	kvbuf.Append("2222", "33333")
	kvbuf.Reset()
	for kvbuf.Next() {
		fmt.Println(kvbuf.Key(), kvbuf.Value())
	}

	testlen([]byte("dkroroobgh"))

	var eddd []byte = nil
	fmt.Println(eddd == nil)
	fmt.Println(bytes.Compare([]byte("ab"), []byte("abc")))
	var dd3 sync.Map
	dd3.Store("", 33)
	fmt.Println(dd3.Load(""))
	var kdk []byte
	fmt.Println(kdk == nil)

	var ms, ms2 MemStatus
	ms.All = 1
	ms.Free = 2
	ms.Self = 3
	ms.Used = 4
	msctt, msctte := json.Marshal(ms)
	fmt.Println(msctte)
	fmt.Println(msctt)
	fmt.Println(string(msctt))
	fmt.Println(len(msctt))
	fmt.Println(ms)
	json.Unmarshal(msctt, &ms2)
	fmt.Println(ms2)
	url := "http://www.baidu.com/index.php"
	fmt.Println(url[strings.LastIndex(url, "/"):])
	arturlf, _ := os.OpenFile("testfile1.txt", os.O_CREATE|os.O_WRONLY, 0666)
	arturlf.Seek(0, os.SEEK_END)
	arturlf.Write([]byte("1"))
	arturlf.Close()

	type UrlTitle struct {
		Url   string
		Title string
	}
	var infols []*UrlTitle
	for i := 0; i < 2; i++ {
		infols = append(infols, &UrlTitle{Url: "dsfdsf", Title: "dsfdsfds435435"})
	}
	infolsctt, _ := json.Marshal(infols)
	fmt.Println(string(infolsctt))

	fmt.Println(regexp.MustCompile("^https?://[a-z0-9\\-]+[.][a-z0-9.\\-]+/.*").Match([]byte("[forbid]:^https://down.ali213.net/pcgame/.*")))
	//fmt.Println(DecListToString("212 29 140 217 143 0 178 4 233 128 9 152 236 248 66 126"))

	var j interface{}
	jsonstr := `[432,"3243214",4324,"dsflksdfj"]`
	err := json.Unmarshal([]byte(jsonstr), &j)
	fmt.Println(err)
	fmt.Println(j)
	fmt.Println(reflect.TypeOf(j.([]interface{})[1]))
	bc := make([]byte, 8)
	binary.BigEndian.PutUint64(bc, math.Float64bits(1.8446744073709552e+19))
	fmt.Println("float64", bc)

	i := int64(-2)

	fmt.Println(i)

	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(i))

	fmt.Println(b)

	i = int64(binary.BigEndian.Uint64(b))
	fmt.Println(i)
	fmt.Println(len(map[string]string{"safsdf": "lskdfkdlf"}))

	kks := float64(-45435)
	fmt.Println(int64(kks))
	fmt.Println(float64(^uint64(0)))
	fmt.Println(math.Float64frombits(binary.BigEndian.Uint64([]byte{255, 255, 255, 255, 255, 255, 255, 255})))
	fmt.Println(ValCompare([]interface{}{"aa", "bb", 44}, []interface{}{"aa", "bb", 44}))
	fmt.Println(uint64(1) << 63)
	fmt.Println(bytes.Compare([]byte{0x11, 0x12, 0x11, 0xAB}, []byte{0x11, 0x12, 0x11}))
	ff, ffe := os.OpenFile("test1.txt", os.O_CREATE|os.O_RDWR, 0666)
	if ffe == nil {
		ff.Seek(0, os.SEEK_END)
		ff.Write([]byte("a"))
		ff.Close()
	}

	ff4, _ := os.OpenFile("test1.txt", os.O_RDONLY, 0666)
	kl := make([]byte, 1000)
	ff4l, _ := ff4.Read(kl)
	fmt.Println(ff4l, len(kl), kl)
	ff4l, _ = ff4.Read(kl)
	fmt.Println(ff4l, len(kl), kl)
	ff4.Close()

	MakePathDirExists("aa/bb/cc")
}

/**/
