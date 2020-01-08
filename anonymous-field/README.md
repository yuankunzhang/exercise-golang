# Anonymous field in a struct

## Anonymous field

A field declared with a type but no explicit field name is an anonymous field, also called an embedded field or an embedding of the type in the struct.

An embedded type must be specified as a type name `T` or as a pointer to a non-interface type name `*T`, and `T` itself may not be a pointer type.

## Field promotion

A field or method `f` of an anonymous field in a struct `x` is called promoted if `x.f` is a legal selector that denotes that field or method `f`.

Promoted fields act like ordinary fields of a struct except that they cannot be used as field names in composite literals of the struct.

## Field access

```go
type Discount struct {
    percent   float32
    startTime uint64
    endTime   uint64
}

func (d Discount) Calculate(originalPrice float32) float32 {
    return originalPrice * d.percent
}

func (d Discount) IsValid(currentTime uint64) bool {
    return currentTime > d.startTime && currentTime < d.endTime
}

// ----------

type PremiumDiscount struct {
    Discount
    additional float32
}

func (d PremiumDiscount) Calculate(originalPrice float32) float32 {
    return d.Discount.Calculate(originalPrice) - additional
}
```

Given `var discount PremiumDiscount`, we can call `discount.IsValid` directly. But the `Discount.Calculate` is hidden by `PremiumDiscount.Calculate`. In this case, the anonymous field can still be referenced as part of the `struct` by using the full name of the embedded struct.

## Anonymous field cannot be pointer to pointer or pointer to interface

Because these types don't have methods. The whole point of anonymous fields is that the methods get promoted.

## Flyweight pattern (享元模式)

A flyweight is an object that minimizes memory usage by sharing as much data as possible with other similar objects; it is a way to use objects in large numbers when a simple repeated representation would use an unacceptable amount of memory.

## Notes

Go offers the embedding mechanism as an alternative to the inheritance mechanism in traditional object oriented languages. It offers the correct level of flexibility without levying additional mountains of complexity on the languages.

---

* https://stackoverflow.com/questions/27733854/embedding-when-to-use-pointer/27733969#27733969
* http://www.hydrogen18.com/blog/golang-embedding.html