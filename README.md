# go-tcknverify
Go dilinde KPS servislerini kullanarak kimlik doğrulaması yapar.

## Kurulum

```bash
go get github.com/sercanarga/go-tcknverify
```

## Kullanım

`Validate` fonksiyonu `boolean` bir değer döndürür. Hata yakalama için fonksiyonunun gönderdiği 2. parametre kullanılabilir.

```go
verify, _ := tcknverify.Validate(
  &tcknverify.ValidateData{
    TCKimlikNo: "11111111111",
    Ad:         "Sercan",
    Soyad:      "Arğa",
    DogumYili:  "1995",
  })
  
if verify {
	fmt.Println("TC Kimlik numarası doğru")
} else {
	fmt.Println("Kimlik doğrulanamadı")
}
```
