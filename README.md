# generate
项目中对数据库表建立实体类代码经常作为一种固定写法，这个工具旨在生成这类代码

目前v0.1版本，只实现了mysql-->go dao层的生成 和 mysql-->java entity&mapper 生成

可下载generateCode.exe可执行文件进行代码生成

后续会持续迭代

若要自定义生成方法，可选择性实现IDatabaseStrategy, ILanguageStrategy, ICodeGenerator, 然后调用Executer执行器生成