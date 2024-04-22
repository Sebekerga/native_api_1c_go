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

Для Счетчик = 1 По 1 Цикл
	ИмяСвойства = СтрШаблон("Property%1", Счетчик);
	Попытка
		ЗначениеСвойства = Объект[ИмяСвойства];
		Сообщить(СтрШаблон("%1: %2", ИмяСвойства, ЗначениеСвойства));
	Исключение
		Сообщить("Unable to get property");
	КонецПопытки;
КонецЦикла;