package jabob_test

import (
	"encoding/json"
	"testing"

	jabob "github.com/go-marshaltemabu/go-jabob"
)

type structureForTesting01 struct {
	Abc int32  `json:"abc"`
	Def bool   `json:"def"`
	Ghi string `json:"ghi"`
}

func checkTestData01(t *testing.T, targetVal *structureForTesting01) {
	if targetVal.Abc != 10 {
		t.Errorf("unexpect value (TargetVal.Abc != 10): %v", targetVal.Abc)
	}
	if targetVal.Def != true {
		t.Errorf("unexpect value (TargetVal.Def != true): %v", targetVal.Def)
	}
	if targetVal.Ghi != "hi" {
		t.Errorf("unexpect value (TargetVal.Ghi != \"hi\"): %v", targetVal.Ghi)
	}
}

func TestMarshalJSON01(t *testing.T) {
	targetInput := struct {
		TargetVal jabob.BytesJSON `json:"target_v"`
	}{
		TargetVal: jabob.BytesJSON{
			Bytes: []byte("{\"abc\":10,\"def\":true,\"ghi\":\"hi\"}"),
		},
	}
	buf, err := json.Marshal(&targetInput)
	if nil != err {
		t.Errorf("cannot json.Marshal input into JSON: %v", err)
		return
	}
	var targetOutput struct {
		TargetVal structureForTesting01 `json:"target_v"`
	}
	if err = json.Unmarshal(buf, &targetOutput); nil != err {
		t.Errorf("cannot json.Unmarshal input from JSON: %v", err)
		return
	}
	checkTestData01(t, &targetOutput.TargetVal)
}

func TestUnmarshalJSON01(t *testing.T) {
	var targetOutput struct {
		TargetVal jabob.BytesJSON `json:"target_v"`
	}
	if err := json.Unmarshal(
		[]byte("{\"target_v\":{\"abc\":10,\"def\":true,\"ghi\":\"hi\"}}"),
		&targetOutput); nil != err {
		t.Errorf("cannot json.Unmarshal input from JSON: %v", err)
		return
	}
	var targetVerify structureForTesting01
	if err := json.Unmarshal(targetOutput.TargetVal.Bytes, &targetVerify); nil != err {
		t.Errorf("cannot json.Unmarshal verify copy: %v", err)
		return
	}
	checkTestData01(t, &targetVerify)
}

func TestMarshalFrom01(t *testing.T) {
	targetData := structureForTesting01{
		Abc: 10,
		Def: true,
		Ghi: "hi",
	}
	var targetOutput struct {
		TargetVal jabob.BytesJSON `json:"target_v"`
	}
	if err := targetOutput.TargetVal.MarshalFrom(targetData); nil != err {
		t.Errorf("cannot pack data into JSON with MarshalFrom: %v", err)
		return
	}
	t.Logf("result of MarshalFrom: %s", targetOutput.TargetVal.String())
	var targetVerify structureForTesting01
	if err := json.Unmarshal(targetOutput.TargetVal.Bytes, &targetVerify); nil != err {
		t.Errorf("cannot json.Unmarshal verify copy: %v", err)
		return
	}
	checkTestData01(t, &targetVerify)
}

func TestUnmarshalInto01(t *testing.T) {
	targetInput := struct {
		TargetVal jabob.BytesJSON `json:"target_v"`
	}{
		TargetVal: jabob.BytesJSON{
			Bytes: []byte("{\"abc\":10,\"def\":true,\"ghi\":\"hi\"}"),
		},
	}
	var targetVerify structureForTesting01
	if err := targetInput.TargetVal.UnmarshalInto(&targetVerify); nil != err {
		t.Errorf("cannot fill target structure with UnmarshalInto: %v", err)
		return
	}
	checkTestData01(t, &targetVerify)
}
