package request

//func TestGetHeader_ValidHeader(t *testing.T) {
//req := httptest.NewRequest("GET", "/auth", nil)

//key := "header_key"
//expectedValue := "header_value"

//req.Header.Add(key, expectedValue)

//actualValue, err := GetHeader(req, key)

//if err != nil {
//t.Fatalf("GetHeader returned error: %v",
//err)
//}

//// Check the status code is what we expect.
//if expectedValue != actualValue {
//t.Fatalf("GetHeader returned wrong value: expected %v, actual %v",
//expectedValue, actualValue)
//}
//}

//func TestGetHeader_MissingHeader(t *testing.T) {
//req := httptest.NewRequest("GET", "/auth", nil)

//expectedValue := headerMissing

//key := "header_key"

//_, err := GetHeader(req, key)

//if err == nil {
//t.Fatalf("GetHeader expected error, but it was nil")
//}

//actualValue := err.Error()

//// Check the status code is what we expect.
//if expectedValue != actualValue {
//t.Fatalf("GetHeader returned wrong value: expected %v, actual %v",
//expectedValue, actualValue)
//}
//}
