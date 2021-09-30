Разработать модульное приложение, позволяющее вести учет ключевых объектов, имеющих некоторый вид связи (агрегация или композиция) с двумя другими объектами, находящимися друг с другом в такой же взаимосвязи, для одной из приведенных в конце предметных областей.

Реализовать в программе возможность вывода на экран следующей информации:

* список ключевых объектов, связанных с одним из представителей первого объекта;
* список ключевых объектов, связанных со каждым представителем второго объекта;
* список ключевых объектов, связанных с одним из представителей второго объекта.

Выбранная тема:  информация о товарах на складе.

Исходный код программы см. в приложении.

Пример вывода программы:

```
Список продуктов:
 - holy-glitter    
 - rough-sunset    
 - dawn-silence    
 - solitary-forest 
 - cool-star       
 - throbbing-fire  
 - bitter-haze     
 - lingering-cherry
 - frosty-sky      
 - holy-river      
 - summer-water    
Список складов:
 - holy-glitter
 - rough-sunset
 - dawn-silence
 - solitary-forest
 - cool-star
 - throbbing-fire
Полный список товаров по складам:
склад holy-glitter
        секция 1
                holy-glitter x27
                cool-star x84
                throbbing-fire x91
                bitter-haze x93
                summer-water x40
        секция 2
                summer-water x75
                dawn-silence x8
                holy-river x29
склад rough-sunset
склад dawn-silence
склад solitary-forest
        секция 1
                rough-sunset x35
                dawn-silence x28
                cool-star x95
                lingering-cherry x52
                holy-river x76
        секция 2
                rough-sunset x26
                cool-star x23
                bitter-haze x1
                lingering-cherry x12
                holy-glitter x72
склад cool-star
        секция 1
                rough-sunset x42
                dawn-silence x24
                solitary-forest x59
                summer-water x55
        секция 2
                throbbing-fire x66
                bitter-haze x84
                lingering-cherry x76
                summer-water x57
                holy-glitter x32
                rough-sunset x1
                dawn-silence x9
                solitary-forest x54
склад throbbing-fire
        секция 1
                lingering-cherry x87
                rough-sunset x91
                dawn-silence x47
                cool-star x62
        секция 2
                dawn-silence x68
                throbbing-fire x87
                bitter-haze x93
                holy-river x54
                summer-water x91
        секция 3
                holy-glitter x3
                dawn-silence x9
                cool-star x9
                throbbing-fire x18
                lingering-cherry x77
                frosty-sky x73
```