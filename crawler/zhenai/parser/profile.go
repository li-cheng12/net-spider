package parser

import (
	"GoDemo/crawler/engine"
	"regexp"
	"GoDemo/crawler/model"
	"strconv"
)
var idRe = regexp.MustCompile(`<p class="brief-info fs14 lh32 c9f">ID：([\d]+)`)
var ageRe = regexp.MustCompile(`<td><span class="label">年龄：</span>([\d]+)岁</td>`)
var genderRe = regexp.MustCompile(`<td><span class="label">性别：</span><span field="">([^<]+)</span></td>`)
var heightRe = regexp.MustCompile(`<td><span class="label">身高：</span>([^<]+)</td>`)
var weightRe = regexp.MustCompile(`<td><span class="label">体重：</span><span field="">([^<]+)</span></td>`)
var incomeRe = regexp.MustCompile(`<td><span class="label">月收入：</span>([^<]+)</td>`)
var marriageRe = regexp.MustCompile(`<td><span class="label">婚况：</span>([^<]+)</td>`)
var educationRe = regexp.MustCompile(`<td><span class="label">学历：</span>([^<]+)</td>`)
var occupationRe = regexp.MustCompile(`<td><span class="label">职业： </span>([^<]+)</td>`)
var hokouRe = regexp.MustCompile(`<td><span class="label">籍贯：</span>([^<]+)</td>`)
var xinzuoRe = regexp.MustCompile(`<td><span class="label">星座：</span>([^<]+)</td>`)
var houseRe = regexp.MustCompile(`<td><span class="label">住房条件：</span><span field="">([^<]+)</span></td>`)
var carRe = regexp.MustCompile(`<td><span class="label">是否购车：</span><span field="">([^<]+)</span></td>`)
func ParseProfile(contents []byte, name string) engine.ParseResult {
	profile := model.Profile{}
	profile.Id = ConvertTOString(contents, idRe)
	profile.Age,_ = strconv.Atoi(ConvertTOString(contents, ageRe))
	profile.Gender = ConvertTOString(contents, genderRe)
	profile.Height,_ = strconv.Atoi(ConvertTOString(contents, heightRe))
	profile.Weight,_ = strconv.Atoi(ConvertTOString(contents, weightRe))
	profile.Income = ConvertTOString(contents, incomeRe)
	profile.Marriage = ConvertTOString(contents, marriageRe)
	profile.Education = ConvertTOString(contents, educationRe)
	profile.Occupation = ConvertTOString(contents, occupationRe)
	profile.Hokou = ConvertTOString(contents, hokouRe)
	profile.Xinzuo = ConvertTOString(contents, xinzuoRe)
	profile.House  = ConvertTOString(contents, houseRe)
	profile.Car = ConvertTOString(contents, carRe)
	profile.Name = name
	result := engine.ParseResult{
		Items:[] interface{}{profile},
	}
	return result
}
func ConvertTOString(contents []byte, re *regexp.Regexp) string {
	match := re.FindSubmatch(contents)
	if len(match) > 1 {
		return string(match[1])
	} else {
		return ""
	}
}
