# ecommerceApp
  В проекте учавствуют:
    
      Rustem Nygmet 200103424
      Telzhan Mukhadas 200103387
      Yerlan Yeldesov 200103084
      Berik Serikbay 200103348
      Aina Akimniyazova 200103139
      
# Скачиваем MAMP, через него заходим в PhpMyAdmin, Создаем бд в MySql с таблицами: users, которые содержут поля id, name, email, username, pass. И дальше запускаем проект
 
Доп команды, если проект не запустился, сначала удалаем go.mod, затем пишем на терминале:
        
        go mod init github.com/Krasav4ik01/ecommerceApp
        
        
 И затем введем команды ниже:
  
  
      go get -u github.com/go-sql-driver/mysql 
      
      go get golang.org/x/crypto/bcrypt
      
      go get github.com/gorilla/sessions
      
      go get github.com/go-playground/validator/v10


