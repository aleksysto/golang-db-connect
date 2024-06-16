package markers
const (
    NullMarker byte = 0xC0
    BooleanFalseMarker byte = 0xC2
    BooleanTrueMarker byte = 0xC3
    //TinyIntMarker byte= 0x00 // #todo
    Int8Marker byte = 0xC8
    Int16Marker byte = 0xC9
    Int32Marker byte = 0xCA
    Int64Marker byte = 0xCB
    FloatMarker byte = 0xC1
    Byte8Marker byte = 0xCC
    Byte16Marker byte = 0xCD
    Byte32Marker byte = 0xCE
    TinyStringMarker byte = 0x80
    String8Marker byte = 0xD0
    String16Marker byte = 0xD1
    String32Marker byte = 0xD2
    TinyListMarker byte = 0x90
    List8Marker byte = 0xD4
    List16Marker byte = 0xD5
    List32Marker byte = 0xD6
    TinyDictMarker byte = 0xA0
    Dict8Marker byte = 0xD8
    Dict16Marker byte = 0xD9
    Dict32Marker byte = 0xDA
    TinyStructMarker byte = 0xB0
)
