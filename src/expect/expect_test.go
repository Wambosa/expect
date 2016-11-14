package expect_test

import (
    "testing"
    "expect"
)


func Test_Property_when_given_a_map_string_keyName_then_reduces_TestCase_Value(t *testing.T){
    
    expecting := expect.TestCase{
        Value: map[string]interface{}{
            "name": "goku",
        },
    }
    
    expecting.Property("name")
    
    if expecting.Value != "goku" {
        t.Errorf("Expecting Property to have reduced the TestCase.Value to goku. But it was set to: %v", expecting.Value)
    }
}

func Test_Property_when_given_a_map_string_keyName_then_stashes_original_map(t *testing.T){
    
    expecting := expect.TestCase{
        Value: map[string]interface{}{
            "technique": "kamehameha",
        },
    }
    
    expecting.Property("technique")
    
    if expecting.Map["technique"] != "kamehameha" {
        t.Errorf("Expecting Property to have stashed the original map. But it was set to: %v", expecting.Map)
    }
}

func Test_Property_when_given_multiple_string_keyName_then_successfully_makes_multiple_checks(t *testing.T){
    
    expecting := expect.TestCase{
        T: t,
        Value: map[string]interface{}{
            "name": "gohan",
            "technique": "kamehameha",
        },
    }
    
    expecting.Property("name").ToBe("gohan")
    expecting.Property("technique").ToBe("kamehameha")
}

func Test_Truthy_when_given_a_struct_then_pass_test(t * testing.T){
    
    type zWarrior struct {
        name string
    }
    
    expecting := expect.TestCase{
        T: t,
        Value: zWarrior{name: "trunks"},
    }
    
    expecting.Truthy()
}

func Test_Truthy_when_given_a_map_then_pass_test(t * testing.T){
    
    expecting := expect.TestCase{
        T: t,
        Value: map[int]string{9000:"over"},
    }
    
    expecting.Truthy()
}

func Test_Truthy_when_given_a_NONEmpty_string_then_pass_test(t * testing.T){
    
    expecting := expect.TestCase{
        T: t,
        Value: "Vegeta",
    }
    
    expecting.Truthy()
}

func Test_Truthy_when_given_a_number_greater_than_zero_then_pass_test(t * testing.T){
    
    expecting := expect.TestCase{
        T: t,
        Value: 9000,
    }
    
    expecting.Truthy()
}

func Test_Truthy_when_given_a_number_less_than_zero_then_pass_test(t * testing.T){
    
    expecting := expect.TestCase{
        T: t,
        Value: -9000,
    }
    
    expecting.Truthy()
}

func Test_Truthy_when_given_nil_then_fails_test(t * testing.T){
    
    expecting := expect.TestCase{
        Value: nil,
    }
    
    if expecting.Truthy() != false {
        t.Error("TestCase.Truthy should return FALSE when given nil")
    }
}

func Test_Truthy_when_given_an_EMPTY_string_then_fail_test(t * testing.T){
    
    expecting := expect.TestCase{
        Value: "",
    }
    
    if expecting.Truthy() != false {
        t.Error("TestCase.Truthy should return FALSE when given an empty string")
    }
}

func Test_Truthy_when_given_a_number_zero_then_fails_test(t * testing.T){
    
    expecting := expect.TestCase{
        Value: 0,
    }
    
    if expecting.Truthy() != false {
        t.Error("TestCase.Truthy should return FALSE when given a zero")
    }
}

func Test_ToBe_when_given_nil_then_compares_correctly(t *testing.T){
    expecting := expect.TestCase{
        Value: nil,
    }
    
    if !expecting.ToBe(nil) {t.Error("The TestCase.Value is not converting properly. The problem may be with expect.toString")}
    if !expecting.ToBe("") {t.Error("The TestCase.Value is not converting properly. The problem may be with expect.toString")}
    if !expecting.ToBe(0) {t.Error("The TestCase.Value is not converting properly. The problem may be with expect.toString")}    
}

func Test_ToBe_when_given_a_number_then_compares_correctly(t *testing.T){
    expecting := expect.TestCase{
        Value: 9000,
    }
    
    if !expecting.ToBe(9000) {t.Error("The TestCase.Value is not converting numbers properly. The problem may be with expect.toString")}
    
    if expecting.ToBe(0) {t.Error("The TestCase.Value is not converting numbers properly. The problem may be with expect.toString")}
    
}

func Test_ToBe_when_given_a_string_then_compares_correctly(t *testing.T){
    expecting := expect.TestCase{
        Value: "raditz",
    }
    
    if !expecting.ToBe("raditz") {t.Error("The TestCase.Value is not converting strings properly. The problem may be with expect.toString")}
    
    if expecting.ToBe("") {t.Error("The TestCase.Value is not converting strings properly. The problem may be with expect.toString")}
}

func Test_ToBe_when_given_a_map_then_compares_correctly(t *testing.T){
    expecting := expect.TestCase{
        Value: map[string]interface{}{"name": "piccolo", "tech": "special beam cannon"},
    }
    
    if !expecting.ToBe(map[string]interface{}{"name": "piccolo", "tech": "special beam cannon"}) {
        t.Error("The TestCase.Value is not converting maps properly. The problem may be with expect.toString")
        
    }
    
    if expecting.ToBe("") {t.Error("The TestCase.Value is not converting maps properly. The problem may be with expect.toString")}
}