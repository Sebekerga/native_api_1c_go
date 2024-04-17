
/* HOW IT WORKS
TODO: Add description
*/

#include <stdbool.h>
#include <stdint.h>
#include <wchar.h>
#include <stdlib.h>
#include <logging.h>

#ifndef ADDIN_H
#define ADDIN_H

#ifndef ADDIN_TYPE
#define ADDIN_TYPE Component
#endif
typedef struct ADDIN_TYPE ADDIN_TYPE;

// forward declarations
typedef struct AddInInterface AddInInterface;

typedef struct InitDoneBaseVTable InitDoneBaseVTable;
static InitDoneBaseVTable *CreateInitDone();

typedef struct LanguageExtenderVTable LanguageExtenderVTable;
static LanguageExtenderVTable *CreateLanguageExtender();

typedef struct LocaleBaseVTable LocaleBaseVTable;
static LocaleBaseVTable *CreateLocaleBase();

typedef struct UserLanguageBaseVTable UserLanguageBaseVTable;
static UserLanguageBaseVTable *CreateUserLanguageBase();

typedef struct MemoryManager MemoryManager;
typedef struct Connection Connection;

typedef struct AddInInterface
{
    // predefined interface
    InitDoneBaseVTable *init_done;
    LanguageExtenderVTable *language_extender;
    LocaleBaseVTable *locale;
    UserLanguageBaseVTable *user_language;

    // component storage
    MemoryManager *mem_manager;
    Connection *connection;
    ADDIN_TYPE *component;

} AddInInterface;

static AddInInterface *CreateGenericComponent()
{
    logToConsole("Creating AddIn in C");
    AddInInterface *addin = (AddInInterface *)malloc(sizeof(AddInInterface));
    addin->init_done = CreateInitDone();
    addin->language_extender = CreateLanguageExtender();
    addin->locale = CreateLocaleBase();
    addin->user_language = CreateUserLanguageBase();
    logToConsole("AddIn created in C");

    return addin;
}

static void InitComponent(AddInInterface *addin, ADDIN_TYPE *component)
{
    addin->component = component;
}

// represent successful connection with 1C that can be used to perform non trivial operations
typedef struct ConnectionVTable
{
    // de-allocator pointer
    int dtor_a;
// for linux only
#if defined(__linux__)
    int dtor_b;
#endif

    // `add_error` function pointer, used to send error to 1C
    // gets reference to self, message code, source, description and flag
    void (*add_error)(AddInInterface *self, unsigned short code, const int16_t *source, const int16_t *description, long flag);
} ConnectionVTable;

typedef struct Connection
{
    ConnectionVTable *vtable;
} Connection;

typedef struct InitDoneBaseVTable
{
    // de-allocator pointer
    int dtor_a;
// for linux only
#if defined(__linux__)
    int dtor_b;
#endif

    // `init` function pointer, used to initialize component
    // gets reference to self and pointer to 1C connection
    // should return true if initialization was successful
    bool (*init)(AddInInterface *self, Connection *connection);

    // `set_mem_manager` function pointer, used to set memory manager
    // gets reference to self and pointer to 1C memory manager
    // should return true if initialization was successful
    bool (*set_mem_manager)(AddInInterface *self, MemoryManager *mem_manager);

    // `get_info` function pointer, used to get version of the library
    // gets reference to self
    // should return `long`, representing version of the library
    long int (*get_info)(AddInInterface *self);

    // `done` function pointer, used to free resources
    // gets reference to self
    void (*done)(AddInInterface *self);
} InitDoneBaseVTable;

static bool _init(AddInInterface *self, Connection *connection)
{
    logToConsole("Init function was called");
    return true;
}

static bool _set_mem_manager(AddInInterface *self, MemoryManager *mem_manager)
{
    logToConsole("Set memory manager function was called");
    self->mem_manager = mem_manager;
    logToConsole("Memory manager was set");
    return true;
}

static long int _get_info(AddInInterface *self)
{
    logToConsole("Get info function was called");
    return 2000;
}

static void _done(AddInInterface *self)
{
    logToConsole("Done function was called");
}

