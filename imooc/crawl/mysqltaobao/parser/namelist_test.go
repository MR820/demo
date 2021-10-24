/**
 * Createby GoLand
 * User xzw jsjxzw@163.com
 * Date 2020/5/29
 * Time 4:54 下午
 */

package parser

import (
	"io/ioutil"
	"testing"
)

func TestParseNameList(t *testing.T) {
	contents, err := ioutil.ReadFile("namelist_test_data.html")
	if err != nil {
		panic(err)
	}

	//fmt.Printf("%s", contents)

	result := ParseNameList(contents)

	const resultSize = 69
	if len(result.Requests) != resultSize {
		t.Errorf("result should have %d requests;but had %d", resultSize, len(result.Requests))
	}
	if len(result.Items) != resultSize {
		t.Errorf("result should have %d items;but had %d", resultSize, len(result.Items))
	}
}
