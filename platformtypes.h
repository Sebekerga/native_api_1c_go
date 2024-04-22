#include <stdint.h>
#include <stdbool.h>
#include <stdlib.h>
#include <wchar.h>

#ifndef PLATFORMTYPES_H
#define PLATFORMTYPES_H

typedef struct PlatformTime
{
    int sec;
    int min;
    int hour;
    int mday;
    int mon;
    int year;
    int wday;
    int yday;
    int isdst;
} PlatformTime;

typedef enum PlatformType
{
    PlatTypeEmpty = 0,
    PlatTypeNull,
    PlatTypeInt16,     // int16_t
    PlatTypeInt32,     // int32_t
    PlatTypeFloat32,   // float
    PlatTypeFloat64,   // double
    PlatTypeDate,      // DATE (double)
    PlatTypeTime,      // struct tm
    PlatTypePStr,      // struct str    string
    PlatTypeInterface, // struct iface
    PlatTypeError,     // int32_t errCode
    PlatTypeBool,      // bool
    PlatTypeVariant,   // struct _tVariant *
    PlatTypeInt8,      // int8_t
    PlatTypeUInt8,     // uint8_t
    PlatTypeUInt16,    // uint16_t
    PlatTypeUInt32,    // uint32_t
    PlatTypeInt64,     // int64_t
    PlatTypeUInt64,    // uint64_t
    PlatTypeInt,       // int   Depends on architecture
    PlatTypeUInt,      // unsigned int  Depends on architecture
    PlatTypeHResult,   // long hRes
    PlatTypeWStr,      // struct wstr
    PlatTypeBlob,      // means in struct str binary data contain
    PlatTypeClsID,     // UUID

    PlatTypeUndefined = 0xFFFF,
} PlatformType;

typedef struct WCharString
{
    wchar_t *str;
    uint32_t len;
} WCharString;

typedef union PlatformValue
{
    int8_t i8Val;
    int16_t shortVal;
    int32_t lVal;
    int intVal;
    unsigned int uintVal;
    int64_t llVal;
    uint8_t ui8Val;
    uint16_t ushortVal;
    uint32_t ulVal;
    uint64_t ullVal;
    int32_t errCode;
    long hRes;
    float fltVal;
    double dblVal;
    bool bVal;
    char chVal;
    wchar_t wchVal;
    double date;
    char IDVal[16];
    struct PlatformVar *pvarVal;
    struct PlatformTime tmVal;
    struct
    {
        void *pInterfaceVal;
        char InterfaceID[16];
    };
    struct WCharString wstrVal;
    struct
    {
        wchar_t *pwstrVal;
        uint32_t wstrLen;
    };
} PlatformValue;

typedef struct PlatformVar
{
    PlatformValue value;
    uint32_t elements;
    PlatformType ty;

} PlatformVar;

#endif // PLATFORMTYPES_H