static InitDoneBaseVTable *CreateInitDone()
{
    InitDoneBaseVTable *returnStruct = (InitDoneBaseVTable *)malloc(sizeof(InitDoneBaseVTable));
    returnStruct->init = _init;
    returnStruct->set_mem_manager = _set_mem_manager;
    returnStruct->get_info = _get_info;
    returnStruct->done = _done;
    return returnStruct;
}

typedef struct LanguageExtenderVTable
{
    // de-allocator pointer
    int dtor_a;
// for linux only
#if defined(__linux__)
    int dtor_b;
#endif

    // `register_extension_as` function pointer, used to register extension
    // gets reference to self and pointer to extension name
    // returns true if extension was registered successfully
    bool (*register_extension_as)(AddInInterface *self, wchar_t *extension_name);

    // `get_n_props` function pointer, used to get number of properties
    // gets reference to self and return number of properties
    long int (*get_n_props)(AddInInterface *self);

    // `find_prop` function pointer, used to find property by name
    // gets reference to self and pointer to property name
    // returns index of property or -1 if property was not found
    long int (*find_prop)(AddInInterface *self, wchar_t *prop_name);

    // `get_prop_name` function pointer, used to get property name
    // gets reference to self, index of property and index of language
    // returns pointer to property name
    wchar_t *(*get_prop_name)(AddInInterface *self, long int prop_index, long int lang_index);

    // `get_prop_val` function pointer, used to get property value
    // gets reference to self, index of property and pointer to value
    // returns true if value was got successfully
    bool (*get_prop_val)(AddInInterface *self, long int prop_index, void *value);

    // `set_prop_val` function pointer, used to set property value
    // gets reference to self, index of property and pointer to value
    // returns true if value was set successfully
    bool (*set_prop_val)(AddInInterface *self, long int prop_index, void *value);

    // `is_prop_readable` function pointer, used to check if property is readable
    // gets reference to self and index of property
    // returns true if property is readable
    bool (*is_prop_readable)(AddInInterface *self, long int prop_index);

    // `is_prop_writable` function pointer, used to check if property is writable
    // gets reference to self and index of property
    // returns true if property is writable
    bool (*is_prop_writable)(AddInInterface *self, long int prop_index);

    // `get_n_methods` function pointer, used to get number of methods
    // gets reference to self and return number of methods
    long int (*get_n_methods)(AddInInterface *self);

    // `find_method` function pointer, used to find method by name
    // gets reference to self and pointer to method name
    // returns index of method or -1 if method was not found
    long int (*find_method)(AddInInterface *self, wchar_t *method_name);

    // `get_method_name` function pointer, used to get method name
    // gets reference to self, index of method and index of language
    // returns pointer to method name
    wchar_t *(*get_method_name)(AddInInterface *self, long int method_index, long int lang_index);

    // `get_n_params` function pointer, used to get number of method parameters
    // gets reference to self and index of method
    // returns number of parameters
    long int (*get_n_params)(AddInInterface *self, long int method_index);

    // `get_param_default_val` function pointer, used to get default value of method parameter
    // gets reference to self, index of method, index of parameter and pointer to value
    // returns true if value was got successfully
    bool (*get_param_default_val)(AddInInterface *self, long int method_index, long int param_index, void *value);

    // `call_as_proc` function pointer, used to call method as procedure
    // gets reference to self, index of method and pointer to parameters
    // returns true if method was called successfully
    bool (*call_as_proc)(AddInInterface *self, long int method_index, void *params);

    // `call_as_func` function pointer, used to call method as function
    // gets reference to self, index of method, pointer to parameters and pointer to result
    // returns true if method was called successfully
    bool (*call_as_func)(AddInInterface *self, long int method_index, void *params, void *result);

} LanguageExtenderVTable;

