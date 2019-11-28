package libs

import "testing"

func TestBase62(t *testing.T) {
	res := UintToBase62(62)
	if res == "10"{
		t.Logf("success, expect:%v, got:%v\n", "10", res)
	} else{
		t.Errorf("error, expect:%v, got:%v\n", "10", res)
	}
	res = UintToBase62(73)
	if res == "1B"{
		t.Logf("success, expect:%v, got:%v\n", "1B", res)
	} else{
		t.Errorf("error, expect:%v, got:%v\n", "1B", res)
	}

	_, err := Base62ToUint("1+")
	if err != nil{
		t.Logf("success, %v", err)
	} else{
		t.Errorf("error")
	}

	value, _ := Base62ToUint("1B")
	if value == 73{
		t.Logf("success, expect:%v, got:%v\n", 73, value)
	} else{
		t.Errorf("error, expect:%v, got:%v\n", 73, value)
	}

	value, _ = Base62ToUint("EqV5vuleU")
	str := UintToBase62(value)
	if str == "EqV5vuleU" {
		t.Logf("success, expect:%v, got:%v\n", "EqV5vuleU", str)
	} else{
		t.Errorf("error, expect:%v, got:%v\n", "EqV5vuleU", str)
	}
}