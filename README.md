num2words
=========


num2words - Numbers to persian words converter in Go (Golang)

## Usage

First, import package num2words

```import github.com/esmailemami/num2words```

Convert number
```go
  str := num2words.Convert(12) // outputs "دوازده"
  ...
  str := num2words.Convert(1024) // outputs "یک هزار و بیست و چهار"
  ...
  str := num2words.Convert(-123) // outputs "منهای یکصد و بیست و سه"
```

If you want to use english words converter, you can use [Divan's Repository](https://github.com/divan/num2words)