extern bool _register_extension_as(AddInInterface *self, wchar_t *extension_name);
extern long int _get_n_props(AddInInterface *self);
extern long int _find_prop(AddInInterface *self, wchar_t *prop_name);
extern wchar_t *_get_prop_name(AddInInterface *self, long int prop_index, long int lang_index);
extern bool _get_prop_val(AddInInterface *self, long int prop_index, void *value);
extern bool _set_prop_val(AddInInterface *self, long int prop_index, void *value);
extern bool _is_prop_readable(AddInInterface *self, long int prop_index);
extern bool _is_prop_writable(AddInInterface *self, long int prop_index);
extern long int _get_n_methods(AddInInterface *self);
extern long int _find_method(AddInInterface *self, wchar_t *method_name);
extern wchar_t *_get_method_name(AddInInterface *self, long int method_index, long int lang_index);
extern long int _get_n_params(AddInInterface *self, long int method_index);
extern bool _get_param_default_val(AddInInterface *self, long int method_index, long int param_index, void *value);
extern bool _call_as_proc(AddInInterface *self, long int method_index, void *params);
extern bool _call_as_func(AddInInterface *self, long int method_index, void *params, void *result);

static LanguageExtenderVTable *CreateLanguageExtender()
{
    logToConsole("Creating LanguageExtender in C");
    LanguageExtenderVTable *languageExtender = (LanguageExtenderVTable *)malloc(sizeof(LanguageExtenderVTable));
    languageExtender->register_extension_as = _register_extension_as;
    languageExtender->get_n_props = _get_n_props;
    languageExtender->find_prop = _find_prop;
    languageExtender->get_prop_name = _get_prop_name;
    languageExtender->get_prop_val = _get_prop_val;
    languageExtender->set_prop_val = _set_prop_val;
    languageExtender->is_prop_readable = _is_prop_readable;
    languageExtender->is_prop_writable = _is_prop_writable;
    languageExtender->get_n_methods = _get_n_methods;
    languageExtender->find_method = _find_method;
    languageExtender->get_method_name = _get_method_name;
    languageExtender->get_n_params = _get_n_params;
    languageExtender->get_param_default_val = _get_param_default_val;
    languageExtender->call_as_proc = _call_as_proc;
    languageExtender->call_as_func = _call_as_func;

    logToConsole("Created LanguageExtender in C");
    return languageExtender;
}

typedef struct LocaleBaseVTable
{
    // de-allocator pointer
    int dtor_a;
// for linux only
#if defined(__linux__)
    int dtor_b;
#endif

    // `set_locale` function pointer, used to set locale
    // gets reference to self and pointer to locale string
    void (*set_locale)(AddInInterface *self, const wchar_t *locale);

} LocaleBaseVTable;

static void _set_locale(AddInInterface *self, const wchar_t *locale)
{
    logToConsole("Setting locale in C");
}

static LocaleBaseVTable *CreateLocaleBase()
{
    LocaleBaseVTable *vtable = (LocaleBaseVTable *)malloc(sizeof(LocaleBaseVTable));
    vtable->set_locale = _set_locale;
    return vtable;
}

typedef struct UserLanguageBaseVTable
{
    // de-allocator pointer
    int dtor_a;
// for linux only
#if defined(__linux__)
    int dtor_b;
#endif

    // `set_user_language` function pointer, used to set user language
    // gets reference to self and pointer to user language string
    void (*set_user_language)(AddInInterface *self, const wchar_t *user_language);

} UserLanguageBaseVTable;

static void _set_user_language(AddInInterface *self, const wchar_t *user_language)
{
    logToConsole("Setting user language in C");
}

static UserLanguageBaseVTable *CreateUserLanguageBase()
{
    UserLanguageBaseVTable *vtable = (UserLanguageBaseVTable *)malloc(sizeof(UserLanguageBaseVTable));
    vtable->set_user_language = _set_user_language;
    return vtable;
}

typedef struct MemoryManagerVTable
{
    // de-allocator pointer
    int dtor_a;
// for linux only
#if defined(__linux__)
    int dtor_b;
#endif

    // `alloc_memory` function pointer, used to allocate memory
    // gets reference to self, pointer to pointer and size of memory
    // returns true if memory was allocated successfully
    bool (*alloc_memory)(AddInInterface *self, void **memory, size_t size);

    // `free_memory` function pointer, used to free memory
    // gets reference to self and pointer to pointer
    void (*free_memory)(AddInInterface *self, void **memory);
} MemoryManagerVTable;

typedef struct MemoryManager
{
    // pointer to vtable
    MemoryManagerVTable *vtable;
} MemoryManager;

#undef ADDIN_TYPE

#endif