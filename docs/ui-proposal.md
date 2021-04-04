# Пользовательский интерфейс mfd

MFD уже содержит огромное кол-во настроек, параметров и особенностей, которые нужно все помнить, 
для того что бы им можно было продуктивно пользоваться.

Интерфейс должен упростить пользование mfd-generator путем визуализации последовательности команд, 
которые сейчас пишутся полу-вручную через makefile.

Предлагается сделать интерфейс как SPA и/или нативное приложение mac/linux/windows на Flutter. 
Веб версия будет включена в mfd-generator путем вшивания html, ccs и js файлов, 
а у mfd-generator будет режим работы, когда он будет раздавать эти файлы как веб-сервис.

## Этап 1: настройка

2 поля ввода:

* строка подключения, например `postgres://postgres:postgres@localhost:5432/uteka?sslmode=disable`
* путь до `.mfd` файла
* выбирается язык?

## Этап 2: XML

### Настройка маппинга неймспейсов на сущности:

Таблица/список. 
Добавляем неймспейсы, в неймспейсы добавляем сущности и указываем в какой таблице сущность хранится.

![mapping.png](mapping.png)

Маппинг сохраняется в `.mfd` файл в раздел TablesMapping. Из примера выше:
```xml
<TablesMapping>
    <Namespace Name="catalogue">
        <Entity TableName="pharmacies">Pharmacy</Entity>
        <Entity TableName="pharmacyNetworks">PharmacyNetwork</Entity>
        <Entity TableName="cities">City</Entity>
    </Namespace>
    <Namespace Name="common">
        <Entity TableName="users">User</Entity>
    </Namespace>
    <Namespace Name="order">
        <Entity TableName="orders">Order</Entity>
        <Entity TableName="siteUsers">SiteUser</Entity>
    </Namespace>
</TablesMapping>
```

### Остальные настройки xml

* Указываем go-pg версию.
* Список кастомных типов:

![custom-types.png](custom-types.png)

### Этап 3: редактирование namespace.xml файлов

##### Параметры сущности

Можно поменять:
- имя сущности (тогда поменяется и в `.mfd` файле и маппинге).
- Неймспейс ???? (тогда нужно убирать из файла?! и перемещать в другой)
- таблицу (тоже менять в `.mfd` и в маппинге)

##### Редактирование аттрибутов сущностей

Выбираем сущность из списка. Открывается таблица, строки — поля таблицы базы, колонки — настройки: `Name`, `DBName`, `GoType` и другие.

##### Редактирование поисков

То же самое, что и в случае с аттрибутами, только редактируем поиски.

