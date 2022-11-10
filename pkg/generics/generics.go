package generics

type Slice interface {
	[]any
}

type SliceOfPointers interface {
	[]*any
}

type Setter[INPUT any, OUTPUT any] func(i INPUT) OUTPUT

//func WrapWithConversionSetter[INPUT any, OUTPUT any](setterFunc Setter[INPUT, OUTPUT],

type ItemConversionFunc[INPUT any, OUTPUT any] func(i INPUT) OUTPUT
type ConversionClosureFunc[INPUT any, OUTPUT any] func(conversionFunc SliceConversionFunc[INPUT, OUTPUT], input INPUT) (o OUTPUT)

func ConversionClosure[INPUT any, OUTPUT any, FUNC ItemConversionFunc[INPUT, OUTPUT]](f FUNC) (o ItemConversionFunc[INPUT, OUTPUT]) {
	return func(i INPUT) OUTPUT {
		return f(i)
	}
}

type SliceConversionFunc[INPUT any, OUTPUT any] func(i []INPUT) []OUTPUT
type SliceConversionClosureFunc[INPUT any, OUTPUT any] func(conversionFunc SliceConversionFunc[INPUT, OUTPUT], input []INPUT) (o []OUTPUT)

func ConvertSlice[INPUT any, OUTPUT any](conversionFunc ItemConversionFunc[INPUT, OUTPUT], input []INPUT) (o []OUTPUT) {
	o = make([]OUTPUT, 0, len(input))
	for i, v := range input {
		o[i] = conversionFunc(v)
	}
	return o
}

// SliceConversionClosure returns a function closure that calls ConvertSlice with the provided conversion function
// this is leveraged to defer conversion if condition checks are used
func SliceConversionClosure[INPUT any, OUTPUT any](f ItemConversionFunc[INPUT, OUTPUT]) (o SliceConversionFunc[INPUT, OUTPUT]) {
	return func(i []INPUT) []OUTPUT {
		return ConvertSlice(f, i)
	}
}

type ConditionCheck[INPUT any] func(i INPUT) bool

func SetWithConditions[INPUT any, OUTPUT any](setterFunc Setter[INPUT, OUTPUT], i INPUT, conditions ...ConditionCheck[INPUT]) OUTPUT {
	var o OUTPUT
	for _, condition := range conditions {
		if condition(i) != true {
			return o
		}
	}
	return setterFunc(i)
}

func SetWithConversion[INPUT any, COUTPUT any, OUTPUT any](setterFunc Setter[COUTPUT, OUTPUT], conversionFunc ItemConversionFunc[INPUT, COUTPUT], i INPUT) OUTPUT {
	return setterFunc(conversionFunc(i))
}

func SetIfNonNilWithConversion[INPUT any, COUTPUT any, OUTPUT any](setterFunc Setter[COUTPUT, OUTPUT], conversionFunc ItemConversionFunc[*INPUT, COUTPUT], i *INPUT) OUTPUT {
	var o OUTPUT
	if i == nil {
		// return the default value for OUTPUT type
		return o
	}
	items := conversionFunc(i)
	return setterFunc(items)
}

func SetIfNonNil[INPUT any, OUTPUT any](setterFunc Setter[INPUT, OUTPUT], i *INPUT) OUTPUT {
	if i != nil {
		return setterFunc(*i)
	}
	// return the default value for OUTPUT type
	var o OUTPUT
	return o
}

func CopyToNewPointer[INPUT any](i *INPUT) (o *INPUT) {
	if i == nil {
		return nil
	}
	o = &(*i)
	return o
}

func CopyPointerSlice[INPUT any](input []*INPUT) (o []*INPUT) {
	o = make([]*INPUT, 0, len(input))
	for i, v := range input {
		if v == nil {
			o[i] = nil
		}
		o[i] = &(*v)
	}
	return o
}

/*
func SetSliceWithConversion[SETTEROUTPUT any, INPUT any, OUTPUT any](setterFunc Setter[[]OUTPUT, SETTEROUTPUT], conversionFunc SliceConversion[INPUT, OUTPUT], i []INPUT) {
	// return the default value for OUTPUT type
	var o SETTEROUTPUT
	if i != nil {
		items := conversionFunc(i)
		setterFunc(items)
	}
	// return o
}
*/

// func ConvertSlice[INPUT any, OUTPUT any](input []INPUT, conversionFunc ConversionFunction[INPUT, OUTPUT]) (o []OUTPUT) {

func Dereference[INPUT any](i *INPUT) INPUT {
	return *i
}

func DefaultIfNil[INPUT any](i *INPUT) INPUT {
	var d INPUT
	if i == nil {
		return d
	}
	return Dereference(i)
}

/*
func ConvertPointerSlice[INPUT any, OUTPUT any](input []*INPUT, conversionFunc GenericConversionFunction[INPUT, OUTPUT]) (o []OUTPUT) {
	o = make([]OUTPUT, 0, len(input))
	for i, v := range input {
		o[i] = conversionFunc(*v)
	}
	return o
}
*/
