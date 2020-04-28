package main

import (
	"net/http/httptest"
	"testing"
)

func BenchmarkRequest(b *testing.B) {
	b.ReportAllocs()
	r := httptest.NewRequest("GET", "/Promotions/?Years=33&Balance=30050&Rating=600&Age=19&AccountType=Blue", nil)
	for i := 0; i < b.N; i++ {
		w := httptest.NewRecorder()
		getPromotions(w, r)
	}
}

func BenchmarkRequest2(b *testing.B) {
	b.ReportAllocs()
	//r := httptest.NewRequest("GET", "/Promotions/?Years=33&Balance=30050&Rating=600&Age=19&AccountType=Blue", nil)
	var DD int = 0
	for i := 0; i < b.N; i++ {
		DD+=1
		//fmt.Println(DD)
		//w := httptest.NewRecorder()
		//getPromotions(w, r)
	}

}
