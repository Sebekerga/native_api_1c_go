ПутьКомпоненты = "C:\Users\sebek\Documents\native_api_1c_go\out\lib.dll";

Сообщить("Starting loading AddIn");
Если НЕ ПодключитьВнешнююКомпоненту(ПутьКомпоненты, "Test", ТипВнешнейКомпоненты.Native) Тогда
	ВызватьИсключение("Unable to load AddIn");
Иначе
	Сообщить("AddIn loaded successfully");
КонецЕсли;

Сообщить("Starting creating AddIn object");
Попытка
	Объект = Новый("AddIn.Test.A");
	Сообщить("AddIn object created successfully, " + ТипЗнч(Объект));
Исключение
	Сообщить("Unable to create AddIn object");
КонецПопытки;

Сообщить("Starting getting property");
Попытка
	Сообщить("Property value: " + Объект.Свойство);
Исключение
	Сообщить("Unable to get property");
КонецПопытки;