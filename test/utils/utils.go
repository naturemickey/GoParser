package utils

import "io/ioutil"

func WalkDir(path string, test_fun func(string2 string)) {
	dir, err2 := ioutil.ReadDir(path)
	if err2 != nil {
		println(err2.Error())
	}
	for _, fileInfo := range dir {
		if fileInfo.IsDir() {
			WalkDir(path+fileInfo.Name()+"/", test_fun)
		} else {
			test_fun(path + fileInfo.Name())
		}
	}
}
